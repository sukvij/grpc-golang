package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/kataras/iris/v12"
	proto "github.com/sukvij/grpc-golang/protoc"

	"github.com/gin-gonic/gin"
	userController "github.com/sukvij/grpc-golang/user/controller"
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
	// implement rest api
	// r := gin.Default()
	// r.GET("/user", getAllUser)
	// r.GET("/user/:id", getUserById)
	// r.Run(":8000") // 8080
	app := iris.New()
	userController.UserApis(app)
	app.Listen(":8000")
}

func getUserById(c *gin.Context) {
	variable := c.Param("id")
	// userId, _ := strconv.ParseInt(variable, 10, 64)
	req := &proto.UserIdInput{UserId: variable}
	res, _ := client.GellUserById(context.TODO(), req)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}

func getAllUser(c *gin.Context) {
	stream, err := client.GetAllUser(context.TODO(), &proto.Empty{})
	if err != nil {
		fmt.Println("Something error")
		return
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server message:- ", message)
	}
	c.JSON(http.StatusOK, gin.H{
		"message_sent":    1,
		"message_recieve": 2,
	})
}
