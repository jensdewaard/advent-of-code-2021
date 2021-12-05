package day5

import "github.com/jensdewaard/advent-of-code-2021/shared/files"

func SolveA() int {
	lines := MapToLines(ParseLine, files.ReadLines("day5/input"))
	c := NewChart(1000)
	c = FoldChart(c, DrawLine, lines)
	return CountIf(func(i int) bool {
		return i >= 2
	}, c)
}

func SolveB() int {
	lines := MapToLines(ParseLine, files.ReadLines("day5/input"))
	c := NewChart(1000)
	c = FoldChart(c, DrawLineDiagonal, lines)
	return CountIf(func(i int) bool {
		return i >= 2
	}, c)
}
