package day2

import (
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
)

func SolveA() int {
	lines := files.ReadLines("day2/input")
	commands := MapStringsToCommands(lines)
	start := Location{0, 0}
	end := FoldCommands(start, commands)
	return end.Depth * end.Horizontal
}
func SolveB() int {
	lines := files.ReadLines("day2/input")
	commands := MapStringsToCommands(lines)
	start := LocationWithAim{0, 0, 0}
	end := FoldCommandsWithAim(start, commands)
	return end.Depth * end.Horizontal
}
