package services

import (
	"context"
	"fmt"

	"github.com/izabelrodrigues/fullcycle-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	//Insert database
	fmt.Println(req.Nome)

	return &pb.User{
		Id: "123",
		Nome: req.GetNome(),
		Email: req.GetEmail(),
	}, nil
}