package day4

import (
	"log"
	"strings"

	mystrings "github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

type Board struct {
	Numbers []int
	Marked  []bool
}

func NewBoard(lines []string) Board {
	b := Board{}
	b.Numbers = make([]int, 25)
	b.Marked = make([]bool, 25)

	if len(lines) != 5 {
		log.Fatalf("incorrect number of lines given, got %v, need 5", len(lines))
	}
	for i := 0; i < 5; i++ {
		numbers := strings.Fields(lines[i])
		for j := 0; j < 5; j++ {
			b.Numbers[i*5+j] = mystrings.ParseInt(numbers[j])
		}
	}

	return b
}
