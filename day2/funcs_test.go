package day2

import (
	"testing"

	"github.com/jensdewaard/advent-of-code-2021/shared"
)

func TestMove(t *testing.T) {
	tests := []struct {
		start    Location
		command  Command
		expected Location
	}{
		{
			Location{0, 0},
			Command{Up, 5},
			Location{0, -5},
		},
		{
			Location{10, 5},
			Command{Forward, 3},
			Location{13, 5},
		},
	}
	for _, tt := range tests {
		actual := Move(tt.start, tt.command)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func TestFoldCommands(t *testing.T) {
	tests := []struct {
		start    Location
		commands []Command
		expected Location
	}{
		{
			Location{0, 0},
			[]Command{
				{Forward, 5},
				{Down, 5},
				{Forward, 8},
				{Up, 3},
				{Down, 8},
				{Forward, 2},
			},
			Location{15, 10},
		},
	}
	for _, tt := range tests {
		actual := shared.Fold(tt.start, Move, tt.commands)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func TestFoldCommandsWithAim(t *testing.T) {
	tests := []struct {
		start    LocationWithAim
		commands []Command
		expected LocationWithAim
	}{
		{
			LocationWithAim{0, 0, 0},
			[]Command{
				{Forward, 5},
				{Down, 5},
				{Forward, 8},
				{Up, 3},
				{Down, 8},
				{Forward, 2},
			},
			LocationWithAim{15, 60, 10},
		},
	}
	for _, tt := range tests {
		actual := shared.Fold(tt.start, MoveWithAim, tt.commands)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
