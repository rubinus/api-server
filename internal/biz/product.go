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

package biz

import (
	"context"
	"fmt"

	errors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	commonv1 "github.com/nautes-labs/api-server/api/common/v1"
	nautesconfigs "github.com/nautes-labs/pkg/pkg/nautesconfigs"
)

const (
	_ProductKind     = "Product"
	_ProductDestUser = "Argo"
)

type Group struct {
	Id          int32
	Name        string
	Visibility  string
	Description string
	Path        string
	WebUrl      string
	ParentId    int32
}

type ProjectNamespace struct {
	ID        int
	Name      string
	Path      string
	Kind      string
	FullPath  string
	AvatarURL string
	WebURL    string
}

type Project struct {
	Id                int32
	Name              string
	Visibility        string
	Description       string
	Path              string
	WebUrl            string
	SshUrlToRepo      string
	HttpUrlToRepo     string
	PathWithNamespace string
	Namespace         *ProjectNamespace
}

// ProductUsecase is a Product usecase.
type ProductUsecase struct {
	log              *log.Helper
	codeRepo         CodeRepo
	secretRepo       Secretrepo
	gitRepo          GitRepo
	configs          *nautesconfigs.Config
	resourcesUsecase *ResourcesUsecase
	codeRepoUsecase  *CodeRepoUsecase
}

type GroupAndProjectItem struct {
	Group   *Group
	Project *Project
}

func NewProductUsecase(logger log.Logger, codeRepo CodeRepo, secretRepo Secretrepo, gitRepo GitRepo, configs *nautesconfigs.Config, resourcesUsecase *ResourcesUsecase, codeRepoUsecase *CodeRepoUsecase) *ProductUsecase {
	return &ProductUsecase{log: log.NewHelper(logger), codeRepo: codeRepo, secretRepo: secretRepo, gitRepo: gitRepo, configs: configs, resourcesUsecase: resourcesUsecase, codeRepoUsecase: codeRepoUsecase}
}

func (p *ProductUsecase) GetGroupAndDefaultProject(ctx context.Context, productName string) (*GroupAndProjectItem, error) {
	group, err := GetGroup(ctx, p.codeRepo, productName)
	if err != nil {
		if ok := commonv1.IsGroupNotFound(err); ok {
			return nil, nil
		}

		return nil, err
	}

	pid := fmt.Sprintf("%s/%s", group.Path, p.configs.Git.DefaultProductName)
	project, err := GetProject(ctx, p.codeRepo, pid)
	if err != nil {
		if ok := commonv1.IsProjectNotFound(err); ok {
			return nil, nil
		}

		return nil, err
	}

	return &GroupAndProjectItem{
		Group:   group,
		Project: project,
	}, nil
}

func (p *ProductUsecase) GetProduct(ctx context.Context, productName string) (*GroupAndProjectItem, error) {
	item, err := p.GetGroupAndDefaultProject(ctx, productName)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, fmt.Errorf("there are no default projects under this product")
	}

	return item, nil
}

func (p *ProductUsecase) ListProducts(ctx context.Context) ([]*GroupAndProjectItem, error) {
	groups, err := p.codeRepo.ListAllGroups(ctx)
	if err != nil {
		return nil, err
	}

	var products []*GroupAndProjectItem
	for _, group := range groups {
		product, err := p.GetGroupAndDefaultProject(ctx, group.Name)
		if err != nil {
			return nil, err
		}
		if product != nil {
			products = append(products, &GroupAndProjectItem{
				Group:   product.Group,
				Project: product.Project,
			})
		}
	}

	return products, nil
}

func (p *ProductUsecase) SaveProduct(ctx context.Context, productName string, gitOptions *GitGroupOptions) (group *Group, project *Project, err error) {
	group, err = GetGroup(ctx, p.codeRepo, productName)
	e := errors.FromError(err)
	if err != nil && e.Code != 404 {
		return nil, nil, err
	}

	if err != nil && e.Code == 404 {
		group, err = CreateGroup(ctx, p.codeRepo, gitOptions)
		if err != nil {
			return
		}
	} else {
		group, err = UpdateGroup(ctx, p.codeRepo, p.configs, int(group.Id), gitOptions)
		if err != nil {
			return
		}
	}

	project, err = p.saveDefaultProject(ctx, group)
	if err != nil {
		return
	}

	err = p.grantAuthorizationDefaultProject(ctx, project)
	if err != nil {
		return
	}

	return group, project, nil
}

