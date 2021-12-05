package strings

import (
	"log"
	"strconv"
	"strings"
)

func Append(a, b string) string {
	return a + b
}

func MapToString(f func(string) string, ss []string) []string {
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = f(s)
	}
	return ss
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func MapToInt(f func(string) int, ss []string) []int {
	is := make([]int, len(ss))
	for i, s := range ss {
		is[i] = f(s)
	}
	return is
}

func SliceEqual(a, b []string) bool {
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

func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func Fields(s string) []string {
    return strings.Fields(s)
}
