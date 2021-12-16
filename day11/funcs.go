package day11

func Neighbours(g IntGrid, p Position) []Position {
	out := make([]Position, 0)
	spaceAbove := p.Y-1 >= 0
	spaceLeft := p.X-1 >= 0
	spaceBelow := p.Y+1 < len(g)
	spaceRight := p.X+1 < len(g[p.Y])
	if spaceAbove {
		if spaceLeft {
			out = append(out, Position{p.X - 1, p.Y - 1})
		}
		out = append(out, Position{p.X, p.Y - 1})
		if spaceRight {
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
		if spaceLeft {
			out = append(out, Position{p.X - 1, p.Y + 1})
		}
		out = append(out, Position{p.X, p.Y + 1})
		if spaceRight {
			out = append(out, Position{p.X + 1, p.Y + 1})
		}
	}
	return out
}

func MapWithPositions(f func(int) int, g IntGrid, ps []Position) IntGrid {
	for _, p := range ps {
		g[p.Y][p.X] = f(g[p.Y][p.X])
	}
	return g
}

func MapIntGrid(f func(int) int, g IntGrid) IntGrid {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g[y][x] = f(g[y][x])
		}
	}
	return g
}

func FilterBoolGrid(pred func(bool) bool, g BoolGrid) []Position {
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

func FilterIntGrid(pred func(int) bool, g IntGrid) []Position {
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

func Get(g IntGrid, p Position) int {
	return g[p.Y][p.X]
}

func Set(g IntGrid, p Position, i int) {
	g[p.Y][p.X] = i
}

func AboveNine(i int) bool {
	return i > 9
}

func canStillFlash(is IntGrid, bs BoolGrid, p Position) bool {
	return is[p.Y][p.X] >= 9 && !bs[p.Y][p.X]
}

func IntGridEquals(a, b IntGrid) bool {
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
