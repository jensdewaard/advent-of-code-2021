package day2

import (
	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
)

func SolveA() int {
	lines := files.ReadLines("day2/input")
	commands := shared.Map(ParseCommand, lines)
	end := shared.Fold(Location{0, 0}, Move, commands)
	return end.Depth * end.Horizontal
}
func SolveB() int {
	lines := files.ReadLines("day2/input")
	commands := shared.Map(ParseCommand, lines)
	end := shared.Fold(LocationWithAim{0, 0, 0}, MoveWithAim, commands)
	return end.Depth * end.Horizontal
}
