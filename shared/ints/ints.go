package ints

import (
	"math"

	. "github.com/jensdewaard/advent-of-code-2021/shared"
)

func LessThan(a, b int) bool {
	return a < b
}

func Sum(is []int) int {
	return Fold(0, func(a, b int) int { return a + b }, is)
}

func Mult(x, y int) int {
	return x * y
}

func SumIf(is []int, bs []bool) int {
	total := 0
	for i, n := range is {
		if bs[i] {
			total += n
		}
	}
	return total
}

func Median(is []int) int {
	l := len(is)
	if l%2 == 2 {
		return (is[l/2-1] + is[l/2]) / 2
	} else {
		return is[l/2]
	}
}

func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func Diff(a, b int) int {
	return Abs(a - b)
}

func Max(is []int) (int, int) {
	index := 0
	l := math.MinInt
	for i, v := range is {
		if v > l {
			l = v
			index = i
		}
	}
	return index, l
}

func Min(is []int) (int, int) {
	index := 0
	l := math.MaxInt
	for i, v := range is {
		if v < l {
			l = v
			index = i
		}
	}
	return index, l
}

func Average(is []int) int {
	return Sum(is) / len(is)
}

func Add(a, b int) int {
	return a + b
}
