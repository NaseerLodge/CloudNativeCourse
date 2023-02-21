// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"home/naseer/Desktop/cloud_native/labs/lab5/movieapi"
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

	ret, err := c.SetMovieInfo(ctx, &movieapi.MovieData{Title: "Harry Potter", Year: 2001, Director: "Chris Columbus", Cast: []string{"Daniel Radcliffe, Rupert Grint, Emma Watson"}})

	if err != nil {
		log.Printf("SetMovieInfo status: %v", ret)
		log.Fatalf("could not set movie info: %v", err)
	}

	r2, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: "Harry Potter"})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for Harry Potter: %d %s %v", r2.GetYear(), r2.GetDirector(), r2.GetCast())
}
