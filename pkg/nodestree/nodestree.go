// Copyright 2023 Nautes Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodestree

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	CRDContentType = iota
	StringContentType
)

type Resource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata"`
}

type FileOptions struct {
	IgnorePath       []string
	IgnoreFile       []string
	ExclusionsSuffix []string
	MappingResources map[string]interface{}
	ContentType      int
}

type Node struct {
	Name     string
	Path     string
	Children []*Node
	IsDir    bool
	Content  interface{}
	Level    int
	Kind     string
}

type checkFn func(options CompareOptions, in *nodesTree) error

type nodesTree struct {
	fileOptions *FileOptions
	checks      []checkFn
	client      client.Client
	config      *Config
	operators   []NodesOperator
	nodes       *Node
}

type CompareOptions struct {
	Nodes       Node
	ProductName string
}

func NewNodestree(fileOptions *FileOptions, config *Config, client client.Client) NodesTree {
	return &nodesTree{
		fileOptions: fileOptions,
		checks: []checkFn{
			CheckResourceRepeatName,
			CheckEffectiveResourceLayout,
			CheckResouceReference,
			CheckNumberOfResources,
		},
		client: client,
		config: config,
	}
}

func (in *nodesTree) AppendIgnoreFilePath(paths []string) {
	in.fileOptions.IgnorePath = append(in.fileOptions.IgnorePath, paths...)
}

func (in *nodesTree) GetFileOptions() *FileOptions {
	return in.fileOptions
}

func (in *nodesTree) GetResourceLayoutConfigs() *Config {
	return in.config
}