func (p *ProductUsecase) saveDefaultProject(ctx context.Context, group *Group) (*Project, error) {
	defaultProjectPath := fmt.Sprintf("%s/%s", group.Path, p.configs.Git.DefaultProductName)
	project, err := p.codeRepo.GetCodeRepo(ctx, defaultProjectPath)
	if err != nil {
		opt := &GitCodeRepoOptions{
			Gitlab: &GitlabCodeRepoOptions{
				Name: p.configs.Git.DefaultProductName,
			},
		}

		project, err = p.codeRepo.CreateCodeRepo(ctx, int(group.Id), opt)
		if err != nil {
			return nil, err
		}

		user, email, err := p.codeRepo.GetCurrentUser(ctx)
		if err != nil {
			return nil, err
		}

		param := &CloneRepositoryParam{
			URL:   project.HttpUrlToRepo,
			User:  user,
			Email: email,
		}
		localPath, err := p.gitRepo.Clone(ctx, param)
		if err != nil {
			return nil, err
		}
		defer cleanCodeRepo(localPath)

		err = p.resourcesUsecase.SaveDeployConfig(nil, localPath)
		if err != nil {
			return nil, err
		}

		err = p.resourcesUsecase.SaveConfig(ctx, localPath)
		if err != nil {
			return nil, err
		}
	}

	return project, nil
}

func (p *ProductUsecase) grantAuthorizationDefaultProject(ctx context.Context, project *Project) error {
	projectDeployKey, err := p.codeRepoUsecase.saveDeployKey(ctx, int(project.Id), false)
	if err != nil {
		return err
	}

	if err := p.codeRepoUsecase.removeInvalidDeploykey(ctx, int(project.Id), projectDeployKey); err != nil {
		return err
	}

	err = p.secretRepo.AuthorizationSecret(ctx, int(project.Id), _ProductDestUser, string(p.configs.Git.GitType), p.configs.Secret.Vault.MountPath)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUsecase) DeleteProduct(ctx context.Context, productID string) error {
	group, err := p.codeRepo.GetGroup(ctx, productID)
	if err != nil {
		return err
	}

	codeRepos, err := p.codeRepo.ListGroupCodeRepos(ctx, int(group.Id))
	if err != nil {
		return err
	}

	if len(codeRepos) > 1 {
		return fmt.Errorf("here are multiple project in %v, unable to delete", group.Path)
	}

	if len(codeRepos) == 1 {
		defaultProjectPath := fmt.Sprintf("%v/%v", group.Path, p.configs.Git.DefaultProductName)
		project, err := p.codeRepo.GetCodeRepo(ctx, defaultProjectPath)
		if err != nil {
			return err
		}

		err = p.secretRepo.DeleteSecret(ctx, int(project.Id), DefaultUser, string(ReadOnly))
		if err != nil {
			return err
		}
	}

	err = p.codeRepo.DeleteGroup(ctx, int(group.Id))
	if err != nil {
		return err
	}

	return nil
}

func GetProject(ctx context.Context, codeRepo CodeRepo, pid interface{}) (project *Project, err error) {
	project, err = codeRepo.GetCodeRepo(ctx, pid)
	if err != nil {
		return
	}

	return
}

func GetGroup(ctx context.Context, codeRepo CodeRepo, gid interface{}) (group *Group, err error) {
	group, err = codeRepo.GetGroup(ctx, gid)
	if err != nil {
		return
	}

	return
}

func CreateGroup(ctx context.Context, codeRepo CodeRepo, gitOptions *GitGroupOptions) (group *Group, err error) {
	group, err = codeRepo.CreateGroup(ctx, gitOptions)
	if err != nil {
		return
	}

	return
}

func UpdateGroup(ctx context.Context, codeRepo CodeRepo, configClient *nautesconfigs.Config, gid int, git *GitGroupOptions) (group *Group, err error) {
	_, err = codeRepo.ListGroupCodeRepos(ctx, gid)
	if err != nil {
		return
	}

	group, err = codeRepo.UpdateGroup(ctx, gid, git)
	if err != nil {
		return nil, err
	}

	return
}
