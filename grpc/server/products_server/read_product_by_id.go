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

func (s *Server) GetProductByID(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	log.Printf("GetProductByID was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &productstruct{}
	filter := bson.M{"_id": oid}

	res := product_collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find product with specified ID: %v", err),
		)
	}

	return assignProductValues(data), nil
}