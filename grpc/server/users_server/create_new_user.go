package main

import (
	"context"
	"fmt"
	"log"

	pb "obaid/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("CreateNewUser was invoked with %v\n", in)

	data := userStruct{
		Name:     in.Name,
		User:     in.User,
		Password: in.Password,
	}

	res, err := user_collection.InsertOne(ctx, data)

	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.User{
		Id: oid.Hex(),
		Name: in.Name,
		User: in.User,
		Password: in.Password,
	}, nil
}
