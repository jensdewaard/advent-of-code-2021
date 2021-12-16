package ints

import "math"

func MapSlicesToInt(f func([]int) int, as [][]int) []int {
	bs := make([]int, len(as))
	for i, a := range as {
		bs[i] = f(a)
	}
	return bs
}

func MapToInt(f func(int) int, as []int) []int {
	bs := make([]int, len(as))
	for i, a := range as {
		bs[i] = f(a)
	}
	return bs
}

func FlatMapToInt(f func(int) []int, as []int) []int {
	bs := make([]int, 0)
	for _, a := range as {
		bs = append(bs, f(a)...)
	}
	return bs
}

func ZipToBool(f func(int, int) bool, as, bs []int) []bool {
	is := make([]bool, len(as))
	for i, a := range as {
		is[i] = f(a, bs[i])
	}
	return is
}

func LessThan(a, b int) bool {
	return a < b
}

func Window(k int, is []int) [][]int {
	out := make([][]int, len(is)-k+1)
	i := 0
	for i+k <= len(is) {
		out[i] = make([]int, k)
		copy(out[i], is[i:i+(k)])
		i++
	}
	return out
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Fold(acc int, f func(a, b int) int, is []int) int {
	for i := range is {
		acc = f(acc, is[i])
	}
	return acc
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

func Filter(predicate func(int) bool, is []int) []int {
	out := make([]int, 0)
	for _, i := range is {
		if predicate(i) {
			out = append(out, i)
		}
	}
	return out
}

func CountIf(predicate func(int) bool, is []int) int {
	return len(Filter(predicate, is))
}

func Add(a, b int) int {
	return a + b
}
