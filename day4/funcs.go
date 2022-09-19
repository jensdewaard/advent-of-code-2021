package day4

import (
	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
)

func HasWon(b Board) bool {
	Hline1 := b.Marked[0] && b.Marked[1] && b.Marked[2] && b.Marked[3] && b.Marked[4]
	Hline2 := b.Marked[5] && b.Marked[6] && b.Marked[7] && b.Marked[8] && b.Marked[9]
	Hline3 := b.Marked[10] && b.Marked[11] && b.Marked[12] && b.Marked[13] && b.Marked[14]
	Hline4 := b.Marked[15] && b.Marked[16] && b.Marked[17] && b.Marked[18] && b.Marked[19]
	Hline5 := b.Marked[20] && b.Marked[21] && b.Marked[22] && b.Marked[23] && b.Marked[24]

	Vline1 := b.Marked[0] && b.Marked[5] && b.Marked[10] && b.Marked[15] && b.Marked[20]
	Vline2 := b.Marked[1] && b.Marked[6] && b.Marked[11] && b.Marked[16] && b.Marked[21]
	Vline3 := b.Marked[2] && b.Marked[7] && b.Marked[12] && b.Marked[17] && b.Marked[22]
	Vline4 := b.Marked[3] && b.Marked[8] && b.Marked[13] && b.Marked[18] && b.Marked[23]
	Vline5 := b.Marked[4] && b.Marked[9] && b.Marked[14] && b.Marked[19] && b.Marked[24]

	return Hline1 || Hline2 || Hline3 || Hline4 || Hline5 || Vline1 || Vline2 || Vline3 || Vline4 || Vline5
}

func Mark(b Board, num int) Board {
	for i, v := range b.Numbers {
		if v == num {
			b.Marked[i] = true
		}
	}
	return b
}

func ParseBoards(ls []string) []Board {
	index := 0
	boards := make([]Board, 0)
	for index < len(ls) {
		boards = append(boards, NewBoard(ls[index:index+5]))
		index += 6
	}
	return boards
}

func WonBoards(bs []Board) []Board {
	return shared.Filter(
		func(b Board) bool {
			return HasWon(b)
		}, bs,
	)
}

func Score(b Board, lastCalledNumber int) int {
	sum := ints.SumIf(
		b.Numbers, shared.Map(
			func(s bool) bool {
				return !s
			}, b.Marked,
		),
	)
	return sum * lastCalledNumber
}
