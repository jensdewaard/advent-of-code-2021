package day6

import (
	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	fish := shared.Map(
		strings.ParseInt,
		strings.Split(files.ReadLines("day6/input")[0], ","),
	)
	return FishAfterDaysSlice(80, fish)
}

func SolveB() int {
	fish := shared.Map(
		strings.ParseInt,
		strings.Split(files.ReadLines("day6/input")[0], ","),
	)
	ages := []Fish{0, 1, 2, 3, 4, 5, 6, 7, 8}
	offspringPerAge := shared.Map(
		func(f Fish) int {
			return FishAfterDaysSingle(256, f)
		},
		ages,
	)
	offspringPerFish := shared.Map(
		func(f Fish) int {
			return offspringPerAge[f]
		},
		fish,
	)
	return ints.Sum(offspringPerFish)
}
