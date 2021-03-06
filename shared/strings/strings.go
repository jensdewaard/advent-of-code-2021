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
	return out
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

func Join(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

func Filter(predicate func(string) bool, ss []string) []string {
	out := make([]string, 0)
	for _, s := range ss {
		if predicate(s) {
			out = append(out, s)
		}
	}
	return out
}

func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func FoldToString(acc string, f func(string, string) string, ss []string) string {
	for _, s := range ss {
		acc = f(acc, s)
	}
	return acc
}