package grids

import (
	"fmt"
	str "strings"

	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

type Grid [][]int

func NewGrid(x, y int) Grid {
	m := make([][]int, y)
	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
	}
	return m
}

func ParseGrid(ss []string) Grid {
	g := NewGrid(len(ss[0]), len(ss))
	for y := 0; y < len(ss); y++ {
		g[y] = shared.Map(
			strings.ParseInt,
			strings.Split(ss[y], ""),
		)
	}
	return g
}

func SprintGrid(m Grid) string {
	sb := str.Builder{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			sb.WriteString(fmt.Sprint(m[j][i]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

type Position struct {
	X, Y int
}

func Neighbours(g Grid, p Position, diagonal bool) []Position {
	out := make([]Position, 0)
	spaceAbove := p.Y-1 >= 0
	spaceLeft := p.X-1 >= 0
	spaceBelow := p.Y+1 < len(g)
	spaceRight := p.X+1 < len(g[p.Y])
	if spaceAbove {
		if spaceLeft && diagonal {
			out = append(out, Position{p.X - 1, p.Y - 1})
		}
		out = append(out, Position{p.X, p.Y - 1})
		if spaceRight && diagonal {
			out = append(out, Position{p.X + 1, p.Y - 1})
		}
	}
	if spaceLeft {
		out = append(out, Position{p.X - 1, p.Y})
	}
	if spaceRight {
		out = append(out, Position{p.X + 1, p.Y})
	}
	if spaceBelow {
		if spaceLeft && diagonal {
			out = append(out, Position{p.X - 1, p.Y + 1})
		}
		out = append(out, Position{p.X, p.Y + 1})
		if spaceRight && diagonal {
			out = append(out, Position{p.X + 1, p.Y + 1})
		}
	}
	return out
}

func MapWithPositions(f func(int) int, g Grid, ps []Position) Grid {
	for _, p := range ps {
		g[p.Y][p.X] = f(g[p.Y][p.X])
	}
	return g
}

func MapGrid(f func(int) int, g Grid) Grid {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g[y][x] = f(g[y][x])
		}
	}
	return g
}

func FilterGrid(pred func(int) bool, g Grid) []Position {
	out := make([]Position, 0)
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if pred(g[y][x]) {
				out = append(out, Position{x, y})
			}
		}
	}
	return out
}

func FilterPositions(pred func(Position) bool, ps []Position) []Position {
	out := make([]Position, 0)
	for _, p := range ps {
		if pred(p) {
			out = append(out, p)
		}
	}
	return out
}

func GridEquals(a, b Grid) bool {
	if len(a) != len(b) {
		return false
	}
	for y := 0; y < len(a); y++ {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := 0; x < len(a[y]); x++ {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}

func MapPositions(f func(p Position) Position, ps []Position) []Position {
	out := make([]Position, len(ps))
	for i, p := range ps {
		out[i] = f(p)
	}
	return out
}

func FoldGridLines(acc int, f func(int, int) int, g Grid) []int {
	out := make([]int, len(g))
	for i, l := range g {
		lAcc := acc
		for _, v := range l {
			lAcc = f(lAcc, v)
		}
		out[i] = lAcc
	}
	return out
}
