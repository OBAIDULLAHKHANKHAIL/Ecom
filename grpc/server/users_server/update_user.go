package main

import (
	"context"
	"log"

	pb "obaid/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("UpdateUser was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := userStruct{
		Name:     in.Name,
		User:     in.User,
		Password: in.Password,
	}

	res, err := user_collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find user with ID",
		)
	}

	return &pb.User{
		Id:       oid.Hex(),
		Name:     in.Name,
		User:     in.User,
		Password: in.Password,
	}, nil
}
