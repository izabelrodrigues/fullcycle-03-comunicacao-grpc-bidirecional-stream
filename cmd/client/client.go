package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/izabelrodrigues/fullcycle-grpc-bidirecional-stream/pb"
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
	AddUserStreamBoth(client)

}

func AddUserStreamBoth(client pb.UserServiceClient){

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

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

	//channel para o programa não morrer ao executar as functions
	wait := make(chan int)

	//Função anônima para rodar de forma assincrona o envio
	go func()  {
		for _, req := range reqs {
			fmt.Println("Sending User: ", req.Nome)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	//Função anônima para rodar de forma assincrona a resposta
	go func()  {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetNome(), res.GetStatus())
	
		}
		close(wait)
	}()

	<-wait

}
