package controller

import (
	"fmt"
	"strconv"

	"github.com/kataras/iris/v12"
	proto "github.com/sukvij/grpc-golang/protoc"
	userModel "github.com/sukvij/grpc-golang/user/model"
	userService "github.com/sukvij/grpc-golang/user/service"
)

var client proto.ExampleClient

func UserApis(app *iris.Application, cl proto.ExampleClient) {
	client = cl
	AlluserApis := app.Party("/user")
	{
		AlluserApis.Get("/", getAllUsers)
		AlluserApis.Get("/:id", getUserById)
		AlluserApis.Post("/", createUser)
	}
}

func createUser(ctx iris.Context) {
	user := &userModel.User{}
	ctx.ReadJSON(&user)
	service := &userService.Service{User: user, Client: client}
	result := service.CreateUser()
	ctx.JSON(result)
}

func getAllUsers(ctx iris.Context) {
	user := &userModel.User{}
	service := &userService.Service{User: user, Client: client}
	result := service.GetAllUser()
	ctx.JSON(result)
}

func getUserById(ctx iris.Context) {
	variable := ctx.Params().Get("id")
	userId, _ := strconv.ParseInt(variable, 10, 64)
	fmt.Println("userid : ", userId)
	user := &userModel.User{Id: userId}
	service := &userService.Service{User: user, Client: client}
	result := service.GetUserById()
	ctx.JSON(result)
}
