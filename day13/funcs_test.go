package day13

import (
	"testing"

	"github.com/jensdewaard/advent-of-code-2021/shared/grids"
)

func Test_MirrorPositionHorizontally(t *testing.T) {
	tests := []struct {
		p        grids.Position
		l        int
		expected grids.Position
	}{
		{grids.Position{X: 5, Y: 5}, 3, grids.Position{X: 1, Y: 5}},
		{grids.Position{X: 6, Y: 0}, 5, grids.Position{X: 4, Y: 0}},
	}
	for _, tt := range tests {
		actual := MirrorPositionHorizontally(tt.p, tt.l)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_MirrorPositionVertically(t *testing.T) {
	tests := []struct {
		p        grids.Position
		l        int
		expected grids.Position
	}{
		{grids.Position{X: 0, Y: 14}, 7, grids.Position{X: 0, Y: 0}},
	}
	for _, tt := range tests {
		actual := MirrorPositionVertically(tt.p, tt.l)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
