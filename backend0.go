package main

import (
	srch "TP4/search"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"log"
	 "TP4/rand"
	"net"
	"time"
)

type searchServiceClient0 struct {
}

func (c *searchServiceClient0) MultipleSearch(ctx context.Context, in *srch.Request, opts ...grpc.CallOption) (srch.SearchService_MultipleSearchClient, error) {
	return nil,nil
}

type server0 struct {
}

func (s server0) Search(ctx context.Context, req *srch.Request) (*srch.Result, error) {
	if tr, ok := trace.FromContext(ctx); ok { // HL
		tr.LazyPrintf("Request: %s",(*req).Query) // HL
	}
	d := rand.RandomDuration(10 * time.Second)
	time.Sleep(d)
	return &srch.Result{
		Title: fmt.Sprintf("result for [%s] from backend0", req.Query),
	}, nil
}

func (s server0) MultipleSearch(*srch.Request, srch.SearchService_MultipleSearchServer) error {
	panic("implement me")
}

func main() {

	serverAddress := fmt.Sprintf(":%d", 5002)
	fmt.Printf("Server starts at %s...\n", serverAddress)

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Server failed to listen: %v", err)
	}
	// create gRPC server
	g := grpc.NewServer()
	// register gRPC search service server which must implement the function Search
	srch.RegisterSearchServiceServer(g,new(server0))
	// launch gRPC server
	g.Serve(lis)
}
