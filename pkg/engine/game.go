package engine

import(
	"github.com/LSFN/ode"
	"github.com/LSFN/lsfn/pkg/ship"
	pb "github.com/LSFN/lsfn/api/proto"
)

type Sensor interface {
	Update(ship *Ship)
	Read() (*pb.ShipUpdate_Sensor)
	Describe() (*pb.SensorDescription)
}

type Part interface {
	Register(ship *Ship)
}

type GPSSensor struct{
	id  string
	loc *pb.GalacticCoordinate
}


func (sensor *GPSSensor) Update(ship *Ship) {
	vec := ship.PhysicsBody.Position()
	sensor.loc = &pb.GalacticCoordinate{X: vec[0], Y: vec[1], Z: vec[2]}
}

func (sensor GPSSensor) Read() (*pb.ShipUpdate_Sensor) {
	return &pb.ShipUpdate_Sensor{Sensor:
		&pb.SensorState{Id: sensor.id, SensorTypeValue:
			&pb.SensorState_GalacticCoordinate{GalacticCoordinate:
				sensor.loc,
			},
		},
	}
}

func (sensor GPSSensor) Describe() (*pb.SensorDescription) {
	return &pb.SensorDescription{
		Id: sensor.id,
		Name: "Galactic Positioning System Location",
		SensorType: pb.SensorDescription_GalacticCoordinates,
	}
}


type GPS struct {
	Id string
}

func (g GPS) Register(ship *Ship) {
	ship.AddSensor(&GPSSensor{id: g.Id + ":sensor"})
}

type Ship struct {
	Description *ship.ShipDescription
	PhysicsBody *ode.Body
	Parts       []Part
	Sensors     []Sensor
}

func (s *Ship) AddSensor(sensor Sensor) {
	// Update before we add it so that it's value is never nil
	sensor.Update(s)
	s.Sensors = append(s.Sensors, sensor)
}

func (s *Ship) RegisterPart(part Part) {
	part.Register(s)
	s.Parts = append(s.Parts, part)
}

type Game struct {
	PhysicsWorld ode.World
	PhysicsSpace ode.Space
	Ships        []*Ship
}

func (s Ship) Describe() (*pb.ShipDescription) {
	sensorDescriptions := make([]*pb.SensorDescription, len(s.Sensors))
	for i, sensor := range s.Sensors {
		sensorDescriptions[i] = sensor.Describe()
	}
	
	return &pb.ShipDescription{
		Id: "0",
		Name: s.Description.Name,
		Sensors: sensorDescriptions,
	}
}

func NewGame() *Game {
	g := Game{}
	g.PhysicsWorld = ode.NewWorld()
	g.PhysicsSpace = ode.NilSpace().NewHashSpace()

	return &g
}

func (g *Game) AddShip(sd *ship.ShipDescription) {
	// Create all the physics objects
	size := ode.V3(sd.Physics.Width, sd.Physics.Length, sd.Physics.Height)

	mass := ode.NewMass()
	mass.SetBox(1, size)
	mass.Adjust(sd.Physics.Mass)

	collision := g.PhysicsSpace.NewBox(size)

	body := g.PhysicsWorld.NewBody()
	body.SetMass(mass)

	collision.SetBody(body)

	// Create the Ship struct
	s := Ship{
		Description: sd,
		PhysicsBody: &body,
	}
	g.Ships = append(g.Ships, &s)
	
	// Create default parts
	s.RegisterPart(GPS{Id: "0"})

	// FIXME Add little example force to push it
	body.AddForceAtPos([]float64{1, 0, 0}, ode.V3(0, 0, 0))
}
