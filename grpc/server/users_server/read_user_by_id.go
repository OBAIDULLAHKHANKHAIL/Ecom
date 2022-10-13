package main

import (
	"context"
	"fmt"
	"log"

	pb "obaid/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserByID(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	log.Printf("GetUserByID was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &userStruct{}
	filter := bson.M{"_id": oid}

	res := user_collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user with specified ID: %v", err),
		)
	}

	return assignUserValues(data), nil
}
