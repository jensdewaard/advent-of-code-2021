package day11

import "github.com/jensdewaard/advent-of-code-2021/shared/files"

func SolveA() int {
	octopi := ParseGrid(files.ReadLines("day11/input"))
	flashes := 0
	f := 0
	for i := 0; i < 100; i++ {
		octopi, f = DoStep(octopi)
		flashes = flashes + f
	}
	return flashes
}

func SolveB() int {
	octopi := ParseGrid(files.ReadLines("day11/input"))
	zeroGrid := NewIntGrid(len(octopi), len(octopi[1]))
	i := 0
	for !IntGridEquals(octopi, zeroGrid) {
		octopi, _ = DoStep(octopi)
		i++
	}
	return i
}

func DoStep(g IntGrid) (IntGrid, int) {
	flashed := NewBoolGrid(len(g), len(g[1]))
	// 1. increase all by 1
	g = MapIntGrid(func(i int) int { return i + 1 }, g)
	// 2. start flashing
	flashables := FilterPositions(
		func(p Position) bool { return canStillFlash(g, flashed, p) },
		FilterIntGrid(AboveNine, g),
	)
	for len(flashables) > 0 {
		for _, p := range flashables {
			g = Flash(g, p)
			flashed[p.Y][p.X] = true
		}
		flashables = FilterPositions(
			func(p Position) bool { return canStillFlash(g, flashed, p) },
			FilterIntGrid(AboveNine, g),
		)
	}
	// 3. reset to zero
	flashedPos := FilterBoolGrid(func(b bool) bool { return b }, flashed)
	for _, p := range flashedPos {
		g[p.Y][p.X] = 0
	}
	return g, len(flashedPos)
}

func Flash(g IntGrid, p Position) IntGrid {
	ns := Neighbours(g, p)
	return MapWithPositions(func(i int) int { return i + 1 }, g, ns)
}
