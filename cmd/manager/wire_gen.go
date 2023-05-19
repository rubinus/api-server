// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nautes-labs/api-server/internal/biz"
	"github.com/nautes-labs/api-server/internal/conf"
	"github.com/nautes-labs/api-server/internal/data"
	"github.com/nautes-labs/api-server/internal/server"
	"github.com/nautes-labs/api-server/internal/service"
	"github.com/nautes-labs/api-server/pkg/cluster"
	"github.com/nautes-labs/api-server/pkg/nodestree"
	"github.com/nautes-labs/pkg/pkg/nautesconfigs"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, nodesTree nodestree.NodesTree, config *configs.Config, client2 client.Client, clusteroperator cluster.ClusterRegistrationOperator) (*kratos.App, func(), error) {
	codeRepo, err := data.NewCodeRepo(config)
	if err != nil {
		return nil, nil, err
	}
	secretrepo, err := data.NewSecretRepo(config)
	if err != nil {
		return nil, nil, err
	}
	gitRepo, err := data.NewGitRepo(config)
	if err != nil {
		return nil, nil, err
	}
	resourcesUsecase := biz.NewResourcesUsecase(logger, codeRepo, secretrepo, gitRepo, nodesTree, config)
	codeRepoBindingUsecase := biz.NewCodeRepoCodeRepoBindingUsecase(logger, codeRepo, secretrepo, nodesTree, resourcesUsecase, config, client2)
	codeRepoUsecase := biz.NewCodeRepoUsecase(logger, codeRepo, secretrepo, nodesTree, config, resourcesUsecase, codeRepoBindingUsecase, client2)
	productUsecase := biz.NewProductUsecase(logger, codeRepo, secretrepo, gitRepo, config, resourcesUsecase, codeRepoUsecase)
	productService := service.NewProductService(productUsecase, config)
	grpcServer := server.NewGRPCServer(confServer, productService, logger)
	projectPipelineRuntimeUsecase := biz.NewProjectPipelineRuntimeUsecase(logger, codeRepo, nodesTree, resourcesUsecase)
	projectPipelineRuntimeService := service.NewProjectPipelineRuntimeService(projectPipelineRuntimeUsecase)
	deploymentRuntimeUsecase := biz.NewDeploymentRuntimeUsecase(logger, codeRepo, nodesTree, resourcesUsecase)
	deploymentruntimeService := service.NewDeploymentruntimeService(deploymentRuntimeUsecase)
	codeRepoService := service.NewCodeRepoService(codeRepoUsecase, config)
	codeRepoBindingService := service.NewCodeRepoBindingService(codeRepoBindingUsecase)
	projectUsecase := biz.NewProjectUsecase(logger, codeRepo, secretrepo, nodesTree, config, resourcesUsecase)
	projectService := service.NewProjectService(projectUsecase)
	environmentUsecase := biz.NewEnviromentUsecase(logger, config, codeRepo, nodesTree, resourcesUsecase)
	environmentService := service.NewEnvironmentService(environmentUsecase)
	dexRepo := data.NewDexRepo(client2)
	clusterUsecase := biz.NewClusterUsecase(logger, codeRepo, secretrepo, resourcesUsecase, config, client2, clusteroperator, dexRepo)
	clusterService := service.NewClusterService(clusterUsecase, config)
	serviceProductGroup := server.NewServiceGroup(projectPipelineRuntimeService, deploymentruntimeService, codeRepoService, codeRepoBindingService, productService, projectService, environmentService, clusterService)
	httpServer := server.NewHTTPServer(confServer, serviceProductGroup)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
