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

func (repository *Repository) CreateUser() (*userMode.User, error) {
	user := repository.User
	client := repository.Client

	req := mapper.MappingFromUserModelToProtoModel(*user)
	res, _ := client.CreateUser(context.TODO(), req)
	fmt.Println(res)

	x := mapper.Mapping(res)
	return &x, nil
}

func (repository *Repository) GetUserById() (*userMode.User, error) {
	client := repository.Client
	fmt.Println("repo - ", repository.User)
	req := &proto.UserIdInput{UserId: repository.User.Id}
	res, _ := client.GellUserById(context.TODO(), req)
	fmt.Println(res)
	x := mapper.Mapping(res)
	return &x, nil
}

func (repository *Repository) GetAllUser() ([]*userMode.User, error) {
	client := repository.Client
	stream, err := client.GetAllUser(context.TODO(), &proto.Empty{})
	if err != nil {
		fmt.Println("Something error")
		return nil, err
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
	return allUsers, nil
}
