package main

import (
	"io"
	"net"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/LSFN/lsfn/api/proto"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) Join(ctx context.Context, in *pb.JoinServer) (*pb.Welcome, error) {
	log.Printf("Received message to join")
	return &pb.Welcome{Ship: &pb.ShipDescription{Id: "1", Name: "Resolution to Finish"}}, nil
}

func (s *server) Command(stream pb.Lobby_CommandServer) error {
	log.Print("Bidirectional stream opened to command ship")
	if err := stream.Send(&pb.ShipUpdate{Union: &pb.ShipUpdate_Control{Control: &pb.ControlState{Id: "0", ControlTypeValue: &pb.ControlState_Toggle{Toggle: true}}}}); err != nil {
		log.Printf("Failed to send ship update %v", err)
		return err
	}
	
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Print("Got EOF from ship, closing.")
			return nil
		}
		if err != nil {
			log.Printf("Failed to recive from ship: %v", err)
			return err
		}
		log.Printf("Got %v from ship", in)
	}
	log.Print("Closing bidirectional stream with ship")
	return nil
}

func main() {
	log.Printf("Staring server on %v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLobbyServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}
