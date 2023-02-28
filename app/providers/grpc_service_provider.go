package providers

import (
	grpcfacade "gopkg.in/go-mixed/framework.v1/facades/grpc"

	"go-app/app/grpc"
	"go-app/routes"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register() {
	//Add Grpc interceptors
	kernel := grpc.Kernel{}
	grpcfacade.UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	grpcfacade.UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot() {
	//Add routes
	routes.Grpc()
}
