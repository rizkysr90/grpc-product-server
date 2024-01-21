package grpc

import (
	productGRPCService "grpc-server-product/internal/grpc/handler/product"
	productService "grpc-server-product/internal/service/product"
	"grpc-server-product/internal/store"

	"github.com/rizkysr90/my-protobuf/gen/go/personal/productservice/product"
	"google.golang.org/grpc"
)

func New(
	masterData []store.ProductData,
) (*grpc.Server, error) {
	server := grpc.NewServer()
	productService := productService.New(masterData)
	productGRPCService := productGRPCService.New(productService)
	product.RegisterProductServiceServer(server, productGRPCService)
	return server, nil
}
