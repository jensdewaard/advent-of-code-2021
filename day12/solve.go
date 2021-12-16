package day12

import (
	"strings"

	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/graphs"
)

func Solve(f VisitFunc) int {
	lines := files.ReadLines("day12/input")
	graph := graphs.NewGraph()
	for _, l := range lines {
		ss := strings.Split(l, "-")
		graph = graphs.AddEdge(graph, ss[0], ss[1])
	}
	start := graphs.GetNode(graph, "start")
	end := graphs.GetNode(graph, "end")
	paths := EnumeratePaths(graph, start, end, graphs.Path{start}, graphs.Path{start}, f)
	return len(paths)
}

func SolveA() int {
	return Solve(MayVisitOnce)
}

func SolveB() int {
	return Solve(MayVisitTwice)
}
