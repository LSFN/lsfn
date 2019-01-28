package ship

// The Physics struct provides information about the 3D rigid body that defines
// the ship in physics simulations
type Physics struct {
	Width  float64 `yaml:"width"`
	Height float64 `yaml:"height"`
	Length float64 `yaml:"length"`
	Mass   float64 `yaml:"mass"`
}

// The Ship struct is the game engine's internal view of a ship. With all the
// information needed to simulate it.
type ShipDescription struct {
	Name    string  `yaml:"name"`
	Physics Physics `yaml:"physics"`
}
