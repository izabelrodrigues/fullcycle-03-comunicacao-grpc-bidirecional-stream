package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/izabelrodrigues/fullcycle-grpc-stream/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)

	AddUserVerbose(client)

}

func AddUser(client pb.UserServiceClient) {
	
	req := &pb.User {
		Id:		"0",
		Nome:	"Iza-client",
		Email:	"iza@iza.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User {
		Id:		"0",
		Nome:	"Iza-client",
		Email:	"iza@iza.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())


	}
}