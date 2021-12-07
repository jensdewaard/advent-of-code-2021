package day7

import (
	"sort"

	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	crabs := strings.MapToInt(
		strings.ParseInt,
		strings.Split(
			files.ReadLines("day7/input")[0],
			",",
		),
	)
	sort.Ints(crabs)
	median := ints.Median(crabs)
	return CalculateFuelUsage(median, crabs, ints.Diff)
}

func SolveB() int {
	crabs := strings.MapToInt(
		strings.ParseInt,
		strings.Split(
			files.ReadLines("day7/input")[0],
			",",
		),
	)
	sort.Ints(crabs)
	position := ints.Average(crabs)
	return CalculateFuelUsage(position, crabs, CrabFuel)
}
