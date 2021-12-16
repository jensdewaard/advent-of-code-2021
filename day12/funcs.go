package day12

import (
	"strings"

	"github.com/jensdewaard/advent-of-code-2021/shared/graphs"
)

func IsLargeCave(n graphs.Node) bool {
	return strings.ToUpper(n.Label) == n.Label
}

func MayVisitOnce(visited graphs.Path, n graphs.Node) bool {
	return !graphs.PathContains(visited, n)
}

func MayVisitTwice(visited graphs.Path, n graphs.Node) bool {
	if n.Label == "start" {
		return false
	}
	return !graphs.PathContains(visited, n) ||
		graphs.PathUnique(
			PathFilter(
				func(n graphs.Node) bool { return strings.ToLower(n.Label) == n.Label },
				visited,
			),
		)
}

type VisitFunc func(graphs.Path, graphs.Node) bool

func EnumeratePaths(
	g graphs.Graph,
	s, e graphs.Node,
	travelled, avoid graphs.Path,
	f VisitFunc,
) []graphs.Path {
	if s == e {
		return []graphs.Path{{e}}
	}
	ns := graphs.Neighbours(g, s)
	out := make([]graphs.Path, 0)
	for _, n := range ns {
		if IsLargeCave(n) || f(avoid, n) {
			// we consider a next node if it is either a large cave
			// or a small cave we haven't travelled through yet
			ps := EnumeratePaths(g, n, e, graphs.Path{n}, append(avoid, n), f)
			for _, p := range ps {
				if len(p) != 0 {
					out = append(out, append(travelled, p...))
				}
			}
		}
	}
	return out
}

func PathFilter(pred func(graphs.Node) bool, p graphs.Path) graphs.Path {
	out := make(graphs.Path, 0)
	for _, n := range p {
		if pred(n) {
			out = append(out, n)
		}
	}
	return out
}
