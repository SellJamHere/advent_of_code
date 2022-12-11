package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(puzzleInput, "\n")

	// part 1
	clock := 1
	xReg := 1
	i := 0
	signal := 0
	var currentOp Operator
	for {
		if currentOp == nil {
			if i > len(lines)-1 {
				break
			}

			line := lines[i]
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "noop":
				currentOp = NoOp{}
			case "addx":
				amount, _ := strconv.Atoi(parts[1])
				currentOp = &AddOp{amount: amount}
			}

			i++
		}

		if clock%40 == 20 {
			signal += clock * xReg
		}

		if !currentOp.Cycle(&xReg) {
			currentOp = nil
		}

		clock++
	}

	fmt.Printf("part 1 xreg: %d clock: %d, strength: %d\n", xReg, clock, signal)

	// part 2
	screen := NewScreen(40, 6)

	clock = 1
	xReg = 1
	i = 0
	signal = 0
	currentOp = nil
	for {
		if currentOp == nil {
			if i > len(lines)-1 {
				break
			}

			line := lines[i]
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "noop":
				currentOp = NoOp{}
			case "addx":
				amount, _ := strconv.Atoi(parts[1])
				currentOp = &AddOp{amount: amount}
			}

			i++
		}

		if clock%40 == 20 {
			signal += clock * xReg
		}

		screen.Draw(xReg, clock)

		if !currentOp.Cycle(&xReg) {
			currentOp = nil
		}

		clock++
	}

	fmt.Println(screen)
}

type Operator interface {
	Cycle(x *int) bool // returns true when processing, false when finished
}

type NoOp struct{}

func (n NoOp) Cycle(x *int) bool {
	return false
}

type AddOp struct {
	amount int
	cycle  int
}

func (a *AddOp) Cycle(x *int) bool {
	if a.cycle < 1 {
		a.cycle++
		return true
	}

	*x += a.amount

	return false
}

type Screen [][]string

func NewScreen(x, y int) Screen {
	s := make([][]string, y)
	for i := 0; i < y; i++ {
		s[i] = make([]string, x)
		for j := 0; j < x; j++ {
			s[i][j] = "."
		}
	}

	return s
}

func (s *Screen) Draw(spriteLocation, cycle int) {
	cycle--
	row := cycle / 40
	col := cycle % 40

	if math.Abs(float64(spriteLocation-col)) <= 1 {
		(*s)[row][col] = "#"
	}
}

func (s Screen) String() string {
	builder := strings.Builder{}
	for _, line := range s {
		for _, col := range line {
			builder.WriteString(string(col))
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

const puzzleInput2 = `noop
addx 3
addx -5`

const puzzleInput1 = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

const puzzleInput = `noop
noop
addx 5
addx 3
noop
addx 14
addx -12
noop
addx 5
addx 1
noop
addx 19
addx -15
noop
noop
noop
addx 7
addx -1
addx 4
noop
noop
addx 5
addx 1
addx -38
noop
addx 21
addx -18
addx 2
addx 2
noop
addx 3
addx 5
addx -6
addx 11
noop
addx 2
addx 19
addx -18
noop
addx 8
addx -3
addx 2
addx 5
addx 2
addx 3
addx -2
addx -38
noop
addx 3
addx 4
addx 5
noop
addx -2
addx 5
addx -8
addx 12
addx 3
addx -2
addx 5
addx 11
addx -31
addx 23
addx 4
noop
noop
addx 5
addx 3
addx -2
addx -37
addx 1
addx 5
addx 2
addx 12
addx -10
addx 3
addx 4
addx -2
noop
addx 6
addx 1
noop
noop
noop
addx -2
addx 7
addx 2
noop
addx 3
addx 3
addx 1
noop
addx -37
addx 2
addx 5
addx 2
addx 32
addx -31
addx 5
addx 2
addx 9
addx 9
addx -15
noop
addx 3
addx 2
addx 5
addx 2
addx 3
addx -2
addx 2
addx 2
addx -37
addx 5
addx -2
addx 2
addx 5
addx 2
addx 16
addx -15
addx 4
noop
addx 1
addx 2
noop
addx 3
addx 5
addx -1
addx 5
noop
noop
noop
noop
addx 3
addx 5
addx -16
noop`
