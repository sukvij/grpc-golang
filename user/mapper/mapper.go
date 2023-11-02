package mapper

import (
	proto "github.com/sukvij/grpc-golang/protoc"
	userMode "github.com/sukvij/grpc-golang/user/model"
)

func Mapping(protoModel *proto.User) userMode.User {
	user := &userMode.User{}
	user.Id = protoModel.Id
	user.Fname = protoModel.FName
	user.City = protoModel.City
	user.Height = float64(protoModel.Height)
	user.Phone = protoModel.Phone
	user.Married = protoModel.Married
	return *user
}

func MappingFromUserModelToProtoModel(user userMode.User) *proto.User {
	protomodel := &proto.User{}
	protomodel.Id = user.Id
	protomodel.FName = user.Fname
	protomodel.City = user.City
	protomodel.Height = float32(user.Height)
	protomodel.Phone = user.Phone
	protomodel.Married = user.Married
	return protomodel
}
