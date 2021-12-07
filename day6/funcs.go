package day6

import (
	"fmt"

	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
)

func Age(f Fish) []Fish {
	if f == 0 {
		return []Fish{6, 8}
	}
	return []int{f - 1}
}

func FishAfterDaysSlice(days int, fs []Fish) int {
	perFish := ints.MapToInt(
		func(f Fish) int {
			return FishAfterDaysSingle(days, f)
		},
		fs,
	)
	return ints.Sum(perFish)
}

var memo = make(map[string]int, 0)

func FishAfterDaysSingle(days int, f Fish) int {
	if days == 0 {
		return 1
	}
	key := fmt.Sprintf("(%d, %d)", days, f)
	_, ok := memo[key]
	if ok {
		return memo[key]
	}
	tomorrow := Age(f)
	result := FishAfterDaysSlice(days-1, tomorrow)
	memo[key] = result
	return result
}
