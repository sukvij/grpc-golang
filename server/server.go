package main

import (
	"context"
	"errors"
	"net"
	"strconv"
	"time"

	proto "github.com/sukvij/grpc-golang/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var users map[int32]*proto.User

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
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

func (s *server) GellUserById(ctx context.Context, req *proto.UserIdInput) (*proto.User, error) {
	in := req.UserId
	userId, _ := strconv.ParseInt(in, 10, 64)
	user := users[int32(userId)]
	return user, nil
}

func (s *server) GetAllUser(req *proto.Empty, strem proto.Example_GetAllUserServer) error {
	users = make(map[int32]*proto.User)
	user1 := &proto.User{Id: 1, Name: "salman"}
	user2 := &proto.User{Id: 2, Name: "vinay"}
	user3 := &proto.User{Id: 3, Name: "vijju"}
	user4 := &proto.User{Id: 4, Name: "vijendra"}
	user5 := &proto.User{Id: 5, Name: "aamir"}
	users[1] = user1
	users[2] = user2
	users[3] = user3
	users[4] = user4
	users[5] = user5
	// users = append(users, user1)
	// users = append(users, user2)
	// users = append(users, user3)
	// users = append(users, user4)
	// users = append(users, user5)
	for _, element := range users {
		err := strem.Send(element)
		time.Sleep(time.Second)
		if err != nil {
			return errors.New("unable to send data from server")
		}
	}
	return nil
}
