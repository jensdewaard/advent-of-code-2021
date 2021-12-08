package day8

import "strings"

type LitSegments = string

func Difference(l, m LitSegments) LitSegments {
	out := ""
	for _, c := range l {
		if !strings.ContainsRune(m, c) {
			out += string(c)
		}
	}
	for _, c := range m {
		if !strings.ContainsRune(l, c) &&
			!strings.ContainsRune(out, c) {
			out += string(c)
		}
	}
	return out
}

func Intersect(l, m LitSegments) LitSegments {
	out := ""
	for _, c := range l {
		if strings.ContainsRune(m, c) {
			out += string(c)
		}
	}
	return out
}

func SegmentContains(l, m LitSegments) bool {
	for _, c := range m {
		if !strings.ContainsRune(l, c) {
			return false
		}
	}
	return true
}

type RecognizedNumber struct {
	Segments LitSegments
	Number   int
}

type Mapping = []byte

func NewMapping() Mapping {
	return []byte{'?', '?', '?', '?', '?', '?', '?'}
}

func GetZero(m Mapping) LitSegments {
	return string([]byte{m[0], m[1], m[2], m[4], m[5], m[6]})
}

func GetTwo(m Mapping) LitSegments {
	return string([]byte{m[0], m[2], m[3], m[4], m[6]})
}

func GetThree(m Mapping) LitSegments {
	return string([]byte{m[0], m[2], m[3], m[5], m[6]})
}

func GetFive(m Mapping) LitSegments {
	return string([]byte{m[0], m[1], m[3], m[5], m[6]})
}

func GetSix(m Mapping) LitSegments {
	return string([]byte{m[0], m[1], m[3], m[4], m[5], m[6]})
}

func GetNine(m Mapping) LitSegments {
	return string([]byte{m[0], m[1], m[2], m[3], m[5], m[6]})
}
