package day6

import "github.com/jensdewaard/advent-of-code-2021/shared/ints"

func Age(f Fish) []Fish {
	if f == 0 {
		return []Fish{6, 8}
	}
	return []int{f - 1}
}

func FishAfterDays(days int, fs []Fish) int {
	for i := 0; i < days; i++ {
		fs = ints.FlatMapToInt(
			Age,
			fs,
		)
	}
	return len(fs)
}
