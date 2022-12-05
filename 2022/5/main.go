package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// part 1
	stacks, moves := parseInput(puzzleInput)

	for _, move := range moves {
		for i := 0; i < move.qty; i++ {
			val := stacks[move.from-1].pop()
			stacks[move.to-1].push(val)
		}
	}

	result := ""
	for _, stack := range stacks {
		result += stack.peek()
	}

	fmt.Printf("part 1 result: %s\n", result)

	// part 2
	stacks, moves = parseInput(puzzleInput)

	for _, move := range moves {
		val := stacks[move.from-1].remove(move.qty)
		stacks[move.to-1].add(val)
	}

	result = ""
	for _, stack := range stacks {
		result += stack.peek()
	}

	fmt.Printf("part 2 result: %s\n", result)
}

type move struct {
	qty  int
	from int
	to   int
}

func parseInput(i string) ([]stack, []move) {
	lines := strings.Split(i, "\n")

	var initialCreateState []string
	var moveLines []string
	for i, line := range lines {
		if line == "" {
			initialCreateState = lines[:i]
			moveLines = lines[i+1:]
			break
		}
	}

	stackLine := initialCreateState[len(initialCreateState)-1]
	stackNumbers := strings.Fields(stackLine)

	stacks := parseStacks(len(stackNumbers), initialCreateState)

	var moves []move
	for _, m := range moveLines {
		m = strings.TrimPrefix(m, "move ")
		parts1 := strings.Split(m, " from ")
		parts2 := strings.Split(parts1[1], " to ")
		qty, _ := strconv.Atoi(parts1[0])
		from, _ := strconv.Atoi(parts2[0])
		to, _ := strconv.Atoi(parts2[1])
		moves = append(moves, move{
			qty:  qty,
			from: from,
			to:   to,
		})
	}

	return stacks, moves
}

func parseStacks(stackCount int, lines []string) []stack {
	stacks := make([]stack, stackCount)

	for i := len(lines) - 2; i >= 0; i-- {

		line := lines[i]
		j := 0
		for len(line) > 2 {
			next := line[:2]
			nextVal := string(next[1])
			if len(strings.TrimSpace(nextVal)) > 0 {
				stacks[j].push(nextVal)
			}

			if len(line) > 4 {
				line = line[4:]
			} else {
				line = ""
			}
			j++
		}
	}

	return stacks
}

type stack []string

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) pop() string {
	length := len(*s)
	if length == 0 {
		return ""
	}

	top := (*s)[length-1]
	*s = (*s)[:length-1]

	return top
}

func (s *stack) add(n []string) {
	*s = append(*s, n...)
}

func (s *stack) remove(n int) []string {
	length := len(*s)
	if length == 0 {
		return nil
	}

	top := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]

	return top
}

func (s *stack) peek() string {
	return (*s)[len(*s)-1]
}

const puzzleInput1 = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

