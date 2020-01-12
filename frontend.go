package main

import (
	"TP4/rand"
	srch "TP4/search"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type server struct{}

//Constant de connexion
const (
	BACKEND1 = ":5001"
	BACKEND0 = ":5002"
)

var services []srch.SearchServiceClient // list des SearchService (back1 && back0 )
var wg sync.WaitGroup

func (s *server) MultipleSearch(req *srch.Request, stream srch.SearchService_MultipleSearchServer) error {

	wg.Add(len(services)) //ajout nbService au WaitGroup

	for i := 0; i < len(services); i++ { // parcours de mes SearchService
		go func(service srch.SearchServiceClient, req *srch.Request) {
			res, err := service.Search(context.Background(), req)
			fmt.Printf("res :%v\n", res)
			stream.Send(res) // envoie du res dans le stream
			if err != nil {
				stream.SendMsg(err) //envoie des erreurs
			}
			defer wg.Done() // fin de la goroutine
		}(services[i], req)

	}

	wg.Wait()// attente de la fin des goroutines
	stream.Send(nil) // permet au client de sortir de la lecture du stream
	return nil
}

// Search sleeps for a random interval then returns a string
// identifying the query and this backend.
func (s *server) Search(ctx context.Context, req *srch.Request) (*srch.Result, error) {
	if tr, ok := trace.FromContext(ctx); ok { // HL
		tr.LazyPrintf("Request: %s", (*req).Query) // HL
	}
	d := rand.RandomDuration(100 * time.Millisecond)
	time.Sleep(d)
	return &srch.Result{
		Title: fmt.Sprintf("result for [%s] from server", req.Query),
	}, nil
}

func main() {
	// Partie client
	go connBackEnd1() //connect au backend1
	go connBackend0() // connect au backend0

	//Partie serveur
	serverAddress := fmt.Sprintf(":%d", 4000)
	fmt.Printf("Server starts at %s...\n", serverAddress)

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Server failed to listen: %v", err)
	}
	// create gRPC server
	g := grpc.NewServer()
	// register gRPC search service server which must implement the function Search
	srch.RegisterSearchServiceServer(g, new(server))
	// launch gRPC server
	g.Serve(lis)

}

func connBackend0() {
	conn, err := grpc.Dial(BACKEND0, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := srch.NewSearchServiceClient(conn)
	services = append(services, client)

}

func connBackEnd1() {
	conn2, err := grpc.Dial(BACKEND1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client2 := srch.NewSearchServiceClient(conn2)
	services = append(services, client2)

}
