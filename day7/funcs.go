package day7

import (
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
)

func CalculateFuelUsage(position int, crabs []int, fuel func(int, int) int) int {
	dists := ints.MapToInt(
		func(i int) int {
			return fuel(position, i)
		},
		crabs,
	)
	return ints.Sum(dists)
}

func CrabFuel(from, to int) int {
	return Gauss(ints.Diff(from, to))
}

func Gauss(i int) int {
	return (i * (i + 1)) / 2
}
