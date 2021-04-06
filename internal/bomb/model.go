package bomb

type Port int

// This holds the different kinds of ports
const (
	PortDVID Port = iota + 1
	PortParallel
	PortPS2
	PortRJ45
	PortSerial
	PortStereoRCA
)

// Label holds the information for one label object
type Label struct {
	Text string
	Lit  bool
}

// Bomb holds the information for one bomb object
type Bomb struct {
	Config  BombConfig
	Strikes int // Number of strikes in the bomb
	Modules []interface{}
}

// BombConfig is used to explain the configuration of the bomb
type BombConfig struct {
	Labels    []Label // Labels within a bomb
	Batteries int     // Number of batteries in the bomb
	Ports     []Port  // Port
	Serial    string
}
