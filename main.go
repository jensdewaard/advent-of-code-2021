package main

import (
	"os"

	"github.com/jensdewaard/advent-of-code-2021/day1"
)

func main() {
	day := "default"
	if len(os.Args) >= 2 {
		day = os.Args[2]
	}

	switch day {
	case "1":
		println(day1.SolveA())
		println(day1.SolveB())
	case "default":
		fallthrough
	default:
		println(day1.SolveB())
	}
}
