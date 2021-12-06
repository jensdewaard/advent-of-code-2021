package day6

import (
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	fish := strings.MapToInt(
		strings.ParseInt,
		strings.Split(files.ReadLines("day6/input")[0], ","),
	)
	return FishAfterDays(80, fish)
}

func SolveB() int {
	fish := strings.MapToInt(
		strings.ParseInt,
		strings.Split(files.ReadLines("day6/input")[0], ","),
	)
	return FishAfterDays(256, fish)
}
