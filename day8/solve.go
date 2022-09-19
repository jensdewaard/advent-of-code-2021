package day8

import (
	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	lines := files.ReadLines("day8/input")
	outputs := shared.Map(
		func(s string) string {
			return strings.Split(s, " | ")[1]
		},
		lines,
	)
	allOutputs := strings.Join(outputs, " ")
	recognizedNumbers := MapToRecognizedNumber(
		Recognize,
		strings.Fields(allOutputs),
	)
	equalsOneFourSevenOrEight := func(r RecognizedNumber) bool {
		return r.Number == 1 || r.Number == 4 || r.Number == 7 || r.Number == 8
	}
	onesFoursSevensAndEights := Filter(
		equalsOneFourSevenOrEight,
		recognizedNumbers,
	)
	return len(onesFoursSevensAndEights)
}

func SolveB() int {
	lines := files.ReadLines("day8/input")
	return ints.Sum(
		shared.Map(
			SolveForLine,
			lines,
		),
	)
}
