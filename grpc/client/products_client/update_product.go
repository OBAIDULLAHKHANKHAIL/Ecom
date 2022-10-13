package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func UpdateProduct(c pb.ProductManagementClient, id string) {
	log.Println("---UpdateProduct was invoked---")
	product := &pb.Product{
		Id:       id,
		Name: "obaid",
	}

	_, err := c.UpdateProduct(context.Background(), product)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("User was updated")
}