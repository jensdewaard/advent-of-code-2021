package day1

import (
	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	lines := files.ReadLines("day1/input")
	depths := shared.Map(strings.ParseInt, lines)
	increases := shared.Zip(
		ints.LessThan,
		depths[:len(depths)-1],
		depths[1:],
	)
	return shared.CountIf(
		func(t bool) bool {
			return t
		}, increases,
	)
}

func SolveB() int {
	lines := files.ReadLines("day1/input")
	depths := shared.Map(strings.ParseInt, lines)
	windows := shared.Window(3, depths)
	summedWindows := shared.Map(ints.Sum, windows)
	increases := shared.Zip(
		ints.LessThan,
		summedWindows[:len(summedWindows)-1],
		summedWindows[1:],
	)
	return shared.CountIf(
		func(t bool) bool {
			return t
		}, increases,
	)
}
