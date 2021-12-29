package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	grpcPort = "0.0.0.0:50052"
	restPort = "0.0.0.0:8000"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("--> unary interceptor: %v", info.FullMethod)
	return handler(ctx, req)
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatalf("Error in starting server %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(AuthInterceptor),
	)
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

func dummyAuth(token string) error {
	log.Printf("Token: %v", token)
	return status.Errorf(codes.Unauthenticated, "Could not authenticate...")
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
		gatewayAddr := "0.0.0.0:8000"
		gwServer := &http.Server{
			Addr: gatewayAddr,
			// Handle authentication through auth interceptor
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.Println("->> Verifying Authentication")
				bearer := r.Header.Get("Authorization")
				// Call grpc auth server
				err := dummyAuth(bearer)
				if err == nil {
					mux.ServeHTTP(w, r)
					return
				}
				// Case: Invalid auth token, write message to response writer object
				w.WriteHeader(http.StatusUnauthorized)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					log.Printf("Error weiting to response writer: %v", err)
					return
				}
			}),
		}
		// http server
		log.Fatalln(gwServer.ListenAndServe())
	}()

	StartGrpcServer()
}
