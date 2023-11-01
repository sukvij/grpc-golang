package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	userMode "github.com/sukvij/grpc-golang/model"
	proto "github.com/sukvij/grpc-golang/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient
var validate *validator.Validate

func main() {
	// Connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)
	// implement rest api
	r := gin.Default()
	r.GET("/user", getAllUser)
	r.GET("/user/:id", getUserById)
	r.POST("/user", createUser)
	r.Run(":8000") // 8080

}

func createUser(c *gin.Context) {
	user := &userMode.User{}
	if err := c.BindJSON(&user); err != nil {
		return
	}

	validate = validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errs.Error(),
		})
		return
	}
	fmt.Println(errs)
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
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

	var allUsers []*proto.User
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		allUsers = append(allUsers, message)
		fmt.Println("Server message:- ", message)
	}
	c.JSON(http.StatusOK, gin.H{
		"allusers": allUsers,
	})
}
