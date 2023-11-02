package repository

import (
	"context"
	"fmt"
	"io"

	proto "github.com/sukvij/grpc-golang/protoc"
	mapper "github.com/sukvij/grpc-golang/user/mapper"
	userMode "github.com/sukvij/grpc-golang/user/model"
)

type Repository struct {
	User   *userMode.User
	Client proto.ExampleClient
}

func NewRepository(User *userMode.User, Client proto.ExampleClient) *Repository {
	return &Repository{User: User, Client: Client}
}

func (repository *Repository) CreateUser() *userMode.User {
	user := repository.User
	client := repository.Client

	req := &proto.User{Id: user.Id, FName: user.Fname, City: user.City, Phone: user.Phone, Height: float32(user.Height), Married: user.Married}
	res, _ := client.CreateUser(context.TODO(), req)
	fmt.Println(res)

	x := mapper.Mapping(res)
	return &x
}

func (repository *Repository) GetUserById() *userMode.User {
	client := repository.Client
	fmt.Println("repo - ", repository.User)
	req := &proto.UserIdInput{UserId: repository.User.Id}
	res, _ := client.GellUserById(context.TODO(), req)
	fmt.Println(res)
	x := mapper.Mapping(res)
	return &x
}

func (repository *Repository) GetAllUser() []*userMode.User {
	client := repository.Client
	stream, err := client.GetAllUser(context.TODO(), &proto.Empty{})
	if err != nil {
		fmt.Println("Something error")
		return nil
	}

	var allUsers []*userMode.User

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		temp := mapper.Mapping(message)
		allUsers = append(allUsers, &temp)
		fmt.Println("Server message:- ", message)
	}
	return allUsers
}