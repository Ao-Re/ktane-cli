package simplewire

import "testing"

func TestSimpleWire_SolveSimpleWires(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		sw   *SimpleWire
		want int
	}{
		{
			name: "Yellow Yellow Black",
			sw: &SimpleWire{
				OddSerial: false,
				NumWires:  3,
				Wires:     []WireColor{Yellow, Yellow, Black},
			},
			want: 1,
		},
		{
			name: "Yellow Red White",
			sw: &SimpleWire{
				OddSerial: false,
				NumWires:  3,
				Wires:     []WireColor{Yellow, Red, White},
			},
			want: 2,
		},
		{
			name: "Blue Red Blue",
			sw: &SimpleWire{
				OddSerial: false,
				NumWires:  3,
				Wires:     []WireColor{Blue, Red, Blue},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sw.SolveSimpleWires(); got != tt.want {
				t.Errorf("SimpleWire.SolveSimpleWires() = %v, want %v", got, tt.want)
			}
		})
	}
}
