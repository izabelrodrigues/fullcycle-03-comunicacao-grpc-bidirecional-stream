package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/izabelrodrigues/fullcycle-grpc-client-stream/pb"
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
	AddUsers(client)

}


func AddUsers(client pb.UserServiceClient) {
	
	reqs := []*pb.User{
		{
			Id:    "i1",
			Nome:  "Iza1",
			Email: "iza1@iza.com",
		},
		{
			Id:    "i2",
			Nome:  "Iza2",
			Email: "iza2@iza.com",
		},
		{
			Id:    "i3",
			Nome:  "Iza3",
			Email: "iza3@iza.com",
		},
		{
			Id:    "i4",
			Nome:  "Iza4",
			Email: "iza4@iza.com",
		},
		{
			Id:    "i5",
			Nome:  "Iza5",
			Email: "iza5@iza.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	//Pecorrendo a lista e enviando
	for _, req := range reqs {

		stream.Send(req)
		time.Sleep(time.Second * 3)

	}

	res, err := stream.CloseAndRecv()


	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}