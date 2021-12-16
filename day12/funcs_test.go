package day12

import (
	"testing"

	"github.com/jensdewaard/advent-of-code-2021/shared/graphs"
)

func Test_IsLargeCave(t *testing.T) {
	tests := []struct {
		n        graphs.Node
		expected bool
	}{
		{graphs.Node{Label: "AA"}, true},
		{graphs.Node{Label: "start"}, false},
		{graphs.Node{Label: "b"}, false},
		{graphs.Node{Label: "HN"}, true},
	}
	for _, tt := range tests {
		actual := IsLargeCave(tt.n)
		if actual != tt.expected {
			t.Errorf("for %s expected %v, got %v", tt.n.Label, tt.expected, actual)
		}
	}
}

func Test_EnumeratePathsOnePath(t *testing.T) {
	start := graphs.Node{Label: "start"}
	end := graphs.Node{Label: "end"}
	g := graphs.Graph{
		Nodes: []graphs.Node{start, end},
		Edges: []graphs.Edge{{Start: start, End: end}},
	}
	ps := EnumeratePaths(g, start, end, graphs.Path{start}, graphs.Path{start}, MayVisitOnce)
	if len(ps) != 1 {
		t.Errorf("expected 1 path, got %v", len(ps))
	}
	if ps[0][0] != start || ps[0][1] != end {
		t.Errorf("expected path start->end, got %s", graphs.SprintPath(ps[0]))
	}
}

func Test_EnumeratePathsTwoPaths(t *testing.T) {
	start := graphs.Node{Label: "start"}
	end := graphs.Node{Label: "end"}
	a := graphs.Node{Label: "a"}
	b := graphs.Node{Label: "b"}
	g := graphs.Graph{
		Nodes: []graphs.Node{start, end, a, b},
		Edges: []graphs.Edge{
			{Start: start, End: a},
			{Start: start, End: b},
			{Start: a, End: end},
			{Start: b, End: end},
		},
	}
	ps := EnumeratePaths(g, start, end, graphs.Path{start}, graphs.Path{start}, MayVisitOnce)
	if len(ps) != 2 {
		t.Errorf("expected 2 paths, got %v", len(ps))
	}
	if ps[0][0] != start || ps[0][1] != a || ps[0][2] != end {
		t.Errorf("expected path start->a->end, got %s", graphs.SprintPath(ps[0]))
	}
	if ps[1][0] != start || ps[1][1] != b || ps[1][2] != end {
		t.Errorf("expected path start->b->end, got %s", graphs.SprintPath(ps[1]))
	}
}

func Test_VisitTwice(t *testing.T) {
	start := graphs.Node{Label: "start"}
	end := graphs.Node{Label: "end"}
	a := graphs.Node{Label: "a"}
	b := graphs.Node{Label: "b"}
	c := graphs.Node{Label: "c"}
	g := graphs.Graph{
		Nodes: []graphs.Node{start, end, a, b, c},
		Edges: []graphs.Edge{
			{Start: start, End: a},
			{Start: start, End: b},
			{Start: start, End: c},
			{Start: a, End: b},
			{Start: b, End: c},
			{Start: a, End: end},
			{Start: b, End: end},
			{Start: c, End: end},
		},
	}
	ps := EnumeratePaths(g, start, end, graphs.Path{start}, graphs.Path{start}, MayVisitTwice)
	if len(ps) != 17 {
		t.Errorf("expected 17 paths, got %v", len(ps))
	}
}
