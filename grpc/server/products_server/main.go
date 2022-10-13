//go:build !test
// +build !test

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

	product_collection = client.Database("productsdb").Collection("product")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterProductManagementServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}



/*
///////////////////////////
//        Terminal       //
///////////////////////////

protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import grpc/proto/*.proto

go build -o bin/user/server.exe ./grpc/server/users_server

go build -o bin/user/client.exe ./grpc/client/users_client



*/


