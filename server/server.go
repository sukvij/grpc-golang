package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sukvij/grpc-golang/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedExampleServer
}

func main() {
	lis, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		log.Fatalf("failed to satrt the server %v", tcpErr)
	} else {
		fmt.Println("server satrted at port 9000")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExampleServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	if errs := grpcServer.Serve(lis); errs != nil {
		log.Fatalf("failed to start : %v", errs)
	}
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserByIdInput) (*pb.User, error) {
	fmt.Println("received request from client:", req.UserId)
	fmt.Println("hello from server")
	// user := make(map[int]User)
	user1 := &pb.User{Id: 1, Fname: "Streve", City: "LA"}
	// user2 := User{Id: 2, Fname: "Vijendra", City: "sikar"}
	// user[1] = user1
	// user[2] = user2
	return user1, nil
}
