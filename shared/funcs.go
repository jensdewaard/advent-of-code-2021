package shared

func Filter[T any](predicate func(T) bool, is []T) []T {
	out := make([]T, 0)
	for _, i := range is {
		if predicate(i) {
			out = append(out, i)
		}
	}
	return out
}

func CountIf[T any](predicate func(T) bool, is []T) int {
	return len(Filter(predicate, is))
}

func Fold[S, T any](acc T, f func(a T, b S) T, is []S) T {
	for i := range is {
		acc = f(acc, is[i])
	}
	return acc
}

func Window[T any](k int, is []T) [][]T {
	out := make([][]T, len(is)-k+1)
	i := 0
	for i+k <= len(is) {
		out[i] = make([]T, k)
		copy(out[i], is[i:i+(k)])
		i++
	}
	return out
}

func Zip[S, T any](f func(S, S) T, as, bs []S) []T {
	is := make([]T, len(as))
	for i, a := range as {
		is[i] = f(a, bs[i])
	}
	return is
}

func Map[S, T any](f func(S) T, as []S) []T {
	bs := make([]T, len(as))
	for i, a := range as {
		bs[i] = f(a)
	}
	return bs
}

func FlatMap[S, T any](f func(S) []T, as []S) []T {
	bs := make([]T, 0)
	for _, a := range as {
		bs = append(bs, f(a)...)
	}
	return bs
}
