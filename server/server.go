package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	proto "github.com/sukvij/grpc-golang/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var users map[int64]*proto.User

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	users = make(map[int64]*proto.User)
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer() // engine
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) CreateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	fmt.Println(req)
	users[req.Id] = req
	return req, nil
}

func (s *server) GellUserById(ctx context.Context, req *proto.UserIdInput) (*proto.User, error) {
	userId := req.UserId
	user := users[userId]
	return user, nil
}

func (s *server) GetAllUser(req *proto.Empty, strem proto.Example_GetAllUserServer) error {
	// users = make(map[int64]*proto.User)
	user1 := &proto.User{Id: 1, FName: "salman"}
	user2 := &proto.User{Id: 2, FName: "vinay"}
	user3 := &proto.User{Id: 3, FName: "vijju"}
	user4 := &proto.User{Id: 4, FName: "vijendra"}
	user5 := &proto.User{Id: 5, FName: "aamir"}
	users[1] = user1
	users[2] = user2
	users[3] = user3
	users[4] = user4
	users[5] = user5
	for _, element := range users {
		err := strem.Send(element)
		time.Sleep(time.Second)
		if err != nil {
			return errors.New("unable to send data from server")
		}
	}

	return nil
}
