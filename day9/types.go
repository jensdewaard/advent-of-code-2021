package day9

import "fmt"

type Grid [][]int

func NewGrid(x, y int) Grid {
	m := make([][]int, y)
	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
	}
	return m
}

func Get(m Grid, p Position) int {
	return m[p.Y][p.X]
}

func PrintGrid(m Grid) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(Get(m, Position{j, i}))
		}
		fmt.Println()
	}
}

type Position struct {
	X, Y int
}