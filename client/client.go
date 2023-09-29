package client

import (
	"context"
	"fmt"
	"log"

	"github.com/roysyboy/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:8080" // Change this to the actual server address

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := invoicer.NewInvoicerClient(conn)

	// Prepare the request message
	request := &invoicer.CreateRequest{
		// Fill in the request fields as needed
	}

	// Call the gRPC service method
	response, err := client.Create(context.Background(), request)
	if err != nil {
		// Handle errors, including gRPC-specific errors
		// if status.Code(err) == status.Code(codes.Unimplemented) {
		// 	log.Fatalf("Method Create not implemented by the server: %v", err)
		// }
		log.Fatalf("RPC error: %v", err)
	}

	// Process the response
	fmt.Printf("Response: %v\n", response)
}
