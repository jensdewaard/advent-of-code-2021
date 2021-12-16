package day11

import (
	"fmt"

	str "strings"

	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

type IntGrid [][]int
type BoolGrid [][]bool

func NewIntGrid(x, y int) IntGrid {
	m := make([][]int, y)
	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
	}
	return m
}

func NewBoolGrid(x, y int) BoolGrid {
	g := make([][]bool, y)
	for i := 0; i < y; i++ {
		g[i] = make([]bool, x)
	}
	return g
}

func GetInt(m IntGrid, p Position) int {
	return m[p.Y][p.X]
}

func ParseGrid(ss []string) IntGrid {
	g := NewIntGrid(len(ss[0]), len(ss))
	for y := 0; y < len(ss); y++ {
		g[y] = strings.MapToInt(
			strings.ParseInt,
			strings.Split(ss[y], ""),
		)
	}
	return g
}

func SprintGrid(m IntGrid) string {
	sb := str.Builder{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			sb.WriteString(fmt.Sprint(GetInt(m, Position{j, i})))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

type Position struct {
	X, Y int
}
