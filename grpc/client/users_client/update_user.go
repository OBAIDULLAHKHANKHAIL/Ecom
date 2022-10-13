package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func UpdateUser(c pb.UserManagementClient, id string) {
	log.Println("---updateUser was invoked---")
	user := &pb.User{
		Id:       id,
		Name: "obaid",
		User: "regular",
		Password: "12345",
	}

	_, err := c.UpdateUser(context.Background(), user)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("User was updated")
}