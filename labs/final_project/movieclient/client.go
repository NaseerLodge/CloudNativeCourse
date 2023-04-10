package main

import (
	"context"
	"fmt"

	pb "github.com/NaseerLodge/CloudNativeCourse/labs/final_project/movieapi"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("[::]:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	req := &pb.AdditionRequest{
		FirstNumber:  2,
		SecondNumber: 3,
	}

	res, err := client.Addition(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", res.Result)
}

/*
// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/NaseerLodge/CloudNativeCourse/labs/final_project/movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())
}
*/
