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

func (s *Server) UpdateProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	log.Printf("UpdateProduct was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := productstruct{
		Name:        in.Name,
		Price:       in.Price,
		Description: in.Description,
		Password:    in.Password,
		CreatedBy:   in.Createdby,
		CreatedAt:   in.Createdat,
		UpdatedBy:   in.Updatedby,
		UpdatedAt:   in.Updatedat,
	}
	
	res, err := product_collection.UpdateOne(
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
			"Cannot find Product with ID",
		)
	}

	return &pb.Product{
		Id:          oid.Hex(),
		Name:        in.Name,
		Price:       in.Price,
		Description: in.Description,
		Password:    in.Password,
		Createdby:   in.Createdby,
		Createdat:   in.Createdat,
		Updatedby:   in.Updatedby,
		Updatedat:   in.Updatedat,
	}, nil
}