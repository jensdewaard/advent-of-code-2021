package day8

import (
	"fmt"
	"log"

	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

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

func PrintMapping(m Mapping) {
	fmt.Printf(" %s%s%s%s \n", string(m[0]), string(m[0]), string(m[0]), string(m[0]))
	fmt.Printf("%s    %s\n", string(m[1]), string(m[2]))
	fmt.Printf("%s    %s\n", string(m[1]), string(m[2]))
	fmt.Printf(" %s%s%s%s \n", string(m[3]), string(m[3]), string(m[3]), string(m[3]))
	fmt.Printf("%s    %s\n", string(m[4]), string(m[5]))
	fmt.Printf("%s    %s\n", string(m[4]), string(m[5]))
	fmt.Printf(" %s%s%s%s \n", string(m[6]), string(m[6]), string(m[6]), string(m[6]))
}

func Decode(l LitSegments, rs []RecognizedNumber) int {
	for _, r := range rs {
		if EquallyLit(l, r.Segments) {
			return r.Number
		}
	}
	log.Fatalf("unable to decode %s", l)
	return -1
}

func EquallyLit(l, s LitSegments) bool {
	if len(l) != len(s) {
		return false
	}
	diff := Difference(l, s)
	return diff == ""
}

func SolveForLine(l string) int {
	splitted := strings.Split(l, "|")
	inputs := MapToRecognizedNumber(
		Recognize,
		strings.Fields(splitted[0]),
	)
	isRecognized := func(r RecognizedNumber) bool {
		return r.Number != -1
	}
	rs := Solve(
		Filter(isRecognized, inputs),
		Filter(
			func(r RecognizedNumber) bool { return !isRecognized(r) },
			inputs,
		),
	)
	outputs := strings.Fields(splitted[1])
	os := strings.MapToString(
		func(ls LitSegments) string {
			return fmt.Sprint(Decode(ls, rs))
		},
		outputs,
	)
	return strings.ParseInt(
		strings.FoldToString("", strings.Append, os),
	)
}

func Solve(rs []RecognizedNumber, us []RecognizedNumber) []RecognizedNumber {
	if len(rs) == 10 {
		// all numbers mapped
		return rs
	}
	if len(us) == 0 {
		log.Fatalf("unable to create consistent mapping: %v", rs)
	}
	u := us[0]
	sevenSegments := GetNumber(rs, 7).Segments
	fourSegments := GetNumber(rs, 4).Segments
	if len(u.Segments) == 5 {
		// 2, 3, 5
		if SegmentContains(u.Segments, sevenSegments) {
			u.Number = 3
		} else if len(Intersect(u.Segments, fourSegments)) == 3 {
			u.Number = 5
		} else {
			u.Number = 2
		}
	} else if len(u.Segments) == 6 {
		// 0, 6, 9
		if SegmentContains(u.Segments, fourSegments) {
			u.Number = 9
		} else if SegmentContains(u.Segments, sevenSegments) {
			u.Number = 0
		} else {
			u.Number = 6
		}
	}
	if u.Number == -1 {
		log.Fatalf("could not recognize %s", u.Segments)
	}
	if !ContainsNumber(rs, u.Number) {
		rs = append(rs, u)
	}
	return Solve(rs, us[1:])
}

func GetNumber(rs []RecognizedNumber, n int) RecognizedNumber {
	for _, r := range rs {
		if r.Number == n {
			return r
		}
	}
	return RecognizedNumber{"", -1}
}

func ContainsNumber(rs []RecognizedNumber, n int) bool {
	if len(rs) == 0 {
		return false
	}
	if rs[0].Number == n {
		return true
	}
	return ContainsNumber(rs[1:], n)
}

func MappingContains(m Mapping, c byte) bool {
	for _, d := range m {
		if d == c {
			return true
		}
	}
	return false
}

func RecognizedEqual(rs, ls []RecognizedNumber) bool {
	if len(ls) != len(rs) {
		return false
	}
	for _, l := range ls {
		if !ContainsNumber(rs, l.Number) {
			return false
		}
	}
	return true
}
