package day13

import (
	"fmt"

	"github.com/jensdewaard/advent-of-code-2021/shared"
	"github.com/jensdewaard/advent-of-code-2021/shared/files"
	"github.com/jensdewaard/advent-of-code-2021/shared/grids"
	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func SolveA() int {
	pos, folds := ParseInput(files.ReadLines("day13/input"))
	paper := grids.NewGrid(0, 0)
	for _, p := range pos {
		paper = AddDot(paper, p)
	}
	if folds[0].Type == "h" {
		paper = FoldPaperVertically(paper, folds[0].Line)
	} else if folds[0].Type == "v" {
		paper = FoldPaperHorizontally(paper, folds[0].Line)
	}
	foldedGrid := grids.FoldGridLines(0, ints.Add, paper)
	return shared.Fold(0, ints.Add, foldedGrid)
}

func SolveB() int {
	pos, folds := ParseInput(files.ReadLines("day13/input"))
	paper := grids.NewGrid(0, 0)
	for _, p := range pos {
		paper = AddDot(paper, p)
	}
	for _, f := range folds {
		if f.Type == "h" {
			paper = FoldPaperVertically(paper, f.Line)
		} else if f.Type == "v" {
			paper = FoldPaperHorizontally(paper, f.Line)
		}
	}
	fmt.Print(SprintPaper(paper))
	return shared.Fold(0, ints.Add, grids.FoldGridLines(0, ints.Add, paper))
}

func ParseInput(ss []string) ([]grids.Position, []Fold) {
	ps := make([]grids.Position, 0)
	fs := make([]Fold, 0)
	for _, s := range ss {
		if s == "" {
			continue
		}
		if s[0] == 'f' { // fold
			locs := strings.Split(strings.Split(s, " along ")[1], "=")
			if locs[0] == "y" {
				fs = append(fs, Fold{strings.ParseInt(locs[1]), "h"})
			} else {
				fs = append(fs, Fold{strings.ParseInt(locs[1]), "v"})
			}
		} else {
			locs := strings.Split(s, ",")
			ps = append(
				ps, grids.Position{
					X: strings.ParseInt(locs[0]),
					Y: strings.ParseInt(locs[1]),
				},
			)
		}
	}
	return ps, fs
}
