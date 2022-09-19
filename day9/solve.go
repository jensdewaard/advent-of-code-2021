package day9

import (
	"sort"

	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	lines := files.ReadLines("day9/input")
	g := NewGrid(len(lines[0]), len(lines))
	for y := 0; y < len(lines); y++ {
		g[y] = shared.Map(strings.ParseInt, strings.Split(lines[y], ""))
	}
	ls := Filter(IsLowPoint, g)
	calcRisk := func(p Position) int { return Risk(g, p) }
	rs := MapPosition(calcRisk, ls)
	return ints.Sum(rs)
}

func SolveB() int {
	lines := files.ReadLines("day9/input")
	g := NewGrid(len(lines[0]), len(lines))
	for y := 0; y < len(lines); y++ {
		g[y] = shared.Map(strings.ParseInt, strings.Split(lines[y], ""))
	}
	ls := Filter(IsLowPoint, g)
	bs := make([][]Position, len(ls))
	for i, l := range ls {
		bs[i] = Basin(g, []Position{l})
	}
	ss := make([]int, len(bs))
	for i, b := range bs {
		ss[i] = len(b)
	}
	sort.Ints(ss)
	return shared.Fold(1, ints.Mult, ss[len(ss)-3:])
}
