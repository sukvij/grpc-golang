package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	proto "github.com/sukvij/grpc-golang/protoc"
	userMode "github.com/sukvij/grpc-golang/user/model"
	userRepo "github.com/sukvij/grpc-golang/user/repository"
)

var validate *validator.Validate

type Service struct {
	User   *userMode.User
	Client proto.ExampleClient
}

func NewService(user *userMode.User, Client proto.ExampleClient) *Service {
	return &Service{User: user, Client: Client}
}

func (service *Service) CreateUser() *userMode.User {
	user := service.User
	validate = validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		fmt.Println(errs.Error())
		return nil
	}
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res := repo.CreateUser()
	return res
}

func (service *Service) GetAllUser() *userMode.User {
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res := repo.GetAllUser()
	return res
}

func (service *Service) GetUserById() *userMode.User {
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res := repo.GetUserById()
	return res
}
