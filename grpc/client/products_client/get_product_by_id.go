package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func GetProductByID(c pb.ProductManagementClient, id string) *pb.Product {
	log.Println("---GetProductByID was invoked---")

	req := &pb.ProductID{Id: id}
	res, err := c.GetProductByID(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("Product was read: %v\n", res)
	return res
}