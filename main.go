package main

import (
	"os"

	"github.com/jensdewaard/advent-of-code-2021/day1"
	"github.com/jensdewaard/advent-of-code-2021/day2"
	"github.com/jensdewaard/advent-of-code-2021/day3"
	"github.com/jensdewaard/advent-of-code-2021/day4"
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
	case "2":
		println(day2.SolveA())
		println(day2.SolveB())
	case "3":
		println(day3.SolveA())
		println(day3.SolveB())
	case "4":
		println(day4.SolveA())
		println(day4.SolveB())
	default:
		println(day4.SolveA())
		println(day4.SolveB())
	}
}
