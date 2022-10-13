package main

import (
	"log"

	pb "obaid/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewProductManagementClient(conn)

	CreateNewProduct(c)

	// id := createBlog(c)
	// readBlog(c, id)						 // valid
	// // readBlog(c, "aNonExistingID")	 // not valid
	// updateBlog(c, id)
	// listBlog(c)
	// deleteBlog(c, id)
}