package main

import (
	"context"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	proto "github.com/sukvij/grpc-golang/protoc"
	userMode "github.com/sukvij/grpc-golang/user/model"
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
	app := iris.New()
	userApis := app.Party("user")
	{
		userApis.Get("", getAllUser)
		userApis.Get("/:id", getUserById)
		userApis.Post("", createUser)
	}
	app.Listen(":8000")
}

func createUser(c iris.Context) {
	var user userMode.User
	if err := c.ReadJSON(&user); err != nil {
		return
	}

	validate = validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		c.JSON(errs.Error())
		return
	}
	req := &proto.User{Id: user.Id, FName: user.Fname, City: user.City, Phone: user.Phone, Height: float32(user.Height), Married: user.Married}
	res, _ := client.CreateUser(context.TODO(), req)
	c.JSON(res)
}

func getUserById(c iris.Context) {
	variable := c.Params().Get("id")
	// userId, _ := strconv.ParseInt(variable, 10, 64)
	req := &proto.UserIdInput{UserId: variable}
	res, _ := client.GellUserById(context.TODO(), req)
	fmt.Println(res)
	c.JSON(res)
}

func getAllUser(c iris.Context) {
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
	c.JSON(allUsers)
}
