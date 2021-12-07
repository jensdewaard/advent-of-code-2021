package day7

import (
	"sort"

	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	input :=
		strings.MapToInt(
			strings.ParseInt,
			strings.Split(
				files.ReadLines("day7/input")[0],
				",",
			),
		)
	sort.Ints(input)
	median := ints.Median(input)
	return CalculateFuelUsage(median, input)
}

func SolveB() int {
	return 0
}
