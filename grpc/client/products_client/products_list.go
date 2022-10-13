package main

import (
	"context"
	"io"
	"log"

	pb "obaid/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func GetProducts(c pb.ProductManagementClient) {
	log.Println("---GetProducts was invoked---")
	stream, err := c.GetProducts(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling GetProducts: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
