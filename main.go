package main

import (
	"context"
	"log"
	"net"

	"github.com/roysyboy/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	log.Println("Generated server registrar")
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	log.Println("running server, listening to: " + lis.Addr().String())
	serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve : %s", err)
	}
	// log.Println("running server, listening to: " + lis.Addr().String())
}
