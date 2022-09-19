package strings

import (
	"log"
	"strconv"
	"strings"
)

func Append(a, b string) string {
	return a + b
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
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
