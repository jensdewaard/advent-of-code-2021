package day7

import "github.com/jensdewaard/advent-of-code-2021/shared/ints"

func CalculateFuelUsage(position int, crabs []int) int {
	dists := ints.MapToInt(
		func(i int) int {
			return ints.Diff(position, i)
		},
		crabs,
	)
	return ints.Sum(dists)
}
