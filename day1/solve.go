package day1

import (
	"github.com/jensdewaard/advent-of-code-2021/shared/bools"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	lines := files.ReadLines("day1/input")
	depths := strings.MapToInt(strings.ParseInt, lines)
	increases := ints.ZipToBool(
		ints.LessThan,
		depths[:len(depths)-1],
		depths[1:],
	)
	return bools.CountTrue(increases)
}

func SolveB() int {
	lines := files.ReadLines("day1/input")
	depths := strings.MapToInt(strings.ParseInt, lines)
	windows := ints.Window(3, depths)
	summedWindows := ints.MapSlicesToInt(ints.Sum, windows)
	increases := ints.ZipToBool(
		ints.LessThan,
		summedWindows[:len(summedWindows)-1],
		summedWindows[1:],
	)
	return bools.CountTrue(increases)
}
