package main

import (
	pb "obaid/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userStruct struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	User     string             `bson:"user"`
	Password string             `bson:"password"`
}

func assignUserValues(data *userStruct) *pb.User {
	return &pb.User{
		Id:       data.ID.Hex(),
		Name: 	  data.Name,
		User: 	  data.User,
		Password: data.Password,
	}
}