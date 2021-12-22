package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	gw "questionService/libs"
	question "questionService/libs"
	"questionService/methods"
)

const (
	grpcPort = "0.0.0.0:50051"
	restPort = "0.0.0.0:8000"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatalf("Error in starting server %v", err)
	}

	s := grpc.NewServer()
	question.RegisterQuestionServiceServer(s, &methods.Server{})
	reflection.Register(s)

	log.Printf("Listening grpc at : %v", lis.Addr())
	log.Printf("Listening rest at : %v", restPort)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to Serve %v", err)
		}
	}()

	// wait for ctrl c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	log.Println("Stopping Server...")
	s.Stop()
	log.Println("Stopping Listener...")
	lis.Close()
	log.Println("Server Stopped.")
}

func main() {
	log.Println("Starting Question Service...")
	// Better logging with file names
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Thread for grpc gateway REST Server
	go func() {
		// mux
		mux := runtime.NewServeMux()
		// register
		err := gw.RegisterQuestionServiceHandlerServer(context.Background(), mux, &methods.Server{})
		if err != nil {
			panic(err.Error())
		}

		// http server
		log.Fatalln(http.ListenAndServe(restPort, mux))
	}()

	StartGrpcServer()
}
