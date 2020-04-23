package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/mpedrozoduran/gogrpcserver/timeproto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	PORT = 9091
)

func main() {
	endpoint := "localhost:9091"
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterTimeManagerHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalf("Error when trying to register handler: %v", err)
	}
	err = http.ListenAndServe("localhost:8081", mux)
	if err != nil {
		log.Fatalf("Error when starting listener: %v", err)
	}
}
