package game

import(
	"github.com/LSFN/ode"
	"github.com/LSFN/lsfn/pkg/ship"
)

type Game struct {
	PhysicsWorld ode.World
	PhysicsSpace ode.Space
	ships        []*ship.Ship
}

func NewGame() *Game {
	g := Game{}
	g.PhysicsWorld = ode.NewWorld()
	g.PhysicsSpace = ode.NilSpace().NewHashSpace()

	return g
}

func (g *Game) AddShip(ship ship.Ship) (*ode.Body) {
	size := ode.V3(ship.Physics.Width, ship.Physics.Length, ship.Physics.Height)

	mass := ode.NewMass()
	mass.SetBox(1, size)
	mass.Adjust(ship.Physics.Mass)

	collision := g.PhysicsSpace.NewBox(size)

	body := g.PhysicsWorld.NewBody()
	body.SetMass(mass)

	collision.setBody(body)

	return &body
}
