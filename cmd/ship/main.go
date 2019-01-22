package main

import (
	"time"
	"context"
	"log"

	"google.golang.org/grpc"
	pb "github.com/LSFN/lsfn/api/proto"
)

const (
	address = "localhost:50051"
)

type server struct{}

func (s *server) Join(ctx context.Context, in *pb.JoinServer) (*pb.Welcome, error) {
	log.Printf("Received message to join")
	return &pb.Welcome{Ship: &pb.ShipDescription{Id: "1", Name: "Resolution to Finish"}}, nil
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLobbyClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Join(ctx, &pb.JoinServer{})
	if err != nil {
		log.Fatalf("could not join server: %v", err)
	}
	log.Printf("Joined server, assinged ship %s", r.Ship.Name)
}
