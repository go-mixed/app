package providers

import (
	"gopkg.in/go-mixed/framework.v1/facades"

	"go-app/app/grpc"
	"go-app/routes"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register() {
	//Add Grpc interceptors
	kernel := grpc.Kernel{}
	facades.Grpc.UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc.UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot() {
	//Add routes
	routes.Grpc()
}
