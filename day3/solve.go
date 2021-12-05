package day3

import (
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	lines := files.ReadLines("day3/input")
	counts := make([]BitCount, len(lines[0]))

	finalCounts := FoldToBitCount(
		counts,
		func(bc []BitCount, s string) []BitCount {
			return ZipWith(Count, bc, s)
		},
		lines,
	)
	gammaRate := FoldToString("", strings.Append, finalCounts)
	epsilonRate := Invert(gammaRate)
	return int(AsDecimal(gammaRate) * AsDecimal(epsilonRate))
}

func SolveB() int {
	return 0
}