// Compare comparison between file tree and standard layout
func (in *nodesTree) Compare(options CompareOptions) error {
	config, err := NewConfig()
	if err != nil {
		return err
	}

	if len(config.Sub) > 0 && len(options.Nodes.Children) > 0 {
		for _, fn := range in.checks {
			err := fn(options, in)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (in *nodesTree) InsertNodes(nodes, resource *Node) (*Node, error) {
	mapping := make(map[string]*Node)
	NodesToMapping(nodes, mapping)

	if node, ok := mapping[resource.Path]; ok {
		if resource.IsDir {
			node.Children = append(node.Children, resource.Children...)
		} else {
			node.Content = resource.Content
		}
	} else {
		subPath := filepath.Dir(resource.Path)
		subName := strings.Split(subPath, "/")
		len := len(subName)
		name := subName[len-1]
		leval := resource.Level - 1

		subNode := &Node{
			Name:     name,
			Path:     subPath,
			Children: []*Node{resource},
			IsDir:    true,
			Level:    leval,
		}

		return in.InsertNodes(nodes, subNode)
	}

	return nodes, nil
}

func (in *nodesTree) GetNode(nodes *Node, kind, targetNodeName string) (node *Node) {
	for _, n := range nodes.Children {
		if !n.IsDir {
			k1 := GetResourceValue(n.Content, "TypeMeta", "Kind")
			n1 := GetResourceValue(n.Content, "ObjectMeta", "Name")
			if n1 == targetNodeName && k1 == kind {
				return n
			}
		}

		node = in.GetNode(n, kind, targetNodeName)
		if node != nil {
			return
		}
	}

	return
}

func (in *nodesTree) RemoveNode(nodes *Node, targetNode *Node) (*Node, error) {
	mapping := make(map[string]*Node)
	NodesToMapping(nodes, mapping)
	subPath := filepath.Dir(targetNode.Path)

	subNode, ok := mapping[subPath]
	if !ok {
		return nil, fmt.Errorf("the resource doesn't match beacase it cannot be found in %s", subPath)
	}

	for i, node := range subNode.Children {
		if node.Path == targetNode.Path {
			subNode.Children = append(subNode.Children[:i], subNode.Children[i+1:]...)
			break
		}
	}

	if len(subNode.Children) == 0 {
		parentPath := filepath.Dir(subPath)
		parentNode, ok := mapping[parentPath]
		if !ok {
			return nil, fmt.Errorf("the resource doesn't match beacase it cannot be found in %s", parentPath)
		}
		for i, node := range parentNode.Children {
			if node.Path == subPath {
				parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
			}
		}
	}

	return nodes, nil
}

func (i *nodesTree) AppendOperators(operator NodesOperator) {
	i.operators = append(i.operators, operator)
}

func (r *nodesTree) FilterIgnoreByLayout(path string) error {
	ignorePaths := make([]string, 0)
	layoutConfigs := r.GetResourceLayoutConfigs()

	layoutNames := make(map[string]struct{})
	for _, config := range layoutConfigs.Sub {
		layoutNames[config.Name] = struct{}{}
	}

	dir, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open dir: %s", err)
		return err
	}
	defer dir.Close()

	subDirs, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Printf("failed to read sub dir: %s", err)
		return err
	}

	ignorePaths = make([]string, 0, len(subDirs))

	for _, subdir := range subDirs {
		if _, exists := layoutNames[subdir]; !exists {
			ignorePaths = append(ignorePaths, subdir)
		}
	}

	r.AppendIgnoreFilePath(ignorePaths)

	return nil
}

func (i *nodesTree) Load(path string) (root Node, err error) {
	if strings.TrimSpace(path) == "" {
		err = fmt.Errorf("file or directory cannot be empty")
		return
	}

	root.Path = path
	root.Level += 1

	var file fs.FileInfo
	file, err = os.Stat(root.Path)
	if err != nil {
		return
	}
	root.Name = file.Name()
	root.IsDir = file.IsDir()

	if root.IsDir {
		err = explorerRecursive(&root, i.fileOptions, i.operators)
		if err != nil {
			return
		}
	}

	i.nodes = &root

	return
}

func (i *nodesTree) GetNodes() (*Node, error) {
	if i.nodes == nil {
		return nil, fmt.Errorf("the nodes is nill, please load the nodes")
	}

	return i.nodes, nil
}

// explorerRecursive traverse of the file tree
func explorerRecursive(node *Node, fileOptions *FileOptions, operators []NodesOperator) error {
	sub, err := ioutil.ReadDir(node.Path)
	if err != nil {
		return fmt.Errorf("directory does not exist or cannot be opened, %w", err)
	}

	for _, f := range sub {
		var tmp = path.Join(node.Path, f.Name())
		var child = &Node{
			Name:  f.Name(),
			Path:  tmp,
			IsDir: f.IsDir(),
			Level: node.Level + 1,
		}

		if ok := fileFiltering(fileOptions, f.Name(), tmp); ok {
			continue
		}

		if f.IsDir() {
			node.Children = append(node.Children, child)
			err = explorerRecursive(child, fileOptions, operators)
			if err != nil {
				return err
			}
		} else {
			fileType := path.Ext(f.Name())
			child.Name = strings.TrimSuffix(f.Name(), fileType)

			if fileOptions.ContentType == StringContentType {
				buffer, err := ioutil.ReadFile(child.Path)
				if err != nil {
					return err
				}
				child.Content = string(buffer)
			} else {
				cr, err := convertResource(child, operators)
				if err != nil {
					return err
				}
				child.Content = cr
			}

			node.Children = append(node.Children, child)
		}
	}

	return nil
}

func convertResource(child *Node, operators []NodesOperator) (cr interface{}, err error) {
	buffer, err := ioutil.ReadFile(child.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read %v file content, err: %w", child.Name, err)
	}

	if len(buffer) == 0 {
		return nil, fmt.Errorf("invalid file, content is empty in %s", child.Path)
	}

	jsonData, err := yaml.YAMLToJSON(buffer)
	if err != nil {
		return nil, fmt.Errorf("conversion yaml to json, err: %w", err)
	}

	var resource = &Resource{}
	err = json.Unmarshal(jsonData, resource)
	if err != nil {
		return nil, fmt.Errorf("the file is unable to unmarshal standard k8s structure, file path: %s, err: %w", child.Path, err)
	}

	if resource.Kind == "" {
		return nil, fmt.Errorf("node %s kind is empty", child.Name)
	}

	child.Kind = resource.Kind

	for _, o := range operators {
		cr = o.CreateResource(child.Kind)
		if cr != nil {
			err = json.Unmarshal(jsonData, &cr)
			if err != nil {
				return nil, err
			}

			break
		}
	}

	if cr == nil {
		return nil, fmt.Errorf("when loading the nodes tree, unable to generate resource node tree of type '%s'", resource.Kind)
	}

	return cr, nil
}

func fileFiltering(option *FileOptions, name, pathstr string) bool {
	if ok := IsInSlice(option.IgnoreFile, name); ok {
		return true
	}

	if ok := IsInSlice(option.IgnorePath, name); ok {
		return true
	}

	if ok := InContainsDir(pathstr, option.IgnorePath); ok {
		return true
	}

	if ok := IsInSlice(option.ExclusionsSuffix, path.Ext(name)); ok {
		return true
	}

	return false
}
