package main

import (
	"github.com/kataras/iris/v12"

	proto "github.com/sukvij/grpc-golang/protoc"
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
	app := iris.New()
	userController.UserApis(app, client)
	// userApis := app.Party("user")
	// {
	// 	userApis.Get("", getAllUser)
	// 	userApis.Get("/:id", getUserById)
	// 	userApis.Post("", createUser)
	// }
	app.Listen(":8000")
}

// func createUser(c iris.Context) {
// 	var user userMode.User
// 	if err := c.ReadJSON(&user); err != nil {
// 		return
// 	}

// 	validate = validator.New()
// 	errs := validate.Struct(user)
// 	if errs != nil {
// 		c.JSON(errs.Error())
// 		return
// 	}
// 	req := &proto.User{Id: user.Id, FName: user.Fname, City: user.City, Phone: user.Phone, Height: float32(user.Height), Married: user.Married}
// 	res, _ := client.CreateUser(context.TODO(), req)
// 	c.JSON(res)
// }
