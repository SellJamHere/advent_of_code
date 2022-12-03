package main

import (
	"fmt"
	"strings"
)

const (
	// part 1 mappings
	rock1     = "A"
	paper1    = "B"
	scissors1 = "C"

	rock2     = "X"
	paper2    = "Y"
	scissors2 = "Z"

	win  = 6
	lose = 0
	draw = 3

	// part 2 mappings
	mustLose = "X"
	mustDraw = "Y"
	mustWin  = "Z"
)

var (
	playMap = map[string]int{
		rock2:     1,
		paper2:    2,
		scissors2: 3,
	}

	// part 2 mapping
	scoreMap = map[string]int{
		mustLose: 0,
		mustDraw: 3,
		mustWin:  6,
	}
)

func main() {
	matches := strings.Split(puzzleInput, "\n")

	// part 1
	score := 0
	for _, match := range matches {
		plays := strings.Split(match, " ")
		playScore := play(plays[0], plays[1])

		score += playScore + playMap[plays[1]]
	}

	fmt.Printf("part 1 score: %d\n", score)

	// part 2
	score = 0
	for _, match := range matches {
		plays := strings.Split(match, " ")
		move := chooseMove(plays[0], plays[1])

		score += playMap[move] + scoreMap[plays[1]]
	}

	fmt.Printf("part 2 score: %d\n", score)
}

func play(a, b string) int {
	switch a {
	case rock1:
		switch b {
		case rock2:
			return draw
		case paper2:
			return win
		case scissors2:
			return lose
		}
	case paper1:
		switch b {
		case rock2:
			return lose
		case paper2:
			return draw
		case scissors2:
			return win
		}
	case scissors1:
		switch b {
		case rock2:
			return win
		case paper2:
			return lose
		case scissors2:
			return draw
		}
	}

	return 0
}

func chooseMove(a, desiredOutcome string) string {
	switch a {
	case rock1:
		switch desiredOutcome {
		case mustLose:
			return scissors2
		case mustDraw:
			return rock2
		case mustWin:
			return paper2
		}
	case paper1:
		switch desiredOutcome {
		case mustLose:
			return rock2
		case mustDraw:
			return paper2
		case mustWin:
			return scissors2
		}
	case scissors1:
		switch desiredOutcome {
		case mustLose:
			return paper2
		case mustDraw:
			return scissors2
		case mustWin:
			return rock2
		}
	}

	return ""
}

