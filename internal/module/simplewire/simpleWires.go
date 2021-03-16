package simplewire

import (
	"math/rand"
	"strconv"

	"github.com/Ao-Re/ktane-cli/internal/bomb"
)

type WireColor int

const (
	Yellow WireColor = iota + 1
	Red
	Blue
	Black
	White
)

type SimpleWire struct {
	OddSerial bool
	NumWires  int
	Wires     []WireColor
}

func (sw *SimpleWire) Init(config bomb.BombConfig) {
	lastSerial := string(config.Serial[len(config.Serial)-1])
	num, err := strconv.Atoi(lastSerial)
	if err == nil {
		sw.OddSerial = num%2 == 0
	}

	sw.Wires = generate()
	sw.NumWires = len(sw.Wires)
}

func generate() []WireColor {
	numWires := rand.Intn(5) + 1

	wires := make([]WireColor, numWires)

	for i := 0; i < numWires; i++ {
		wires[i] = WireColor(rand.Intn(5) + 1)
	}

	return wires
}

func (sw *SimpleWire) CheckLogic(idx int) bool {
	solution := sw.SolveSimpleWires()
	return idx == solution
}

func (sw *SimpleWire) SolveSimpleWires() int {
	colors := make(map[WireColor]int)

	for _, wire := range sw.Wires {
		colors[wire]++
	}

	switch len(sw.Wires) {
	case 3:
		if colors[Red] == 0 {
			return 1
		} else if sw.Wires[len(sw.Wires)-1] == White {
			return 2
		} else if colors[Blue] > 1 {
			for i := len(sw.Wires) - 1; i >= 0; i-- {
				if sw.Wires[i] == Blue {
					return i
				}
			}
		}
	case 4:
		if colors[Red] >= 1 && sw.OddSerial {
			for i := len(sw.Wires) - 1; i >= 0; i-- {
				if sw.Wires[i] == Red {
					return i
				}
			}
		} else if sw.Wires[len(sw.Wires)-1] == Yellow && colors[Red] == 0 {
			return 0
		} else if colors[Yellow] > 1 {
			return 3
		}
		return 1
	case 5:
		if sw.Wires[len(sw.Wires)-1] == Black && sw.OddSerial {
			return 3
		} else if colors[Red] == 1 && colors[Yellow] > 1 {
			return 0
		} else if colors[Black] == 0 {
			return 1
		}
		return 0
	case 6:
		if colors[Yellow] == 0 && sw.OddSerial {
			return 2
		} else if colors[Yellow] == 1 && colors[White] > 1 {
			return 3
		} else if colors[Red] == 0 {
			return 5
		}
		return 3
	}

	return -1
}
