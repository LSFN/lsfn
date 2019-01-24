package ship

type Physics struct {
	Width  uint `yaml:"width"`
	Height uint `yaml:"height"`
	Length uint `yaml:"length"`
}

type Part interface {
	name() string
}

type Ship struct {
	Name    string  `yaml:"name"`
	Physics Physics `yaml:"physics"`
	Parts   []Part  `yaml:"parts"`
}
