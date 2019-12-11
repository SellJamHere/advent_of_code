package intcode

import (
	"fmt"
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

	arb Opcode = 9

	halt Opcode = 99
)

type Mode int

const (
	pos Mode = 0
	imd Mode = 1
	rel Mode = 2
)

type instruction struct {
	opcode Opcode

	val1     int
	val1Mode Mode

	val2     int
	val2Mode Mode

	val3     int
	val3Mode Mode
}

func (i instruction) instructionPointerOffset() int {
	if i.opcode == add || i.opcode == mul {
		return 4
	} else if i.opcode == sav || i.opcode == out {
		return 2
	}

	return -1
}

type Memory map[int]int

func NewMemory(program []int) Memory {
	mem := Memory{}
	for i, val := range program {
		mem[i] = val
	}

	return mem
}

func (m Memory) Copy() Memory {
	mem := Memory{}
	for i, val := range m {
		mem[i] = val
	}

	return mem
}

func getInstruction(instructionPointer *int, program map[int]int) (instruction, error) {
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
		opcode != arb &&
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
		inst.val1 = program[ip+1]
		inst.val1Mode = defaultModes[0]
		inst.val2 = program[ip+2]
		inst.val2Mode = defaultModes[1]
		inst.val3 = program[ip+3]
		inst.val3Mode = defaultModes[2]
		*instructionPointer = ip + 4
	} else if opcode == jmt || opcode == jmf {
		inst.val1 = program[ip+1]
		inst.val1Mode = defaultModes[0]
		inst.val2 = program[ip+2]
		inst.val2Mode = defaultModes[1]
		*instructionPointer = ip + 3
	} else if opcode == sav || opcode == out || opcode == arb {
		inst.val1 = program[ip+1]
		inst.val1Mode = defaultModes[0]
		*instructionPointer = ip + 2
	}

	return inst, nil
}

func addValues(inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	sum := val1 + val2

	setValue(sum, inst.val3, inst.val3Mode, memory, relativeBase)
}

func mulValues(inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	product := val1 * val2

	setValue(product, inst.val3, inst.val3Mode, memory, relativeBase)
}

func saveValues(inst instruction, in int, memory map[int]int, relativeBase int) {
	setValue(in, inst.val1, inst.val1Mode, memory, relativeBase)
}

// Immediate output is verified as possible
func outputValues(inst instruction, memory map[int]int, relativeBase int) int {
	return getValue(inst.val1, inst.val1Mode, memory, relativeBase)
}

func jumpIfTrue(instructionPointer *int, inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	if val1 != 0 {
		*instructionPointer = val2
	}
}

func jumpIfFalse(instructionPointer *int, inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	if val1 == 0 {
		*instructionPointer = val2
	}
}

func lessThan(inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	destVal := 0
	if val1 < val2 {
		destVal = 1
	}

	setValue(destVal, inst.val3, inst.val3Mode, memory, relativeBase)
}

func equals(inst instruction, memory map[int]int, relativeBase int) {
	val1 := getValue(inst.val1, inst.val1Mode, memory, relativeBase)
	val2 := getValue(inst.val2, inst.val2Mode, memory, relativeBase)

	destVal := 0
	if val1 == val2 {
		destVal = 1
	}

	setValue(destVal, inst.val3, inst.val3Mode, memory, relativeBase)
}

func adjustRelativeBase(inst instruction, memory map[int]int, relativeBase *int) {
	switch inst.val1Mode {
	case pos:
		*relativeBase += memory[inst.val1]
	case imd:
		*relativeBase += inst.val1
	case rel:
		*relativeBase += memory[*relativeBase+inst.val1]
	}
}

func getValue(val int, mode Mode, memory map[int]int, relativeBase int) int {
	value := 0
	switch mode {
	case pos:
		value = memory[val]
	case imd:
		value = val
	case rel:
		value = memory[val+relativeBase]
	}

	return value
}

func setValue(input, val int, mode Mode, memory map[int]int, relativeBase int) {
	switch mode {
	case pos:
		memory[val] = input
	case imd:
		fmt.Println("immediate mode used for setValue")
	case rel:
		memory[val+relativeBase] = input
	}
}

func RunProgram(tag int, memory map[int]int, input <-chan int, output chan<- int) error {
	instructionPointer := 0
	inst, err := getInstruction(&instructionPointer, memory)
	if err != nil {
		return err
	}

	relativeBase := 0

	for inst.opcode != halt {
		switch inst.opcode {
		case add:
			addValues(inst, memory, relativeBase)
		case mul:
			mulValues(inst, memory, relativeBase)
		case sav:
			in := <-input
			saveValues(inst, in, memory, relativeBase)
		case out:
			out := outputValues(inst, memory, relativeBase)
			output <- out
		case jmt:
			jumpIfTrue(&instructionPointer, inst, memory, relativeBase)
		case jmf:
			jumpIfFalse(&instructionPointer, inst, memory, relativeBase)
		case lt:
			lessThan(inst, memory, relativeBase)
		case eq:
			equals(inst, memory, relativeBase)
		case arb:
			adjustRelativeBase(inst, memory, &relativeBase)
		}

		inst, err = getInstruction(&instructionPointer, memory)
		if err != nil {
			return err
		}
	}

	return nil
}
