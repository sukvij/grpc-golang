package controller

import (
	"github.com/kataras/iris/v12"
	userService "github.com/sukvij/grpc-golang/user/service"
)

func UserApis(app *iris.Application) {
	AllApis := app.Party("/user")
	{
		AllApis.Get("/", getAllUsers)
		// AllApis.Get("/{userId}", getUser)
		// AllApis.Post("/", createUser)
	}
}

func getAllUsers(ctx iris.Context) {
	result, err := userService.GetAllUser()
	if err != nil {
		ctx.JSON(err)
		return
	}
	ctx.JSON(result)
}
