package day13

import (
	"strings"

	"github.com/jensdewaard/advent-of-code-2021/shared/grids"
)

type Paper = grids.Grid
type Fold struct {
	Line int
	Type string
}

func FoldPaperHorizontally(p Paper, line int) Paper {
	dots := grids.FilterGrid(
		func(i int) bool { return i > 0 },
		p,
	)
	flippedDots := grids.MapPositions(
		func(p grids.Position) grids.Position { return MirrorPositionHorizontally(p, line) },
		dots,
	)
	for _, d := range flippedDots {
		p = AddDot(p, d)
	}
	return CutPaperHorizontal(p, line)
}

func CutPaperHorizontal(p Paper, line int) Paper {
	t := make(Paper, len(p))
	for i := range p {
		t[i] = p[i][:line]
	}
	return t
}

func FoldPaperVertically(p Paper, line int) Paper {
	dots := grids.FilterGrid(
		func(i int) bool { return i > 0 },
		p,
	)
	flippedDots := grids.MapPositions(
		func(p grids.Position) grids.Position { return MirrorPositionVertically(p, line) },
		dots,
	)
	for _, d := range flippedDots {
		p = AddDot(p, d)
	}
	return CutPaperVertical(p, line)
}

func CutPaperVertical(p Paper, line int) Paper {
	t := make(Paper, line)
	for i := 0; i < line; i++ {
		t[i] = p[i]
	}
	return t
}

func MirrorPositionHorizontally(p grids.Position, line int) grids.Position {
	return grids.Position{
		X: p.X - 2*(p.X-line),
		Y: p.Y,
	}
}

func MirrorPositionVertically(p grids.Position, line int) grids.Position {
	return grids.Position{
		X: p.X,
		Y: p.Y - 2*(p.Y-line),
	}
}

func AddDot(p Paper, c grids.Position) Paper {
	p = GrowPaper(p, c.Y+1, c.X+1)
	p[c.Y][c.X] = 1
	return p
}

func GrowPaper(p Paper, h, w int) Paper {
	if len(p) < h {
		t := make(Paper, h)
		copy(t, p)
		for i := len(p); i < len(t); i++ {
			t[i] = make([]int, w)
		}
		p = t
	}
	for y := 0; y < len(p); y++ {
		if len(p[y]) < w {
			t := make([]int, w)
			copy(t, p[y])
			p[y] = t
		}
	}
	return p
}

func SprintPaper(p Paper) string {
	sb := strings.Builder{}
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			if p[y][x] > 0 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
