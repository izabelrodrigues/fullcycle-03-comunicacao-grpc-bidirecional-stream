package main

import (
	"log"
	"net"

	"github.com/izabelrodrigues/fullcycle-grpc-stream/pb"
	"github.com/izabelrodrigues/fullcycle-grpc-stream/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main()  {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	//Para testar com evans
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}