package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	program, err := readInput(input)
	if err != nil {
		panic(err)
	}

	memory := make([]int, len(program))
	copy(memory, program)

	// adjust program
	memory[1] = 12
	memory[2] = 2

	runProgram(memory)

	fmt.Printf("postition zero value: %d\n", memory[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(memory, program)

			memory[1] = noun
			memory[2] = verb

			runProgram(memory)
			if memory[0] == 19690720 {
				fmt.Printf("noun: %d, verb: %d\n100 * %d + %d = %d\n", noun, verb, noun, verb, 100*noun+verb)
				return
			}
		}
	}
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

type Opcode int

var add Opcode = 1
var mul Opcode = 2
var halt Opcode = 99

type instruction struct {
	opcode  Opcode
	val1Adr int
	val2Adr int
	destAdr int
}

func getInstruction(index int, program []int) (instruction, error) {
	opcode := program[index]
	if Opcode(opcode) != add && Opcode(opcode) != mul && Opcode(opcode) != halt {
		return instruction{}, fmt.Errorf("unknown opcode: %d", opcode)
	}

	if Opcode(opcode) == halt {
		return instruction{opcode: Opcode(opcode)}, nil
	}

	val1 := program[index+1]
	val2 := program[index+2]
	dest := program[index+3]

	inst := instruction{Opcode(opcode), val1, val2, dest}

	return inst, nil
}

func addValues(inst instruction, program []int) {
	sum := program[inst.val1Adr] + program[inst.val2Adr]
	program[inst.destAdr] = sum
}

func mulValues(inst instruction, program []int) {
	product := program[inst.val1Adr] * program[inst.val2Adr]
	program[inst.destAdr] = product
}

func runProgram(memory []int) error {
	instructionPointer := 0
	inst, err := getInstruction(instructionPointer, memory)
	if err != nil {
		return err
	}

	for inst.opcode != halt {
		switch inst.opcode {
		case add:
			addValues(inst, memory)
		case mul:
			mulValues(inst, memory)
		}

		instructionPointer += 4
		inst, err = getInstruction(instructionPointer, memory)
		if err != nil {
			return err
		}
	}

	return nil
}

const input = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,13,19,1,10,19,23,2,9,23,27,1,6,27,31,1,10,31,35,1,35,10,39,1,9,39,43,1,6,43,47,1,10,47,51,1,6,51,55,2,13,55,59,1,6,59,63,1,10,63,67,2,67,9,71,1,71,5,75,1,13,75,79,2,79,13,83,1,83,9,87,2,10,87,91,2,91,6,95,2,13,95,99,1,10,99,103,2,9,103,107,1,107,5,111,2,9,111,115,1,5,115,119,1,9,119,123,2,123,6,127,1,5,127,131,1,10,131,135,1,135,6,139,1,139,5,143,1,143,9,147,1,5,147,151,1,151,13,155,1,5,155,159,1,2,159,163,1,163,6,0,99,2,0,14,0`
