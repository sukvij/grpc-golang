package service

import (
	proto "github.com/sukvij/grpc-golang/protoc"
	userRepo "github.com/sukvij/grpc-golang/user/repository"
)

func GetAllUser() (*[]proto.User, error) {
	// var result *[]userModel.User
	result, err := userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}
