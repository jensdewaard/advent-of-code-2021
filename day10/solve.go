package day10

import (
	"sort"

	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
)

func SolveA() int {
	lines := files.ReadLines("day10/input")
	parse := func(s string) LineStatus {
		return FindClosingDelimiter(s, []byte{})
	}
	filter := func(l LineStatus) bool {
		return l.err == "corrupt"
	}
	ds := Filter(filter, MapToStatus(parse, lines))
	scores := MapToInt(Score, ds)
	return ints.Sum(scores)
}

func SolveB() int {
	lines := files.ReadLines("day10/input")
	parse := func(s string) LineStatus {
		return FindClosingDelimiter(s, []byte{})
	}
	filter := func(l LineStatus) bool {
		return l.err == "incomplete"
	}
	ds := Filter(filter, MapToStatus(parse, lines))
	scores := shared.Map(
		ScoreMissing,
		MapToString(func(ls LineStatus) string { return ls.missing }, ds),
	)
	sort.Ints(scores)
	return scores[len(scores)/2]
}
