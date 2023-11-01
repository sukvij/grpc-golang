package main

// Welcome to channel go guruji

// Topic grpc bi-directional streaming

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	proto "github.com/sukvij/grpc-golang/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	// Connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)

	r := gin.Default()
	r.GET("/sent-message/:message", clientConnectionServer)
	r.Run(":8000") // 8080

}

func clientConnectionServer(c *gin.Context) {
	variable := c.Param("message")
	userId, _ := strconv.ParseInt(variable, 10, 64)
	req := &proto.GetUserByIdInput{UserId: int32(userId)}
	fmt.Println(userId)
	res, _ := client.GetUser(context.TODO(), req)

	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
