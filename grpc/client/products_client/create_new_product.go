package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func CreateNewProduct(c pb.ProductManagementClient) string {
	log.Println("---CreateNewProduct was invoked---")

	product := &pb.NewProduct{
		Name: "obaid",
	}

	res, err := c.CreateNewProduct(context.Background(), product)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("New Product has been created: %v\n", res)
	return res.Id
}


