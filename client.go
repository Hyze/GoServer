package main

import (
	srch "TP4/search"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := srch.NewSearchServiceClient(conn)
	req := &srch.Request{Query: "test"}
	res, err := client.MultipleSearch(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := res.Recv() // parcours le stream
		if err == io.EOF {		// j'ai pas reussis a lui envoyer de EOF
			res.CloseSend()
			return
		}
		if msg == nil{  // arrete la lecture du stream
			return
		}
		fmt.Printf("\nRes : %v\n", msg.Title)
	}

	defer conn.Close() //ferme la connection avec le front

}