const puzzleInput = `C Z
C Z
A X
A X
B Z
B Z
B Z
A Z
B X
A X
A X
A X
C Z
C Z
C X
A X
A X
A X
C Z
B Z
C Z
A Y
B Z
A X
C Y
A X
A X
C Y
C Z
A Y
B Z
A X
C Y
B Z
B Z
B Z
A X
C X
C X
B Z
A X
C Z
A X
B Z
A Y
C X
A X
C Z
C Z
B Z
C Y
C X
C X
C X
C X
A Y
C Y
C Z
C Y
A X
C Y
A X
B Z
A Y
C X
A X
B Y
A X
C X
C X
C Z
A X
C X
A X
C X
B Z
A Z
B Y
B Z
B Z
A X
C Y
B X
A X
A X
B Z
A X
C X
C Z
C Z
A X
C Z
C Z
C Z
C X
C Z
A X
B Z
B Y
C Z
B Z
B Z
A X
B Z
C Y
C Z
A X
A Y
C Z
A X
B Y
C Y
C Z
A X
C Y
C X
A X
C X
A X
C Z
A X
C Y
A Y
C Z
C X
C Y
C Y
A Y
A Y
C Y
C Y
A Y
A X
C Z
C Z
C X
C X
C Z
A Y
C Z
C Z
A Y
A Y
B Z
A X
B Z
C Y
B Z
C Z
C Y
C Y
B Y
C X
A X
A X
A X
A Y
A X
C Z
C Z
C Z
A X
B Z
A X
B Z
C Y
C Z
A Y
C Y
A X
A X
C Z
B X
C Z
C Z
B Z
B Z
B Z
C Z
C X
C Z
A X
B Z
C Z
C X
A X
A Y
B Y
A X
B Z
A X
C Z
C Z
C X
B Z
C Z
C Z
C X
A X
A X
B Y
C Y
C Z
B Y
A X
B Z
A X
B X
B Z
C X
C X
C Y
C X
B Y
B Z
C X
A Y
C X
C Y
A X
C Z
C Z
C X
B Z
A Y
B Z
A Y
B X
A X
C Z
B Z
C X
C X
C X
A X
C X
B Z
C Z
B Y
B Z
A X
A X
C Z
B Z
A X
A Z
C Y
B Z
C X
C Z
A Y
C X
A X
A X
A Y
B Z
A X
B Z
A Y
B X
C X
C Y
A X
C Z
B X
C Z
C X
C X
C Z
A X
B Z
C Z
C X
C Z
B Y
A X
C X
C Y
A X
C Z
A X
B Z
C X
C X
B Z
C Z
C Z
C Y
A X
C Z
C X
A X
C Y
C Z
C Z
C Y
B Y
B Z
A X
B Y
C X
A X
A X
C Y
C Z
B Z
C X
C X
B Z
C Z
A X
A Y
A X
C Y
A Y
B Y
A X
A X
C X
C X
B Z
A Y
B Z
C Z
A X
A Y
A X
C X
A X
C Z
B Y
C Z
A X
B Z
C Z
C Z
A X
A X
B Z
B Z
B Z
A X
C X
A X
A X
C Y
C Y
C X
C X
A X
A Y
B Y
C Z
C Z
B Y
C X
C Z
A Y
C X
C Z
C X
B Z
C Y
C Z
C X
C Y
C Z
B Z
C Z
B Z
A X
B Y
B Y
C Z
B Z
C X
C Z
C Z
C Z
B Y
B Z
C Z
B Y
A X
A X
C X
A X
C X
A X
B Z
A X
A X
A X
C Z
A Y
A X
C Z
C X
A X
C X
A X
B Z
C Z
C Z
C Z
A X
A X
C Z
A X
B Z
C X
C Z
C Y
B Z
A X
B Z
A X
A Z
B Y
B Z
B Z
B Z
A Y
C Z
A X
B Y
C X
B X
B Z
C X
B Y
C Z
A X
C X
B Z
A X
B Z
A X
A X
B Y
C Y
C X
C X
C Z
A X
C Z
A X
B Y
C X
C Y
B Y
B Y
A Y
C X
A Y
C Z
C Y
A X
C Z
C Z
B Z
A X
A X
A X
A X
B Z
C Z
C X
C X
C Z
A X
B Z
B Y
C Y
A Y
C X
B Z
A Y
B Z
C Z
B Z
C X
A X
B Z
A X
A Y
C Y
B Z
B Z
C X
C Z
C Z
C Z
A X
B Z
A Y
A X
C Y
A Y
B Z
A Y
C Z
A Y
A Y
C X
A X
B Z
A Y
B Z
C X
A X
B Y
B Z
C Z
C Z
B Z
A X
A X
B Y
C Z
C Y
B Z
C Z
A Y
A X
C X
C Z
A X
A X
A Z
C X
C X
B X
B Z
A Y
C X
A Y
C X
C Z
A Y
B Z
A X
B X
A Y
C X
C Y
A Y
C X
B Y
C Y
A X
B Z
A X
C Z
A X
A X
C Z
C Z
B X
C X
A Y
B Z
C Z
A X
A X
C Z
C Y
A X
C X
C Z
C Y
C X
C Y
C X
C Z
C Z
B Z
A X
A Y
B Z
A X
B Z
C X
B Y
A X
A X
A X
A X
B Z
C Z
B Z
B Y
C Z
C X
C Z
C X
B Z
C Z
A X
C Z
C Z
A Y
C X
A X
B Z
A X
B Z
B Y
A X
A X
A X
A Y
C Z
A X
B Y
A Y
B Y
C Y
A X
A X
C Y
A X
C Z
B Z
C X
C Z
A X
B Z
C Z
B Z
B Z
B Z
C X
C Z
A X
A Y
C X
C Z
B Z
B Y
C Z
B Y
B Z
C Z
A Y
B Z
C X
C Y
A X
B Y
A X
C Z
C X
C Z
A X
C X
B Z
B X
C Z
B Z
A Y
A X
A Y
C Z
C Y
A X
B X
A Z
C Y
C X
C Y
C X
C Y
B Z
C X
B Y
C Z
C Y
B Z
A Y
B Y
C Z
C X
B Z
A Y
C Z
C Z
C Z
C Y
C X
C X
B Y
C Z
A Y
C Z
C X
B Y
C X
B Z
B Z
B Z
B Z
C Y
C Z
C X
B X
C Z
C Z
C Z
C X
B Y
B Z
C X
A X
C Y
B X
C X
C X
C Z
A X
A X
B Z
B Z
A X
C Z
A X
A Y
C X
B X
C Y
C Y
A X
A X
C Z
C Z
B Y
B Y
B Z
C X
A X
A Z
C Z
B Y
C Y
C X
C Y
B Z
C X
A X
A X
C Z
C Y
B X
C Z
B Z
A X
C X
A Y
C Y
C Z
B Z
C Z
A X
C Y
B Z
C Y
B X
C Z
C Z
A X
A X
C X
C Z
C X
C Y
C Z
A Y
C X
C Z
B Y
B Y
C Z
A X
C Y
A X
C X
C X
B Y
C Z
C Z
A X
A Y
C Z
B Z
C Y
A X
A X
C Z
C Z
A X
C Z
C Z
B Z
B Z
A X
A Y
A X
C Z
C X
A X
B X
C Z
C Z
B Y
A X
B Z
C X
C Z
C Z
C Z
C X
A Y
A X
A Y
A X
A Y
C Z
C Z
B Z
B Y
A X
B Z
C Z
C X
C Y
C Z
A X
C Y
B Y
B Z
C Z
B Y
C X
C Z
B Y
B Z
B Y
C X
C Z
C X
C Z
B Y
B Z
C X
A X
A Y
B Z
B Y
A X
A X
C Z
A X
A X
C X
A X
C Z
C X
A X
B X
A Y
C X
A X
C Y
A X
C Z
A X
A X
C Z
A X
A X
C Z
A X
B Y
B Z
A X
C X
A X
C X
B Z
C Y
A Y
C Z
B Z
A X
C X
A X
B Y
C X
B Y
B Z
A X
B Z
C X
B Z
A X
A X
C X
A X
C X
A Y
C Z
A X
B Y
B Z
C Z
C Z
A X
C Y
B Z
C X
B Z
B Z
C Z
A X
C X
C X
A Y
A X
C Z
B Z
A X
B X
A X
B Z
B Z
C Z
B Y
C Z
A X
C Z
C Z
C Z
C X
A X
A X
C Y
B Y
C Z
B Y
A X
B Z
A Y
C Y
B Y
C Z
C X
C X
A X
B Y
A X
A X
C Z
C X
C Y
A X
A X
A X
C X
B Z
B Y
A X
B Y
C X
C Z
B Z
A X
C X
C Z
B Y
A X
C Z
B Z
B Z
A X
B Y
A X
B Y
C Z
C Z
B Z
C Z
C Y
A X
B Z
C Z
A Y
C Z
B Z
B Z
C Z
B Y
C Z
C Z
B Z
A Y
C X
C Y
B Z
C X
C Z
A X
A X
C Y
A Z
C Z
C X
B Z
B Z
C Y
A X
A Y
C Z
A X
A X
B Z
A X
C X
C X
B X
B X
A Z
B Y
A X
A Y
C Z
A X
A X
B Z
B Z
C Y
A X
C Y
C Z
A Y
C Y
B X
C X
A X
B Z
A X
A X
A X
A X
A X
B X
A X
C Z
B Z
B Z
B X
A X
C Z
B Z
C Z
C Z
C Y
A X
C Y
C Y
C Z
A Y
C Z
A X
C X
C Z
A X
A Y
C Z
C X
A X
C X
B Y
C Z
B Z
A X
C Y
B Z
B Y
A X
C Z
A X
B Z
C Z
C Y
C Z
A X
A X
B Z
B Y
A X
C X
C Y
B X
C Z
A Y
C X
B Y
B Z
A Y
C Y
C Z
C Z
C X
A X
A X
C X
C X
C Z
C X
C X
B X
B Z
B Z
C X
C Z
C X
A X
A X
A X
C Z
C Y
B Z
C X
C X
A X
C X
C Y
A X
B Y
A X
C Z
C Z
B Z
A X
A X
A X
B Z
C X
A X
B Y
B Z
C X
B Y
C X
B Y
C Z
B X
A X
A X
B Z
A X
A X
A X
B Z
A X
C Z
C X
C Z
A X
C Y
C Y
C Z
C Z
C Z
A X
C Z
C X
C Z
C Z
A X
B Z
C X
C Y
B X
B Z
B Z
A X
C Y
B Y
C Z
B Z
C Z
C X
A Z
C Z
B X
C X
C Z
C Z
C X
C Z
C Z
A X
B Z
C Z
C X
A X
B Z
C Y
C Y
B Y
C X
A X
A X
C Z
B Z
C X
C X
C X
B Y
C Z
B X
C X
C Z
A X
A X
C Z
C Z
C X
C X
C Z
A X
C Z
C X
C Z
C Z
A X
C Z
A X
A X
A X
B Z
A X
B Z
C Z
C Z
A X
C Z
C Z
B Z
C Z
A Y
C X
A X
A X
C X
A X
C X
C X
A Y
B Y
C Z
B Z
C X
B Y
C Y
C X
C Z
A X
C Y
B Z
B Y
A X
B Y
B Y
A X
C X
C X
A X
A Y
C Y
B Z
A X
B Z
A Y
B Y
C Z
C Z
C Z
A Y
B Y
B Y
C Z
B Z
C Z
A X
C Y
A X
B Y
B Z
C X
C Y
A X
B Z
B Z
A X
B Z
C Y
A X
C X
C Z
C Z
A X
C Y
A X
C X
C X
B Z
B Y
A X
C Y
B Z
B Y
A X
C Z
A X
B Y
A X
A X
A X
C X
A X
C Z
B Z
C Z
A X
A X
B Y
C X
C Y
C Z
A X
C Z
B Z
A X
B Z
A X
B Y
B Z
C Z
C Z
C Y
C X
A X
C Z
A Y
A X
B Z
B Y
A X
C X
C Z
B Z
A X
C Z
C X
C Y
A X
A X
C Z
A X
A X
C X
C X
A X
C X
C Z
A X
A X
C Y
A Y
B Z
C X
C X
B Y
B Z
A X
B Z
C Z
C Z
A Y
C Y
C Z
A X
C Z
A X
C Y
C X
B X
B Z
A X
A X
B Z
A X
C X
C X
A Y
A X
A X
C Y
B Z
A X
B Y
C X
B Z
C Y
B Y
A X
C Z
B Y
B Z
A X
B Y
C X
C Z
A X
B Z
A X
C Y
C Z
A X
A X
C Y
C Y
C X
A Y
C Z
B Z
A X
A Y
C Z
C Z
C Z
C Z
B Y
A Y
A X
A X
C X
A X
B Y
C Z
B Z
A Y
C X
C X
A Y
A X
A X
C X
C Z
C X
A X
A Y
C X
C Z
C X
A X
A X
C Z
C X
A X
C X
C Z
A X
A X
C Z
B Y
A X
C Z
C Y
C Z
C Z
B X
B Z
B X
A X
C Z
B X
C Z
A X
C Z
A Y
A Z
B Z
C X
A X
A X
A X
A X
A X
C Z
A X
B X
A Y
A X
C X
A X
A Y
A Y
B Y
C X
A Y
A X
C Z
C Z
C Y
B Z
C X
C Z
A Y
C X
C Y
C X
A X
B Y
C Z
C Z
A X
C Z
B Z
B Z
C X
C X
A X
B Y
C X
C X
A X
C Z
B Y
C Y
C Z
A X
B Z
C Z
B Y
C X
A X
B Y
B Z
A X
C Y
B Z
C Z
C X
A X
C X
A X
B Y
C X
A X
C Y
B Z
B Z
B Z
C Y
A X
A X
A Y
C Z
B X
B Z
A X
A X
C X
A X
C X
A X
C X
A X
A X
A X
A Y
A Y
B Y
A X
B Y
A X
B Y
B Z
B Y
A Y
C Z
C Z
C Z
C Z
A X
C Z
C X
C Y
C Z
B Z
B Z
B Z
C Y
C Z
B Z
C X
A Y
A X
C X
C Y
A Y
B Y
C X
C X
B Y
C Z
A X
C X
C X
C X
C X
C X
C X
B Z
C Y
A X
A X
B Y
B Z
A X
C X
C Z
A X
B X
B Z
C Y
B Z
C Z
C Y
A X
C Z
B Y
C X
A X
A Y
A X
C X
A Y
A Y
C Z
B Z
C Z
C X
C Y
B Y
A X
C X
C Z
B Y
A X
B X
B Z
C Z
B Z
B Z
A X
A Y
A X
A X
B Z
C Z
A X
A Y
C Z
C Z
C Y
A X
B Z
C X
C X
A Y
A X
C X
C Y
C Z
C Y
C X
A X
B Z
C Y
B Z
C X
A X
A X
C X
C Z
C Y
C Y
A X
A Y
C Z
A Y
B Y
C Y
A Y
C Z
B Y
C X
A X
B Z
C Z
A X
A X
B Z
C X
B Z
C Z
C Z
A Y
A X
C X
A X
B Z
C X
A X
A X
B Y
A Z
B Z
C X
C Z
C X
C Z
C X
A X
C Z
B Y
C Y
C Z
A X
C Z
A X
A Z
A X
A X
C Z
A X
B Z
C X
B Z
A X
C X
A X
C X
C Z
B Z
A Y
C Y
C Z
C Z
C X
C Z
A X
A X
B Y
B Z
C Y
B Z
A Y
A X
A X
B Z
B X
A X
B Y
B Z
C X
A X
A X
C Z
B Y
C X
B Y
A Y
A Z
A X
B Y
C X
C Z
C Z
C Y
C Z
A X
B X
C X
A Y
A X
C Z
C Y
C X
B Z
A X
C X
C Z
A X
C Z
A X
C X
C X
A Y
C Z
A X
C X
C Z
C Y
A X
A Y
A Y
C Z
C X
A X
C Z
A Y
A Y
A X
B Z
C Z
B Y
C Z
B Y
C X
C Z
B Y
A X
A X
B Z
A X
C Z
A X
A X
B Y
A X
B Z
A Y
A X
A Y
C Z
B Z
A X
B X
B Y
B Z
A X
A Z
B Y
B Z
C X
B Z
A X
B Z
C Z
B Y
C X
A X
A Y
B Y
A Y
C X
C X
B Y
A Y
B Z
C Z
B Z
A Y
A X
B Z
B Z
A Y
C Z
C Z
B Z
C X
B Z
B Y
B Z
B Z
A X
A X
C X
C Z
B Z
A X
B Z
C Y
C Z
A Z
A X
A X
C Z
A X
C Z
C Z
C Z
B Z
A X
A Y
B X
A X
A X
A X
C Z
B X
B X
A X
C Z
A X
B X
A X
A X
B Z
A X
C Z
C Y
A X
B Z
A X
C Z
B Z
A X
A X
B Y
B Y
C Y
A X
B Z
C Y
A X
C Z
B Z
A X
C Z
B Z
C X
B X
B Z
C Z
A X
B Z
A X
A X
A X
A X
C Z
B Z
B Y
C Z
A X
C Z
A X
C X
C X
A X
B Y
C X
B Z
A Y
C Y
B Z
C Y
C X
A X
B Z
C X
B Y
A X
A Y
B Y
B Y
A X
C X
C Z
B Z
A Y
A X
B Y
C X
B Z
B X
C Y
B X
C X
B Y
C X
A X
C Z
B Z
B Y
B Y
A X
A X
A Z
B Z
B Z
C X
C Y
A Y
C Z
C X
A X
C Z
C X
B Z
A Y
B Z
C Z
B Z
B Z
C X
B Z
C X
C X
B Z
A X
B Y
B Z
A X
C X
A X
C Y
A X
C X
B X
C Z
C Y
C Z
B Z
C X
A X
A X
A Y
C Z
C Y
A Y
C Z
B Z
C Z
C Y
A X
A X
C X
C Z
C Z
A Z
C Z
A X
B Z
C X
C Z
A X
B Z
C X
A X
A Y
B Y
C X
A X
A X
C X
C Y
B Z
C X
A X
C X
B Z
B Y
A Y
C Z
A X
C Z
C Z
A X
C Y
B Z
A Y
B Y
B Y
A X
A X
C X
C X
B Y
C Z
C Z
C Z
C X
C Y
C Z
B Z
C Z
A Y
C Z
A X
B Z
B Y
B Z
A X
A X
A X
A X
C X
C Z
B Z
C X
B Z
C Z
A X
C X
C Z
B Y
A X
C X
A X
B Z
B Z
A X
A X
B Y
C Y
C Y
C X
A X
B Z
C Y
B Z
A Y
B Y
B Y
A X
B Z
A X
C X
C Z
A X
B Z
A Y
C Y
B Z
C Z
C Y
A X
C Y
A X
C X
B Z
C X
A X
C Z
A X
B Y
B Z
C X
C Y
B X
A X
B Z
A X
B Y
C X
C X
C Y
C Y
C X
B Y
C Z
C X
B Z
B Y
C X
A Y
C Z
C Z
C X
B X
B Y
A X
A X
C Z
C X
C Y
A X
C Y
A Y
C Z
C X
A X
C Z
C X
A X
B X
C X
C X
C X
B Y
B Z
C X
C Z
A Y
B Y
C X
C X
A X
B Z
C Z
A X
C Y
C Z
A Y
B Z
C X
C Z
A X
C Z
B Z
B Y
B Z
A X
C X
A Y
C X
C Z
B Y
C Y
C Z
C Z
A X
A X
A X
C Z
B Y
C Z
A X
A X
B Z
B Y
B Z
C Z
B Z
A Y
C Y
C Z
B Y
A X
A X
C Z
B Y
C Z
A X
C X
B Y
A Y
B Z
A X
A X
A X
B Y
A X
C Z
B Y
C Z
B Y
C Z
C X
C Y
C X
A X
A X
A X
C Z
C X
C X
B Y
A Y
B Z
B X
C Z
B Y
A X
C Y
B Z
C X
A X
A X
A X
A X
C X
B Y
A X
A X
B Y
A X
B Z
C Z
A X
A X
A X
A X
B Z
A X
C Z
C X
C Y
B Z
C Z
A X
C Y
C Z
A X
A X
C X
B Z
C X
B Z
C Z
A X
A X
A Y
B Y
C Z
B Y
A X
C X
C Z
C Z
C X
A X
A X
C Z
A X
A Y
A X
A X
C Z
C X
C X
C Z
C X
A X
C Z
C Z
B Z
A Y
A Y
B Y
A X
A X
C X
C X
A X
A X
C X
C X
B Y
C Z
A X
B Y
A X
A X
A X
C X
C Y
B Z
B Y
B Z
C Z
C Z
C X
C Z
A X
B Z
A Y
C Z
B Z
A X
C Y
B Z
C Z
C Y
A X
B Y
C Z
A X
A X
A X
C Z
C Z
C X
A Y
C X
B Z
B Y
A X
C X
C X
A Y
A Y
A X
B X
B Z
B Y
B Y
A X
C Y
A X
C Z
C X
C X
C Z
B Z
B Z
C Y
C X
B Z
C Z
A Y
C Y
A X
B Z
A X
C Y
B Z
B Y
C Z
A X
A X
C Z
B Z
B Z
C Z
C Y
C Z
C Z
C X
A Y
A Y
B Z
C Z
B Y
C X
C X
A Z
C Z
A X
A X
C X
A Y
C X
A X
A X
A X
C Y
A X
B Y
A X
B Y
A X
A X
A Y
C X
C Z
A X
C X
B Y
B Z
B Z
A Y
C Z
C X
C X
B Z
B Z
C X
B Y
A X
B Y
A X
A X
C Y
B X
C X
A Z
A Z
A X
C Z
C Y
C Z
C Y`