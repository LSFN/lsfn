package engine

import(
	"github.com/LSFN/ode"
	"github.com/LSFN/lsfn/pkg/ship"
)

type Ship struct {
	Description *ship.ShipDescription
	PhysicsBody *ode.Body
}

type Game struct {
	PhysicsWorld ode.World
	PhysicsSpace ode.Space
	Ships        []*Ship
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
	s := Ship{Description: sd, PhysicsBody: &body}
	g.Ships = append(g.Ships, &s)

	// FIXME Add little example force to push it
	body.AddForceAtPos([]float64{1, 0, 0}, ode.V3(0, 0, 0))
}
