package vm

/*
Assembly File will look like this:

PUSH 1
PUSH 2
ADD
DUMP

-----------------------------------------------------

<instruction> <operand>\n

// PUSH is the only instruction that takes an operand
// all other instructions don't require an operand and can be omitted
*/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(path string) (Program, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("No path provided")
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file := string(raw)

	program := make(Program, 0)
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		inst, err := ParseLine(line)
		if err != nil {
			return nil, err
		}
		program = append(program, inst)
	}

	return program, nil
}

func ParseLine(line string) (Instruction, error) {
	if len(line) == 0 {
		return Instruction{}, fmt.Errorf("No line provided")
	}

	parts := strings.Split(line, " ")
	if len(parts) < 1 {
		return Instruction{}, fmt.Errorf("No instruction provided")
	}

	// if the instruction is a PUSH we need to parse the operand and the data
	if len(parts) > 1 {
		op, err := ParseOperand(parts[0])
		if err != nil {
			return Instruction{}, err
		}
		data, err := strconv.Atoi(parts[1])
		if err != nil {
			return Instruction{}, err
		}
		return Instruction{op, data}, nil
	}
	// otherwise, we just need to parse the instruction
	op, err := ParseOperand(parts[0])
	if err != nil {
		return Instruction{}, err
	}

	return Instruction{op, 0}, nil
}

func ParseOperand(op string) (OP, error) {
	switch op {
	case "PUSH":
		return PUSH, nil
	case "POP":
		return POP, nil
	case "ADD":
		return ADD, nil
	case "SUB":
		return SUB, nil
	case "MUL":
		return MUL, nil
	case "DIV":
		return DIV, nil
	case "MOD":
		return MOD, nil
	case "DUMP":
		return DUMP, nil
	default:
		return OP(0), fmt.Errorf("Invalid instruction: %s", op)
	}
}
