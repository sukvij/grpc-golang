package repository

import (
	"context"
	"fmt"
	"io"

	proto "github.com/sukvij/grpc-golang/protoc"
	// userModel "github.com/sukvij/grpc-golang/user/model"
)

var client proto.ExampleClient

func GetAllUser() (*[]proto.User, error) {
	var result *[]proto.User
	stream, err := client.GetAllUser(context.TODO(), &proto.Empty{})
	if err != nil {
		fmt.Println("Something error")
		return nil, err
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server message:- ", message)
	}
	return result, nil
}
