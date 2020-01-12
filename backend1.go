package main

import (
	"TP4/rand"
	srch "TP4/search"
	"context"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server1 struct{}

func (s server1) Search(ctx context.Context, req *srch.Request) (*srch.Result, error) {
	if tr, ok := trace.FromContext(ctx); ok { // HL
		tr.LazyPrintf("Request: %s",(*req).Query) // HL
	}
	d := rand.RandomDuration(5 * time.Second)
	time.Sleep(d)
	return &srch.Result{
		Title: fmt.Sprintf("result for [%s] from backend1", req.Query),
	}, nil
}

func (s server1) MultipleSearch(*srch.Request, srch.SearchService_MultipleSearchServer) error {
	panic("implement me")
}

func main(){

	serverAddress := fmt.Sprintf(":%d", 5001)
	fmt.Printf("Server starts at %s...\n", serverAddress)

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Server failed to listen: %v", err)
	}
	// create gRPC server
	g := grpc.NewServer()
	// register gRPC search service server which must implement the function Search
	srch.RegisterSearchServiceServer(g, new(server1))
	// launch gRPC server
	g.Serve(lis)

}


