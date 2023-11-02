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

func (service *Service) CreateUser() (*userMode.User, error) {
	user := service.User
	validate = validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		fmt.Println(errs.Error())
		return nil, errs
	}
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res, err := repo.CreateUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *Service) GetAllUser() ([]*userMode.User, error) {
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res, err := repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *Service) GetUserById() (*userMode.User, error) {
	repo := &userRepo.Repository{User: service.User, Client: service.Client}
	res, err := repo.GetUserById()
	if err != nil {
		return nil, err
	}
	return res, nil
}
