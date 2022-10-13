package main

import (
	"context"
	"log"

	pb "obaid/pb"
)

func CreateNewUser(c pb.UserManagementClient) string {
	log.Println("---CreateNewUser was invoked---")

	user := &pb.NewUser{
		Name: "obaid",
		User: "regular",
		Password: "1234",
	}

	res, err := c.CreateNewUser(context.Background(), user)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("New user has been created: %v\n", res)
	return res.Id
}


