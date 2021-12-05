package day4

import (
	"io/ioutil"
	"log"
	"strings"

	mystrings "github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	inputData, _ := ioutil.ReadFile("day4/input")
	input := strings.Split(string(inputData), "\n")
	numbersDrawn := mystrings.MapToInt(
		mystrings.ParseInt,
		strings.Split(input[0], ","),
	)

	boards := ParseBoards(input[2:])
	index := 0

	for len(WonBoards(boards)) == 0 {
		markNumber := func(b Board) Board {
			return Mark(b, numbersDrawn[index])
		}
		boards = MapBoards(markNumber, boards)
		index++
	}
	wonBoards := WonBoards(boards)
	winningBoard := wonBoards[0]
	return Score(winningBoard, numbersDrawn[index-1])
}

func SolveB() int {
	inputData, _ := ioutil.ReadFile("day4/input")
	input := strings.Split(string(inputData), "\n")
	numbersDrawn := mystrings.MapToInt(
		mystrings.ParseInt,
		strings.Split(input[0], ","),
	)

	boards := ParseBoards(input[2:])
	index := 0

	for len(boards) > 1 {
		markNumber := func(b Board) Board {
			return Mark(b, numbersDrawn[index])
		}
		boards = MapBoards(markNumber, boards)
		boards = FilterBoards(func(b Board) bool {
			return !HasWon(b)
		}, boards)
		index++
	}
	lastWinner := boards[0]
	for !HasWon(lastWinner) {
		lastWinner = Mark(lastWinner, numbersDrawn[index])
		index++
	}
	if !HasWon(lastWinner) {
		log.Fatal("hold up the board hasn't won")
	}
	return Score(lastWinner, numbersDrawn[index-1])

}
