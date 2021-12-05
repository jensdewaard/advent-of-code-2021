package day5

type Chart = [][]int

func NewChart(size int) Chart {
	chart := make([][]int, size)
	for i := range chart {
		chart[i] = make([]int, size)
	}
	return chart
}

type Position struct {
	X, Y int
}

type Line struct {
	Start, End Position
}
