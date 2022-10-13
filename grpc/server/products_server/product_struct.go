package main

import (
	pb "obaid/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type productstruct struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Name        string                 `bson:"name"`
	Price       string                 `bson:"price"`
	Description string                 `bson:"description"`
	Password    string                 `bson:"password"`
	CreatedBy   string                 `bson:"createdBy"`
	CreatedAt   *timestamppb.Timestamp `bson:"createdAt"`
	UpdatedBy   string                 `bson:"updatedBy"`
	UpdatedAt   *timestamppb.Timestamp `bson:"updatedAt"`
}

func assignProductValues(data *productstruct) *pb.Product {
	return &pb.Product{
		Id:          data.ID.Hex(),
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Password:    data.Password,
		Createdby:   data.CreatedBy,
		Createdat:   data.CreatedAt,
		Updatedby:   data.UpdatedBy,
		Updatedat:   data.UpdatedAt,
	}
}
