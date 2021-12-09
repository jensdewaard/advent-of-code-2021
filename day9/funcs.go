package day9

import "github.com/jensdewaard/advent-of-code-2021/shared/ints"

func Risk(g Grid, p Position) int {
	return Get(g, p) + 1
}

func Neighbours(g Grid, p Position) []Position {
	out := make([]Position, 0)
	if p.Y-1 >= 0 {
		out = append(out, Position{p.X, p.Y - 1})
	}
	if p.X+1 < len(g[p.Y]) {
		out = append(out, Position{p.X + 1, p.Y})
	}
	if p.Y+1 < len(g) {
		out = append(out, Position{p.X, p.Y + 1})
	}
	if p.X-1 >= 0 {
		out = append(out, Position{p.X - 1, p.Y})
	}
	return out
}

func IsLowPoint(g Grid, p Position) bool {
	ns := Neighbours(g, p)
	lowest := true
	for _, n := range ns {
		lowest = lowest && ints.LessThan(Get(g, p), Get(g, n))
	}
	return lowest
}

func Filter(predicate func(Grid, Position) bool, g Grid) []Position {
	out := make([]Position, 0)
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if predicate(g, Position{x, y}) {
				out = append(out, Position{x, y})
			}
		}
	}
	return out
}

func MapPosition(f func(Position) int, ps []Position) []int {
	is := make([]int, len(ps))
	for i, p := range ps {
		is[i] = f(p)
	}
	return is
}

func FlatMapPosition(f func(Position) []Position, ps []Position) []Position {
	out := make([]Position, 0)
	for _, p := range ps {
		out = append(out, f(p)...)
	}
	return out
}

func Basin(g Grid, ps []Position) []Position {
	getNeighbours := func(p Position) []Position { return Neighbours(g, p) }
	ns := FlatMapPosition(getNeighbours, ps)
	unchanged := true
	for _, n := range ns {
		if !Contains(n, ps) && Get(g, n) != 9 {
			ps = append(ps, n)
			unchanged = false
		}
	}
	if unchanged {
		return ps
	} else {
		return Basin(g, ps)
	}
}

func SliceEqual(ps, qs []Position) bool {
	if len(ps) != len(qs) {
		return false
	}
	for _, p := range ps {
		if !Contains(p, qs) {
			return false
		}
	}
	return true
}

func Contains(q Position, ps []Position) bool {
	for _, p := range ps {
		if p == q {
			return true
		}
	}
	return false
}
