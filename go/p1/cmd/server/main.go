package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"example.com/at/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	api "example.com/at/api/v1"
)

// func main() {
// 	println("Starting")
// 	s := server.NewHTTPServer(":8080")
// 	s.ListenAndServe()
// }

func main() {
	// log.Println("Starting listening on port 8080")
	// port := ":8080"

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Printf("Listening on %s", port)
	// srv := server.NewGRPCServer()
	// GRPC Server
	grpcServer, srv := server.NewGRPCServer()

	// Rest Server
	mux := runtime.NewServeMux()
	err := api.RegisterActivity_LogHandlerServer(context.Background(), mux, &srv)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// reflection.Register(grpcServer)

	log.Println("Starting listening on port 8080")
	// err = http.ListenAndServe(":8080", mux)
	err = http.ListenAndServe(":8080", grpcHandlerFunc(*grpcServer, mux))

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// if err := srv.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}

func grpcHandlerFunc(grpcServer grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(
			r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("GRPC")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("REST")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

// 	api "example.com/at/api/v1"
// 	"example.com/at/internal/server"
// )

// func main() {

// 	// GRPC Server
// 	_, srv := server.NewGRPCServer()

// 	// Rest Server
// 	mux := runtime.NewServeMux()
// 	err := api.RegisterActivity_LogHandlerServer(context.Background(), mux, &srv)
// 	if err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// 	log.Println("Starting listening on port 8080")
// 	err = http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
