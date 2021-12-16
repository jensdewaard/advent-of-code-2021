package main

import (
	"os"

	"github.com/jensdewaard/advent-of-code-2021/day1"
	"github.com/jensdewaard/advent-of-code-2021/day10"
	"github.com/jensdewaard/advent-of-code-2021/day11"
	"github.com/jensdewaard/advent-of-code-2021/day2"
	"github.com/jensdewaard/advent-of-code-2021/day3"
	"github.com/jensdewaard/advent-of-code-2021/day4"
	"github.com/jensdewaard/advent-of-code-2021/day5"
	"github.com/jensdewaard/advent-of-code-2021/day6"
	"github.com/jensdewaard/advent-of-code-2021/day7"
	"github.com/jensdewaard/advent-of-code-2021/day8"
	"github.com/jensdewaard/advent-of-code-2021/day9"
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
	case "5":
		println(day5.SolveA())
		println(day5.SolveB())
	case "6":
		println(day6.SolveA())
		println(day6.SolveB())
	case "7":
		println(day7.SolveA())
		println(day7.SolveB())
	case "8":
		println(day8.SolveA())
		println(day8.SolveB())
	case "9":
		println(day9.SolveA())
		println(day9.SolveB())
	case "10":
		println(day10.SolveA())
		println(day10.SolveB())
	case "11":
		println(day11.SolveA())
		println(day11.SolveB())
	default:
		println(day11.SolveA())
		println(day11.SolveB())
	}
}
