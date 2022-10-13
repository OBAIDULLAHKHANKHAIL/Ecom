//go:build !test
//+build !test

package main

import (
	"context"
	"log"
	"net"

	pb "obaid/pb"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addr string = "0.0.0.0:50051"

func main() {
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	user_collection = client.Database("userdb").Collection("user")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}



/*
///////////////////////////
//        Terminal       //
///////////////////////////

protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import grpc/proto/users.proto

go build -o bin/user/server.exe ./grpc/server/users_server

go build -o bin/user/client.exe ./grpc/client/users_client



*/



// const (
// 	address = "localhost:50051"
// )

// func main() {

// 	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewUserManagementClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	var new_users = make(map[string]int32)
// 	new_users["Alice"] = 43
// 	new_users["Bob"] = 30
// 	for name, age := range new_users {
// 		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
// 		if err != nil {
// 			fmt.Println(err)
// 			log.Fatalf("could not create user: %v", err)
// 		}
// 		log.Printf(`User Details:
// NAME: %s
// AGE: %d
// ID: %d`, r.GetName(), r.GetAge(), r.GetId())
// 	}
// 	params := &pb.GetUsersParams{}
	// r, err := c.GetUsers(ctx, params) 
// 	if err != nil {
// 		log.Fatalf("Could not retreve users: %v", err)
// 	}
// 	log.Print("\nUSER LIST: \n")
// 	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())
// }
