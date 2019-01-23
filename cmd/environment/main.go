package main

import (
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
	if err := stream.Send(&pb.ShipUpdate{Union: &pb.ShipUpdate_Control{Control: &pb.ControlState{Id: "0", ControlTypeValue: &pb.ControlState_Toggle{Toggle: true}}}}); err != nil {
		return err
	}
	
// 	for {
// 		in, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		key := serialize(in.Location)

// 		s.mu.Lock()
// 		s.routeNotes[key] = append(s.routeNotes[key], in)
// 		// Note: this copy prevents blocking other clients while serving this one.
// 		// We don't need to do a deep copy, because elements in the slice are
// 		// insert-only and never modified.
// 		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
// 		copy(rn, s.routeNotes[key])
// 		s.mu.Unlock()

// 		for _, note := range rn {
// 			if err := stream.Send(note); err != nil {
// 				return err
// 			}
// 		}
	// 	}
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
