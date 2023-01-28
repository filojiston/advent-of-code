// https://adventofcode.com/2015/day/23
// solution of advent of code 2015, day23

package day23

import (
	"strconv"
	"strings"

	"github.com/filojiston/advent-of-code/2015/util"
)

type register struct {
	name  string
	value uint32
}

type instruction struct {
	name     string
	register *register
	offset   int
}

func openTuringLock(startA, startB uint32) (uint32, uint32) {
	input := util.ReadInputFile("input.txt")
	registerA := register{name: "a", value: startA}
	registerB := register{name: "b", value: startB}
	instructions := loadInstructions(input, &registerA, &registerB)
	applyInstructions(instructions)
	return registerA.value, registerB.value
}

func loadInstructions(lines []string, registerA, registerB *register) []instruction {
	var instructions []instruction
	for _, line := range lines {
		var name, registerName string
		var register *register
		var offset int
		input := strings.Split(line, " ")
		name = input[0]
		if name == "jmp" {
			offset, _ = strconv.Atoi(input[1])
		} else if name == "jio" || name == "jie" {
			registerName = input[1][:len(input[1])-1]
			offset, _ = strconv.Atoi(input[2])
		} else {
			registerName = input[1]
		}

		if registerName == "a" {
			register = registerA
		} else if registerName == "b" {
			register = registerB
		}

		instructions = append(instructions, instruction{name: name, register: register, offset: offset})
	}
	return instructions
}

func applyInstructions(instructions []instruction) {
	var currentPtr int
	for currentPtr >= 0 && currentPtr < len(instructions) {
		currentInst := instructions[currentPtr]
		switch currentInst.name {
		case "hlf":
			currentInst.register.value /= 2
			currentPtr++
		case "tpl":
			currentInst.register.value *= 3
			currentPtr++
		case "inc":
			currentInst.register.value++
			currentPtr++
		case "jmp":
			currentPtr += currentInst.offset
		case "jie":
			if currentInst.register.value%2 == 0 {
				currentPtr += currentInst.offset
			} else {
				currentPtr++
			}
		case "jio":
			if currentInst.register.value == 1 {
				currentPtr += currentInst.offset
			} else {
				currentPtr++
			}
		}
	}
}
