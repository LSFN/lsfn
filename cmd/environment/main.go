package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/LSFN/lsfn/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"

	"github.com/LSFN/lsfn/pkg/ship"
	"github.com/LSFN/lsfn/pkg/engine"
	"github.com/LSFN/ode"
)

const (
	port = ":50051"
)

type server struct{}

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
}

func loadShip(filename string) (*ship.ShipDescription, error) {
	log.Printf("Loading ship from file %s", filename)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := ship.ShipDescription{}
	err = yaml.Unmarshal(yamlFile, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func simulate(game *engine.Game) {
	log.Print("Starting simulation")
	for {
		game.PhysicsWorld.Step(0.1)
		for _, ship := range game.Ships {
			for _, part := range ship.Parts {
				part.Step(ship)
			}
		}
	}
}

func main() {
	ode.Init(0, ode.AllAFlag)
	
	game := engine.NewGame()
	sh, err := loadShip("configs/ships/resolution-of-dawn.yaml")
	if err != nil {
		log.Fatalf("failed to load ship file: %v", err)
	}
	log.Printf("Loaded ship %v", sh)
	game.AddShip(sh)

	
	simulate(game)
	
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