const puzzleInput = `                [B] [L]     [J]    
            [B] [Q] [R]     [D] [T]
            [G] [H] [H] [M] [N] [F]
        [J] [N] [D] [F] [J] [H] [B]
    [Q] [F] [W] [S] [V] [N] [F] [N]
[W] [N] [H] [M] [L] [B] [R] [T] [Q]
[L] [T] [C] [R] [R] [J] [W] [Z] [L]
[S] [J] [S] [T] [T] [M] [D] [B] [H]
 1   2   3   4   5   6   7   8   9 

move 5 from 4 to 5
move 2 from 5 to 8
move 2 from 9 to 1
move 2 from 9 to 1
move 1 from 5 to 3
move 10 from 5 to 8
move 1 from 4 to 7
move 1 from 1 to 2
move 5 from 3 to 7
move 1 from 2 to 8
move 21 from 8 to 5
move 13 from 5 to 7
move 2 from 9 to 4
move 1 from 7 to 4
move 5 from 1 to 4
move 1 from 5 to 7
move 2 from 2 to 7
move 1 from 3 to 2
move 1 from 1 to 6
move 7 from 5 to 9
move 16 from 7 to 4
move 7 from 9 to 3
move 1 from 7 to 5
move 1 from 3 to 8
move 3 from 2 to 7
move 1 from 8 to 9
move 3 from 3 to 6
move 21 from 4 to 9
move 1 from 5 to 7
move 4 from 4 to 9
move 8 from 6 to 3
move 6 from 7 to 1
move 12 from 9 to 8
move 6 from 7 to 2
move 3 from 6 to 5
move 1 from 6 to 9
move 4 from 8 to 6
move 3 from 8 to 5
move 4 from 1 to 8
move 4 from 6 to 1
move 2 from 1 to 3
move 1 from 5 to 8
move 2 from 2 to 8
move 5 from 8 to 3
move 4 from 2 to 7
move 5 from 8 to 1
move 2 from 1 to 7
move 1 from 8 to 2
move 2 from 1 to 7
move 11 from 9 to 2
move 1 from 8 to 5
move 2 from 9 to 4
move 3 from 9 to 5
move 2 from 5 to 1
move 6 from 5 to 8
move 2 from 4 to 2
move 1 from 5 to 6
move 7 from 1 to 8
move 2 from 2 to 7
move 13 from 8 to 1
move 16 from 3 to 1
move 3 from 2 to 1
move 12 from 7 to 6
move 15 from 1 to 8
move 2 from 3 to 8
move 16 from 1 to 2
move 24 from 2 to 8
move 1 from 1 to 5
move 1 from 5 to 8
move 3 from 6 to 7
move 26 from 8 to 3
move 20 from 3 to 9
move 1 from 2 to 9
move 16 from 9 to 3
move 14 from 3 to 1
move 13 from 1 to 6
move 3 from 3 to 4
move 3 from 9 to 4
move 1 from 7 to 8
move 5 from 8 to 2
move 8 from 8 to 5
move 18 from 6 to 1
move 4 from 8 to 5
move 6 from 4 to 1
move 2 from 2 to 5
move 5 from 3 to 8
move 5 from 8 to 7
move 2 from 5 to 8
move 5 from 5 to 4
move 3 from 2 to 8
move 22 from 1 to 2
move 1 from 1 to 2
move 5 from 8 to 2
move 2 from 5 to 2
move 1 from 1 to 6
move 5 from 5 to 2
move 1 from 9 to 8
move 5 from 4 to 1
move 6 from 6 to 9
move 3 from 1 to 9
move 1 from 1 to 7
move 8 from 9 to 6
move 6 from 7 to 1
move 5 from 6 to 5
move 27 from 2 to 1
move 4 from 5 to 7
move 9 from 1 to 5
move 1 from 9 to 1
move 3 from 6 to 2
move 9 from 2 to 1
move 2 from 7 to 2
move 1 from 8 to 7
move 10 from 5 to 9
move 1 from 9 to 7
move 25 from 1 to 8
move 6 from 7 to 4
move 11 from 1 to 7
move 3 from 8 to 1
move 3 from 2 to 6
move 3 from 8 to 9
move 11 from 8 to 6
move 1 from 2 to 6
move 12 from 6 to 4
move 13 from 4 to 5
move 1 from 6 to 1
move 3 from 7 to 5
move 5 from 8 to 7
move 1 from 7 to 1
move 5 from 1 to 6
move 3 from 6 to 4
move 3 from 8 to 6
move 2 from 5 to 2
move 12 from 5 to 9
move 5 from 6 to 2
move 2 from 5 to 9
move 6 from 4 to 9
move 11 from 7 to 3
move 1 from 2 to 5
move 1 from 7 to 8
move 1 from 5 to 7
move 1 from 7 to 1
move 1 from 8 to 1
move 2 from 4 to 7
move 2 from 6 to 8
move 5 from 3 to 6
move 2 from 7 to 2
move 2 from 2 to 9
move 1 from 2 to 9
move 1 from 1 to 6
move 35 from 9 to 7
move 2 from 8 to 7
move 3 from 3 to 8
move 5 from 2 to 4
move 3 from 3 to 7
move 2 from 4 to 7
move 4 from 6 to 5
move 4 from 5 to 9
move 3 from 4 to 5
move 1 from 8 to 3
move 4 from 9 to 8
move 1 from 9 to 6
move 38 from 7 to 2
move 1 from 3 to 5
move 1 from 1 to 7
move 4 from 7 to 3
move 3 from 6 to 1
move 22 from 2 to 7
move 1 from 5 to 8
move 7 from 8 to 4
move 8 from 2 to 8
move 3 from 5 to 1
move 4 from 3 to 9
move 1 from 8 to 3
move 1 from 3 to 7
move 2 from 2 to 3
move 5 from 8 to 9
move 3 from 9 to 1
move 2 from 1 to 7
move 6 from 2 to 3
move 6 from 3 to 1
move 2 from 3 to 6
move 1 from 6 to 1
move 14 from 7 to 2
move 4 from 1 to 6
move 8 from 1 to 3
move 4 from 3 to 6
move 3 from 9 to 5
move 1 from 8 to 6
move 1 from 8 to 4
move 9 from 7 to 1
move 8 from 2 to 4
move 4 from 2 to 9
move 2 from 2 to 1
move 3 from 5 to 8
move 1 from 8 to 6
move 1 from 7 to 8
move 1 from 6 to 5
move 3 from 9 to 5
move 2 from 9 to 5
move 4 from 3 to 9
move 3 from 6 to 3
move 3 from 6 to 9
move 9 from 4 to 1
move 1 from 9 to 8
move 3 from 3 to 6
move 2 from 7 to 4
move 4 from 8 to 5
move 7 from 5 to 6
move 19 from 1 to 9
move 5 from 9 to 3
move 2 from 1 to 6
move 1 from 4 to 6
move 4 from 3 to 2
move 21 from 9 to 7
move 1 from 1 to 2
move 1 from 9 to 1
move 1 from 1 to 8
move 16 from 7 to 6
move 24 from 6 to 5
move 7 from 4 to 5
move 1 from 8 to 3
move 2 from 2 to 8
move 31 from 5 to 8
move 1 from 4 to 6
move 2 from 6 to 9
move 1 from 7 to 4
move 3 from 7 to 9
move 1 from 4 to 8
move 2 from 3 to 5
move 1 from 2 to 3
move 1 from 3 to 7
move 1 from 7 to 9
move 24 from 8 to 6
move 1 from 8 to 1
move 30 from 6 to 1
move 2 from 5 to 2
move 1 from 6 to 9
move 6 from 9 to 7
move 1 from 6 to 4
move 1 from 4 to 6
move 23 from 1 to 3
move 21 from 3 to 4
move 4 from 2 to 6
move 3 from 6 to 1
move 1 from 5 to 1
move 4 from 1 to 9
move 3 from 9 to 6
move 8 from 1 to 6
move 4 from 8 to 5
move 2 from 7 to 5
move 7 from 4 to 3
move 3 from 4 to 9
move 9 from 3 to 9
move 1 from 7 to 6
move 6 from 5 to 8
move 14 from 6 to 2
move 4 from 8 to 4
move 7 from 4 to 5
move 1 from 7 to 9
move 6 from 4 to 3
move 13 from 2 to 6
move 5 from 3 to 7
move 1 from 3 to 8
move 1 from 8 to 2
move 4 from 8 to 3
move 6 from 6 to 4
move 2 from 2 to 8
move 5 from 4 to 7
move 3 from 7 to 5
move 1 from 7 to 9
move 2 from 3 to 9
move 3 from 7 to 3
move 1 from 7 to 9
move 1 from 7 to 9
move 3 from 4 to 1
move 6 from 6 to 1
move 2 from 7 to 5
move 1 from 3 to 5
move 11 from 9 to 4
move 9 from 4 to 5
move 3 from 3 to 4
move 1 from 3 to 9
move 2 from 8 to 1
move 9 from 1 to 8
move 22 from 5 to 8
move 2 from 1 to 3
move 3 from 4 to 6
move 14 from 8 to 9
move 1 from 3 to 9
move 19 from 9 to 3
move 3 from 9 to 4
move 2 from 7 to 2
move 1 from 4 to 6
move 1 from 3 to 8
move 8 from 3 to 1
move 2 from 9 to 6
move 1 from 2 to 5
move 3 from 4 to 9
move 1 from 2 to 3
move 20 from 8 to 3
move 4 from 9 to 5
move 1 from 4 to 2
move 26 from 3 to 5
move 1 from 8 to 3
move 8 from 1 to 4
move 1 from 3 to 7
move 1 from 2 to 1
move 1 from 1 to 6
move 1 from 6 to 7
move 4 from 5 to 3
move 3 from 4 to 2
move 5 from 5 to 3
move 2 from 2 to 6
move 3 from 3 to 5
move 2 from 4 to 8
move 5 from 3 to 9
move 5 from 9 to 8
move 19 from 5 to 9
move 1 from 5 to 2
move 2 from 7 to 1
move 1 from 1 to 7
move 1 from 7 to 4
move 13 from 9 to 3
move 8 from 6 to 2
move 10 from 3 to 5
move 14 from 5 to 4
move 7 from 8 to 4
move 1 from 6 to 2
move 6 from 3 to 8
move 4 from 9 to 7
move 2 from 9 to 8
move 1 from 7 to 1
move 3 from 2 to 7
move 1 from 5 to 3
move 7 from 8 to 6
move 5 from 6 to 2
move 8 from 4 to 5
move 3 from 5 to 8
move 3 from 8 to 6
move 5 from 7 to 9
move 5 from 3 to 6
move 1 from 9 to 4
move 17 from 4 to 7
move 1 from 8 to 1
move 12 from 7 to 8
move 3 from 1 to 4
move 2 from 4 to 6
move 8 from 6 to 1
move 4 from 6 to 3
move 1 from 7 to 8
move 5 from 5 to 8
move 4 from 7 to 1
move 3 from 2 to 6
move 2 from 5 to 1
move 6 from 1 to 6
move 4 from 3 to 5
move 4 from 5 to 3
move 1 from 4 to 8
move 3 from 3 to 2
move 17 from 8 to 4
move 6 from 6 to 3
move 14 from 4 to 9
move 1 from 3 to 8
move 1 from 7 to 4
move 3 from 8 to 3
move 5 from 2 to 5
move 6 from 1 to 7
move 2 from 6 to 4
move 4 from 5 to 7
move 1 from 1 to 5
move 1 from 6 to 3
move 10 from 7 to 4
move 1 from 5 to 4
move 1 from 2 to 3
move 15 from 4 to 5
move 3 from 3 to 1
move 6 from 2 to 6
move 1 from 2 to 3
move 2 from 4 to 7
move 2 from 7 to 8
move 1 from 4 to 2
move 2 from 1 to 7
move 1 from 7 to 2
move 12 from 9 to 1
move 4 from 9 to 5
move 4 from 6 to 2
move 1 from 7 to 3
move 6 from 2 to 4
move 1 from 8 to 5
move 2 from 4 to 2
move 11 from 1 to 7
move 3 from 1 to 4
move 17 from 5 to 6
move 15 from 6 to 4
move 1 from 8 to 9
move 10 from 4 to 1
move 1 from 3 to 9
move 2 from 6 to 5
move 1 from 2 to 6
move 4 from 5 to 6
move 4 from 1 to 2
move 6 from 6 to 7
move 2 from 2 to 6
move 9 from 4 to 9
move 6 from 1 to 2
move 3 from 4 to 1
move 10 from 9 to 8
move 4 from 2 to 1
move 1 from 1 to 2
move 5 from 8 to 6
move 1 from 2 to 7
move 1 from 9 to 4
move 2 from 6 to 9
move 13 from 7 to 2
move 5 from 7 to 5
move 2 from 5 to 2
move 1 from 4 to 5
move 4 from 8 to 4
move 17 from 2 to 6
move 3 from 4 to 6
move 2 from 9 to 1
move 7 from 6 to 8
move 1 from 5 to 2
move 1 from 4 to 1
move 2 from 9 to 4
move 1 from 3 to 9
move 4 from 3 to 7
move 2 from 8 to 5
move 3 from 7 to 5
move 10 from 5 to 8
move 2 from 2 to 4
move 6 from 1 to 2
move 4 from 6 to 3
move 8 from 2 to 6
move 1 from 7 to 4
move 5 from 4 to 5
move 7 from 6 to 7
move 5 from 3 to 5
move 5 from 5 to 2
move 4 from 8 to 1
move 6 from 1 to 6
move 3 from 3 to 2
move 22 from 6 to 2
move 1 from 9 to 7
move 8 from 8 to 6
move 1 from 7 to 6
move 2 from 5 to 7
move 4 from 8 to 5
move 7 from 6 to 7
move 2 from 6 to 4
move 14 from 2 to 1
move 7 from 1 to 3
move 12 from 7 to 3
move 1 from 4 to 3
move 2 from 5 to 8
move 2 from 8 to 1
move 1 from 4 to 3
move 6 from 2 to 9
move 6 from 9 to 2
move 2 from 2 to 7
move 6 from 7 to 5
move 13 from 3 to 5
move 5 from 2 to 6
move 5 from 6 to 1
move 2 from 3 to 6
move 1 from 6 to 5
move 1 from 6 to 1
move 3 from 1 to 9
move 6 from 2 to 7
move 1 from 2 to 3
move 24 from 5 to 2
move 7 from 3 to 7
move 13 from 7 to 9
move 4 from 1 to 9
move 4 from 1 to 6
move 1 from 5 to 6
move 16 from 9 to 5
move 1 from 6 to 4
move 1 from 5 to 2
move 5 from 1 to 3
move 11 from 2 to 1
move 4 from 9 to 6
move 1 from 4 to 7
move 2 from 3 to 4
move 6 from 6 to 9
move 1 from 1 to 3
move 2 from 9 to 4
move 1 from 7 to 9
move 4 from 2 to 9
move 8 from 9 to 2
move 3 from 3 to 2
move 1 from 9 to 4
move 5 from 1 to 7
move 1 from 4 to 8
move 2 from 1 to 9
move 1 from 8 to 7
move 6 from 5 to 3
move 1 from 5 to 1
move 5 from 2 to 3
move 4 from 1 to 5
move 4 from 7 to 1
move 8 from 5 to 8`
