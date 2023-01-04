// https://adventofcode.com/2015/day/7
// solution of advent of code 2015, day7

// probably there's much better ways to solve this particular problem, but this is what i come with
// for example i've seen people solve this problem with dfs (one of the much better ways)

package day7

import (
	"strconv"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type signal struct {
	value uint16
	valid bool
}

type instruction struct {
	operand1 *wire
	operator operator
	operand2 *wire
	result   *wire
}

var instructions []instruction

func emulateCircuit(filename string) {
	lines := util.ReadInputFile(filename)
	for _, line := range lines {
		instructions = append(instructions, parseInstructions(line))
	}

	for !isCircuitCompleted() {
		applyInstructions()
	}
}

func isCircuitCompleted() bool {
	for _, wire := range wires {
		if !wire.signal.valid {
			return false
		}
	}
	return true
}

func applyInstructions() {
	for _, instruction := range instructions {
		if instruction.operator == nil {
			instruction.result.signal = instruction.operand1.signal
		} else if instruction.operand2 == nil {
			instruction.operator(instruction.operand1, nil, instruction.result)
		} else {
			instruction.operator(instruction.operand1, instruction.operand2, instruction.result)
		}
	}
}

func parseInstructions(line string) instruction {
	splitLine := strings.Split(line, " ")
	var operand1, operand2, result *wire
	var operator operator

	if len(splitLine) == 3 {
		operand1 = parseOperand(splitLine[0])
		result = getWire(splitLine[2])
	} else if len(splitLine) == 4 {
		operator = not
		operand1 = parseOperand(splitLine[1])
		result = getWire(splitLine[3])
	} else {
		operand1 = parseOperand(splitLine[0])
		operator = getOperator(splitLine[1])
		operand2 = parseOperand(splitLine[2])
		result = getWire(splitLine[4])
	}

	return instruction{operand1: operand1, operator: operator, operand2: operand2, result: result}
}

func parseOperand(operandString string) *wire {
	val, err := strconv.ParseUint(operandString, 0, 16)
	if err != nil {
		return getWire(operandString)
	} else {
		valUint16 := uint16(val)
		return createWire("", signal{value: valUint16, valid: true})
	}
}

func resetInstructions() {
	instructions = nil
}
