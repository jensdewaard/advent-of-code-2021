package day5

import "github.com/jensdewaard/advent-of-code-2021/shared/strings"

func DrawLine(c Chart, l Line) Chart {
	for _, p := range Expand(l) {
		c = DrawPoint(c, p)
	}
	return c
}

func DrawLineDiagonal(c Chart, l Line) Chart {
	for _, p := range ExpandDiagonal(l) {
		c = DrawPoint(c, p)
	}
	return c
}

func DrawPoint(c Chart, p Position) Chart {
	c[p.Y][p.X]++
	return c
}

func FoldChart(acc Chart, f func(Chart, Line) Chart, ls []Line) Chart {
	for _, l := range ls {
		acc = f(acc, l)
	}
	return acc
}

func Expand(l Line) []Position {
	out := make([]Position, 0)
	if l.Start.X == l.End.X {
		// horizontal line
		x := l.Start.X
		if l.Start.Y < l.End.Y {
			// going right
			for y := l.Start.Y; y <= l.End.Y; y++ {
				out = append(out, Position{x, y})
			}
		} else { // going left
			for y := l.Start.Y; y >= l.End.Y; y-- {
				out = append(out, Position{x, y})
			}
		}
	} else if l.Start.Y == l.End.Y {
		// vertical line
		y := l.Start.Y
		if l.Start.X < l.End.X {
			// going up
			for x := l.Start.X; x <= l.End.X; x++ {
				out = append(out, Position{x, y})
			}
		} else { // going down
			for x := l.Start.X; x >= l.End.X; x-- {
				out = append(out, Position{x, y})
			}
		}
	}
	return out
}

func ExpandDiagonal(l Line) []Position {
	out := Expand(l)
	x := l.Start.X
	y := l.Start.Y
	var xOp func(int) int
	var yOp func(int) int
	if l.Start.X < l.End.X && l.Start.Y < l.End.Y { // up right
		xOp = func(x int) int { return x + 1 }
		yOp = func(y int) int { return y + 1 }
	} else if l.Start.X < l.End.X && l.Start.Y > l.End.Y { // down right
		xOp = func(x int) int { return x + 1 }
		yOp = func(y int) int { return y - 1 }
	} else if l.Start.X > l.End.X && l.Start.Y < l.End.Y { // up left
		xOp = func(x int) int { return x - 1 }
		yOp = func(y int) int { return y + 1 }
	} else if l.Start.X > l.End.X && l.Start.Y > l.End.Y { // down left
		xOp = func(x int) int { return x - 1 }
		yOp = func(y int) int { return y - 1 }
	}
	diagonal := false
	for x != l.End.X && y != l.End.Y {
		out = append(out, Position{x, y})
		x = xOp(x)
		y = yOp(y)
		diagonal = true
	}
	if diagonal {
		out = append(out, Position{x, y})
	}
	return out
}

func ParsePosition(s string) Position {
	ss := strings.Split(s, ",")
	return Position{strings.ParseInt(ss[0]), strings.ParseInt(ss[1])}
}

func ParseLine(s string) Line {
	ss := strings.Fields(s)
	return Line{
		ParsePosition(ss[0]),
		ParsePosition(ss[2]),
	}
}

func MapToLines(f func(string) Line, ss []string) []Line {
	ls := make([]Line, len(ss))
	for i, s := range ss {
		ls[i] = f(s)
	}
	return ls
}

func CountIf(predicate func(int) bool, c Chart) int {
	count := 0
	for _, column := range c {
		for _, cell := range column {
			if predicate(cell) {
				count++
			}
		}
	}
	return count
}

func PrintChart(c Chart) {
	for x := 0; x < len(c); x++ {
		for y := 0; y < len(c[x]); y++ {
			if c[x][y] == 0 {
				print(".")
			} else {
				print(c[x][y])
			}
		}
		println("")
	}
}
