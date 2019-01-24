package ship

// Information about the 3D rigid body that defines the ship in phyiscs
// simulations
type Physics struct {
	Width  uint `yaml:"width"`
	Height uint `yaml:"height"`
	Length uint `yaml:"length"`
}

// A "part" is something that's on the ship. Generally this will expose
// a number of controls and sensors, along with some functionality that
// will impact the game.
type Part interface {
	name() string
}

// This is the game engine's internal view of a ship. With all the information
// needed to simulate it.
type Ship struct {
	Name    string  `yaml:"name"`
	Physics Physics `yaml:"physics"`
	Parts   []Part  `yaml:"parts"`
}
