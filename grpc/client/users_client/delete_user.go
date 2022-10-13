package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func DeleteUser(c pb.UserManagementClient, id string) {
	log.Println("---DeleteUser was invoked---")
	_, err := c.DeleteUser(context.Background(), &pb.UserID{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("User was deleted")
}