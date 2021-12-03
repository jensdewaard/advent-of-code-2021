package strings

import (
	"log"
	"strconv"
)

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
