package day3

import (
	"log"

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
	lines := files.ReadLines("day3/input")
	oxygen := AsDecimal(RepeatFilter(lines, MostCommon(), "oxygen"))
	co2 := AsDecimal(RepeatFilter(lines, LeastCommon(), "co2"))
	return int(oxygen * co2)
}

func MostCommon() func([]BitCount, int) BitStringFilter {
	return func(bs []BitCount, i int) BitStringFilter {
		bit := bs[i].Bit
		if bs[i].Count == 0 {
			bit = "1"
		}
		return func(b BitString) bool {
			return string(b[i]) == bit
		}
	}
}

func LeastCommon() func([]BitCount, int) BitStringFilter {
	return func(bs []BitCount, i int) BitStringFilter {
		bit := bs[i].Bit
		if bs[i].Count == 0 {
			bit = "1"
		}
		return func(b BitString) bool {
			return string(b[i]) != bit
		}
	}
}

func RepeatFilter(lines []BitString, filter func([]BitCount, int) BitStringFilter, lookingFor string) BitString {
	index := 0
	candidates := lines
	for len(candidates) > 1 {
		counts := make([]BitCount, len(candidates[0]))
		occurences := FoldToBitCount(
			counts,
			func(bc []BitCount, s string) []BitCount {
				return ZipWith(Count, bc, s)
			},
			candidates,
		)
		candidates = Filter(filter(occurences, index), candidates)
		index++
	}
	if len(candidates) != 1 {
		log.Fatalf("no more candidates for %s remaining", lookingFor)
	}
	return candidates[0]
}
