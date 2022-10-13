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

func (s *Server) CreateNewProduct(ctx context.Context, in *pb.NewProduct) (*pb.Product, error) {
	log.Printf("CreateNewProduct was invoked with %v\n", in)

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

	res, err := product_collection.InsertOne(ctx, data)

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
