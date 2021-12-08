package day8

func Recognize(l LitSegments) RecognizedNumber {
	if len(l) == 2 {
		return RecognizedNumber{l, 1}
	} else if len(l) == 3 {
		return RecognizedNumber{l, 7}
	} else if len(l) == 4 {
		return RecognizedNumber{l, 4}
	} else if len(l) == 7 {
		return RecognizedNumber{l, 8}
	}
	return RecognizedNumber{l, -1}
}

func MapToRecognizedNumber(
	f func(LitSegments) RecognizedNumber,
	ls []LitSegments,
) []RecognizedNumber {
	out := make([]RecognizedNumber, len(ls))
	for i, l := range ls {
		out[i] = f(l)
	}
	return out
}

func Filter(predicate func(RecognizedNumber) bool, rs []RecognizedNumber) []RecognizedNumber {
	out := make([]RecognizedNumber, 0)
	for _, r := range rs {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

func UniqueSegments(rs []RecognizedNumber) []RecognizedNumber {
	out := make([]RecognizedNumber, 0)
	have := make([]LitSegments, 0)
	for _, r := range rs {
		if !Contains(have, r.Segments) {
			out = append(out, r)
		}
	}
	return out
}
func Contains(ls []LitSegments, l LitSegments) bool {
	for _, m := range ls {
		if l == m {
			return true
		}
	}
	return false
}
