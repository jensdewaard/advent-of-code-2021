package ints

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
