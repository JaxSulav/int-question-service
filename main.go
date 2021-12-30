package main

import (
	"google.golang.org/grpc"
	"log"
	"questionService/config"
	"questionService/libs/auth"
	"questionService/server"
)

func ConnectAuthClient() auth.AuthClient {
	// Dial to the server address, the connection given by dial will be used to create a new calculator client
	conn, err := grpc.Dial(config.AuthAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to auth server %v \n", err)
	}
	return auth.NewAuthClient(conn)
}

func main() {
	log.Println("Starting Question Service...")

	// Better logging with file names
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Client for auth gRPC
	server.AuthClient = ConnectAuthClient()

	// Thread for grpc gateway REST Server
	go func() {
		server.StartGatewayServer()
	}()

	// gRPC Server
	server.StartGrpcServer()
}
