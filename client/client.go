package main

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

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)

	r := gin.Default()
	r.GET("/user/:id", clientConnectionServer)
	r.Run(":8000")

}

func clientConnectionServer(c *gin.Context) {
	variable := c.Param("id")
	userId, _ := strconv.ParseInt(variable, 10, 64)
	req := &proto.GetUserByIdInput{UserId: int32(userId)}
	fmt.Println(userId)
	res, _ := client.GetUser(context.TODO(), req)

	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
