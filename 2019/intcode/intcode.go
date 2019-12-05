package intcode

import (
	"fmt"
	"math"
	"strconv"
)

type Opcode int

const (
	add Opcode = 1
	mul Opcode = 2
	sav Opcode = 3
	out Opcode = 4

	jmt Opcode = 5
	jmf Opcode = 6
	lt  Opcode = 7
	eq  Opcode = 8

	halt Opcode = 99
)

type Mode int

const (
	pos Mode = 0
	imd Mode = 1
)

type instruction struct {
	opcode Opcode

	val1Adr  int
	val1Mode Mode

	val2Adr  int
	val2Mode Mode

	destAdr  int
	destMode Mode
}

func (i instruction) instructionPointerOffset() int {
	if i.opcode == add || i.opcode == mul {
		return 4
	} else if i.opcode == sav || i.opcode == out {
		return 2
	}

	return -1
}

func getInstruction(instructionPointer *int, program []int) (instruction, error) {
	ip := *instructionPointer

	// Parse Opcode and Mode
	defaultModes := []Mode{pos, pos, pos}
	code := strconv.Itoa(program[ip])
	var op string
	if len(code) >= 2 {
		opCutoff := len(code) - 2
		op = string(code[opCutoff:])
		// remove opcode from end
		code = code[:opCutoff]
		for i, j := len(code)-1, 0; i >= 0; i, j = i-1, j+1 {
			modeInt, err := strconv.Atoi(string(code[i]))
			if err != nil {
				return instruction{}, err
			}

			defaultModes[j] = Mode(modeInt)
		}
	} else {
		op = string(code[0])
	}

	opInt, err := strconv.Atoi(op)
	if err != nil {
		return instruction{}, err
	}

	opcode := Opcode(opInt)
	if opcode != add &&
		opcode != mul &&
		opcode != sav &&
		opcode != out &&
		opcode != jmt &&
		opcode != jmf &&
		opcode != lt &&
		opcode != eq &&
		opcode != halt {
		return instruction{}, fmt.Errorf("unknown opcode: %d", opcode)
	}

	if opcode == halt {
		return instruction{opcode: opcode}, nil
	}

	inst := instruction{
		opcode: opcode,
	}
	if opcode == add || opcode == mul || opcode == lt || opcode == eq {
		inst.val1Adr = program[ip+1]
		inst.val1Mode = defaultModes[0]
		inst.val2Adr = program[ip+2]
		inst.val2Mode = defaultModes[1]
		inst.destAdr = program[ip+3]
		inst.destMode = defaultModes[2]
		*instructionPointer = ip + 4
	} else if opcode == jmt || opcode == jmf {
		inst.val1Adr = program[ip+1]
		inst.val1Mode = defaultModes[0]
		inst.destAdr = program[ip+2]
		inst.destMode = defaultModes[1]
		*instructionPointer = ip + 3
	} else if opcode == sav || opcode == out {
		inst.val1Adr = program[ip+1]
		inst.val1Mode = defaultModes[0]
		*instructionPointer = ip + 2
	}

	return inst, nil
}

func addValues(inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var val2 int
	switch inst.val2Mode {
	case pos:
		val2 = program[inst.val2Adr]
	case imd:
		val2 = inst.val2Adr
	}

	sum := val1 + val2
	// Can a destination be immediate mode?
	if inst.destMode == imd {
		fmt.Printf("immediate add dest: %+v\n", inst)
	}
	program[inst.destAdr] = sum
}

func mulValues(inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var val2 int
	switch inst.val2Mode {
	case pos:
		val2 = program[inst.val2Adr]
	case imd:
		val2 = inst.val2Adr
	}

	product := val1 * val2
	// Can a destination be immediate mode?
	if inst.destMode == imd {
		fmt.Printf("immediate mul dest: %+v\n", inst)
	}
	program[inst.destAdr] = product
}

// Can an input be immediate mode?
func saveValues(inst instruction, in int, program []int) {
	if inst.val1Mode == imd {
		fmt.Printf("immediate save: %+v\n", inst)
	}
	program[inst.val1Adr] = in
}

// Immediate output is verified as possible
func outputValues(inst instruction, program []int) int {
	switch inst.val1Mode {
	case pos:
		return program[inst.val1Adr]
	case imd:
		return inst.val1Adr
	}

	return -1 * math.MaxInt64
}

func jumpIfTrue(instructionPointer *int, inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var dest int
	switch inst.destMode {
	case pos:
		dest = program[inst.destAdr]
	case imd:
		dest = inst.destAdr
	}

	if val1 != 0 {
		*instructionPointer = dest
	}
}

func jumpIfFalse(instructionPointer *int, inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var dest int
	switch inst.destMode {
	case pos:
		dest = program[inst.destAdr]
	case imd:
		dest = inst.destAdr
	}

	if val1 == 0 {
		*instructionPointer = dest
	}
}

func lessThan(inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var val2 int
	switch inst.val2Mode {
	case pos:
		val2 = program[inst.val2Adr]
	case imd:
		val2 = inst.val2Adr
	}

	destVal := 0
	if val1 < val2 {
		destVal = 1
	}

	program[inst.destAdr] = destVal
}

func equals(inst instruction, program []int) {
	var val1 int
	switch inst.val1Mode {
	case pos:
		val1 = program[inst.val1Adr]
	case imd:
		val1 = inst.val1Adr
	}

	var val2 int
	switch inst.val2Mode {
	case pos:
		val2 = program[inst.val2Adr]
	case imd:
		val2 = inst.val2Adr
	}

	destVal := 0
	if val1 == val2 {
		destVal = 1
	}

	program[inst.destAdr] = destVal
}

func RunProgram(memory []int, inputs []int) error {
	instructionPointer := 0
	inst, err := getInstruction(&instructionPointer, memory)
	if err != nil {
		return err
	}

	inputIdx := 0
	for inst.opcode != halt {
		// fmt.Printf("%+v\n", inst)
		// fmt.Println(memory)

		switch inst.opcode {
		case add:
			addValues(inst, memory)
		case mul:
			mulValues(inst, memory)
		case sav:
			saveValues(inst, inputs[inputIdx], memory)
		case out:
			output := outputValues(inst, memory)
			fmt.Println(output)
		case jmt:
			jumpIfTrue(&instructionPointer, inst, memory)
		case jmf:
			jumpIfFalse(&instructionPointer, inst, memory)
		case lt:
			lessThan(inst, memory)
		case eq:
			equals(inst, memory)
		}

		// instructionPointer += inst.instructionPointerOffset()
		inst, err = getInstruction(&instructionPointer, memory)
		if err != nil {
			return err
		}
	}

	return nil
}
