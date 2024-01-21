package main

import (
	"grpc-server-product/internal/grpc"
	"grpc-server-product/internal/store"
	"log"
	"net"
)

func main() {
	masterData := []store.ProductData{}
	getPort, portError := net.Listen("tcp", ":8888")
	if portError != nil {
		log.Fatalf("grpc: failed to listen on port 8888 : %v", portError)
	}
	grpcServer, grpcError := grpc.New(masterData)
	if grpcError != nil {
		log.Fatalf("grpc: failed to construct grpc server")
	}
	log.Println("Server gRPC running on port 8888")
	err := grpcServer.Serve(getPort)
	if err != nil {
		log.Fatalf("grpc: failed to server grpc server")
	}
}
