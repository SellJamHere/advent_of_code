package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SellJamHere/advent_of_code/2019/intcode"
)

func main() {
	// Part 1
	program, err := readInput(input)
	if err != nil {
		panic(err)
	}

	memory := make([]int, len(program))

	perms := generatePermutations([]int{0, 1, 2, 3, 4})

	highestOutput := 0
	for _, perm := range perms {
		nextInput := 0

		for i, phase := range perm {
			copy(memory, program)
			input := make(chan int, 2)
			output := make(chan int)

			input <- phase
			input <- nextInput
			go func() {
				err := intcode.RunProgram(i, memory, input, output)
				if err != nil {
					panic(err)
				}
			}()

			nextInput = <-output
		}

		if nextInput > highestOutput {
			highestOutput = nextInput
		}
	}

	fmt.Println("highest output:", highestOutput)

	// Part 2
	highestOutput = 0
	perms = generatePermutations([]int{5, 6, 7, 8, 9})
	// perms = [][]int{[]int{9, 8, 7, 6, 5}}
	for _, perm := range perms {

		programs := make([][]int, len(perm), len(perm))
		for i := range perm {
			programs[i] = make([]int, len(program))
			copy(programs[i], program)
		}

		inputs := make([]chan int, len(perm), len(perm))
		for i := range perm {
			inputs[i] = make(chan int, 1)
		}

		finalOutput := make(chan int, 1)

		// initialize amplifiers
		for i, phase := range perm {
			memory := programs[i]

			input := inputs[i]
			var output chan int
			if i < len(perm)-1 {
				output = inputs[i+1]
			} else {
				output = finalOutput
			}

			go func(i int, memory []int, input chan int, output chan int) {
				err := intcode.RunProgram(i, memory, input, output)
				if err != nil {
					panic(err)
				}

				close(output)
			}(i, memory, input, output)

			input <- phase
		}

		// Start Amplifiers
		inputs[0] <- 0

		recentOutput := -1
		for out := range finalOutput {
			recentOutput = out
			inputs[0] <- out
		}

		if recentOutput > highestOutput {
			highestOutput = recentOutput
		}
	}

	fmt.Println("highestOutput:", highestOutput)
}

func readInput(programString string) ([]int, error) {
	program := []int{}
	values := strings.Split(programString, ",")
	for _, valueStr := range values {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, err
		}

		program = append(program, value)
	}

	return program, nil
}

func generatePermutations(input []int) [][]int {
	res := [][]int{}

	var helper func([]int, int)
	helper = func(input []int, n int) {
		if n == 1 {
			tmp := make([]int, len(input))
			copy(tmp, input)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(input, n-1)
				if n%2 == 1 {
					tmp := input[i]
					input[i] = input[n-1]
					input[n-1] = tmp
				} else {
					tmp := input[0]
					input[0] = input[n-1]
					input[n-1] = tmp
				}
			}
		}
	}

	helper(input, len(input))

	return res
}

// const input = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`

const input = `3,8,1001,8,10,8,105,1,0,0,21,42,67,84,109,122,203,284,365,446,99999,3,9,1002,9,3,9,1001,9,5,9,102,4,9,9,1001,9,3,9,4,9,99,3,9,1001,9,5,9,1002,9,3,9,1001,9,4,9,102,3,9,9,101,3,9,9,4,9,99,3,9,101,5,9,9,1002,9,3,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,102,3,9,9,101,3,9,9,102,2,9,9,4,9,99,3,9,101,2,9,9,1002,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,99`
