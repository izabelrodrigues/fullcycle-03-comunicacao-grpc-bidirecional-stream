package services

import (
	"io"
	"log"

	"github.com/izabelrodrigues/fullcycle-grpc-bidirecional-stream/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User: req,

		})

		if err != nil {
			log.Fatalf("Error sending stream to the client: %v", err)
		}

	}

}
