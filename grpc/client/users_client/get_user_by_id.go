package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func GetUserByID(c pb.UserManagementClient, id string) *pb.User {
	log.Println("---GetUserByID was invoked---")

	req := &pb.UserID{Id: id}
	res, err := c.GetUserByID(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("User was read: %v\n", res)
	return res
}