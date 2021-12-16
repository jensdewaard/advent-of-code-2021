package graphs

import (
	"fmt"
	"strings"
)

type Graph struct {
	Edges []Edge
	Nodes []Node
}

type Node struct {
	Label string
}

type Edge struct {
	Start Node
	End   Node
}

func NewGraph() Graph {
	e := make([]Edge, 0)
	n := make([]Node, 0)
	return Graph{e, n}
}

func AddEdge(g Graph, s, e string) Graph {
	var start, end Node
	for _, n := range g.Nodes {
		if n.Label == s {
			start = n
		}
		if n.Label == e {
			end = n
		}
	}
	if start.Label == "" {
		start = Node{s}
		g.Nodes = append(g.Nodes, start)
	}
	if end.Label == "" {
		end = Node{e}
		g.Nodes = append(g.Nodes, end)
	}
	g.Edges = append(g.Edges, Edge{start, end})
	return g
}

type Path []Node

func PathContains(p Path, n Node) bool {
	for _, m := range p {
		if n.Label == m.Label {
			return true
		}
	}
	return false
}

func SprintPath(p Path) string {
	sb := strings.Builder{}
	for _, n := range p {
		sb.WriteString(fmt.Sprintf("%s,", n.Label))
	}
	return sb.String()
}

func Neighbours(g Graph, n Node) []Node {
	out := make([]Node, 0)
	for _, e := range g.Edges {
		if e.Start == n {
			out = append(out, e.End)
		}
		if e.End == n {
			out = append(out, e.Start)
		}
	}
	return out
}

func GetNode(g Graph, l string) Node {
	for _, n := range g.Nodes {
		if n.Label == l {
			return n
		}
	}
	return Node{""}
}

func PathUnique(p Path) bool {
	set := make(map[string]struct{})
	for _, n := range p {
		set[n.Label] = struct{}{}
	}
	return len(set) == len(p)
}
