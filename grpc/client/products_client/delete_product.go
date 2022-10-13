package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func DeleteProduct(c pb.ProductManagementClient, id string) {
	log.Println("---DeleteProduct was invoked---")
	_, err := c.DeleteProduct(context.Background(), &pb.ProductID{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Product was deleted")
}