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

type server struct {
	Game *engine.Game
}

func NewServer() (*server) {
	return &server{Game: engine.NewGame()}
}

func (s *server) Join(ctx context.Context, in *pb.JoinServer) (*pb.Welcome, error) {
	log.Printf("Received message to join")
	return &pb.Welcome{Ship: s.Game.Ships[0].Describe()}, nil
}

func (s *server) Command(stream pb.Lobby_CommandServer) error {
	log.Print("Bidirectional stream opened to command ship")

	updates := make(chan *pb.ShipUpdate, 10000)
	stop := make(chan int)
	go func() {
		log.Printf("Make goroutine to take ship updates")
		// TODO close this channel correctly
		for i := range updates {
			if err := stream.Send(i); err != nil {
				log.Printf("Failed to send ship update %v", err)
				return
			}
		}
	}()

	// TODO add this channel to the game in a thread-safe way
	s.Game.Ships[0].Updates = updates
	s.Game.Ships[0].Stop = stop

	defer func() {
		log.Print("Trigger close")
		close(stop)
	}()
	
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
			for _, sensor := range ship.Sensors {
				sensor.Update(ship)
				log.Printf("Foo")
				select {
				case <-ship.Stop:
					log.Print("close triggered")
					close(ship.Updates)
					ship.Updates = nil
				default:
					if ship.Updates != nil {
						u := &pb.ShipUpdate{Union: sensor.Read()}
						ship.Updates <- u

					}
				}
			}
		}
	}
}

func main() {
	ode.Init(0, ode.AllAFlag)

	server := NewServer()

	sh, err := loadShip("configs/ships/resolution-of-dawn.yaml")
	if err != nil {
		log.Fatalf("failed to load ship file: %v", err)
	}
	log.Printf("Loaded ship %v", sh)
	server.Game.AddShip(sh)

	go simulate(server.Game)

	log.Printf("Staring server on %v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLobbyServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
