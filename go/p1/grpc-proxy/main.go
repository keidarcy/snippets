package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "example.com/at/api/v1"
)

func main() {
	var grpcServerEndpoint = "localhost:8080"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterActivity_LogHandlerFromEndpoint(context.Background(),
		mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Listening on port 8081")
	port := ":8081"
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
