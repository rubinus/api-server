// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.6.1
// source: project/v1/project.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationProjectDeleteProject = "/api.project.v1.Project/DeleteProject"
const OperationProjectGetProject = "/api.project.v1.Project/GetProject"
const OperationProjectListProjects = "/api.project.v1.Project/ListProjects"
const OperationProjectSaveProject = "/api.project.v1.Project/SaveProject"

type ProjectHTTPServer interface {
	DeleteProject(context.Context, *DeleteRequest) (*DeleteReply, error)
	GetProject(context.Context, *GetRequest) (*GetReply, error)
	ListProjects(context.Context, *ListsRequest) (*ListsReply, error)
	SaveProject(context.Context, *SaveRequest) (*SaveReply, error)
}

func RegisterProjectHTTPServer(s *http.Server, srv ProjectHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/products/{product_name}/projects/{project_name}", _Project_GetProject0_HTTP_Handler(srv))
	r.GET("/api/v1/products/{product_name}/projects", _Project_ListProjects0_HTTP_Handler(srv))
	r.POST("/api/v1/products/{product_name}/projects/{project_name}", _Project_SaveProject0_HTTP_Handler(srv))
	r.DELETE("/api/v1/products/{product_name}/projects/{project_name}", _Project_DeleteProject0_HTTP_Handler(srv))
}

func _Project_GetProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectGetProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProject(ctx, req.(*GetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReply)
		return ctx.Result(200, reply)
	}
}

func _Project_ListProjects0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectListProjects)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProjects(ctx, req.(*ListsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListsReply)
		return ctx.Result(200, reply)
	}
}

func _Project_SaveProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveRequest
		if err := ctx.Bind(&in.Body); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectSaveProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveProject(ctx, req.(*SaveRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveReply)
		return ctx.Result(200, reply)
	}
}

func _Project_DeleteProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectDeleteProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProject(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteReply)
		return ctx.Result(200, reply)
	}
}

type ProjectHTTPClient interface {
	DeleteProject(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	GetProject(ctx context.Context, req *GetRequest, opts ...http.CallOption) (rsp *GetReply, err error)
	ListProjects(ctx context.Context, req *ListsRequest, opts ...http.CallOption) (rsp *ListsReply, err error)
	SaveProject(ctx context.Context, req *SaveRequest, opts ...http.CallOption) (rsp *SaveReply, err error)
}

type ProjectHTTPClientImpl struct {
	cc *http.Client
}

func NewProjectHTTPClient(client *http.Client) ProjectHTTPClient {
	return &ProjectHTTPClientImpl{client}
}

func (c *ProjectHTTPClientImpl) DeleteProject(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*DeleteReply, error) {
	var out DeleteReply
	pattern := "/api/v1/products/{product_name}/projects/{project_name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProjectDeleteProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) GetProject(ctx context.Context, in *GetRequest, opts ...http.CallOption) (*GetReply, error) {
	var out GetReply
	pattern := "/api/v1/products/{product_name}/projects/{project_name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProjectGetProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) ListProjects(ctx context.Context, in *ListsRequest, opts ...http.CallOption) (*ListsReply, error) {
	var out ListsReply
	pattern := "/api/v1/products/{product_name}/projects"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProjectListProjects))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) SaveProject(ctx context.Context, in *SaveRequest, opts ...http.CallOption) (*SaveReply, error) {
	var out SaveReply
	pattern := "/api/v1/products/{product_name}/projects/{project_name}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectSaveProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Body, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
