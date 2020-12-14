// https://adventofcode.com/2020/day/8

package day8

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Action   string
	Operator int
	Value    int
}

var instructions []Instruction
var ranInstructions = map[int]bool{}
var total int = 0
var scriptPart string

func Exec(scanner *bufio.Scanner, args []string) int {

	scriptPart = args[0]

	// Create array of instructions
	for scanner.Scan() {
		instructions = append(instructions, GetInstructionFromLine(scanner.Text()))
	}

	// Execute instructions until the loop is detected
	ExecInstruction(0, false)
	return total
}

func ExecInstruction(instIndex int, checkingAlt bool) bool {

	instructionToExec := instructions[instIndex]
	// Attemp finding "non-looped fixed path" if applicable
	if scriptPart == "part2" && !checkingAlt {

		// Set original values to reset in case this alt path is not correct
		originalTotal := total
		originalInst := instructionToExec.Action
		originalMap := GetMapCoppy()

		// Temporarely change instructions (nop to jmp or visceversa)
		if originalInst == "nop" {
			instructions[instIndex].Action = "jmp"
		} else if originalInst == "jmp" {
			instructions[instIndex].Action = "nop"
		}

		pathFound := ExecInstruction(instIndex, true)
		if pathFound {
			// Break if fixed path was found
			return true
		} else {
			// Restore originals and continue execution
			total = originalTotal
			instructionToExec.Action = originalInst
			ranInstructions = originalMap
		}
	}

	// Check if instruction was already run
	if ranInstructions[instIndex] {
		// Break the recursion if the instruction was already ran
		return false
	}

	ranInstructions[instIndex] = true

	// Get instruction
	var nextIndex int
	switch instructionToExec.Action {
	case "acc":
		total += instructionToExec.Value * instructionToExec.Operator
		nextIndex = instIndex + 1

	case "jmp":
		nextIndex = instIndex + (instructionToExec.Value * instructionToExec.Operator)
		// Default is nop
	default:
		nextIndex = instIndex + 1
	}

	if nextIndex == len(instructions) {
		// End found
		return true
	}
	return ExecInstruction(nextIndex, checkingAlt)
}

func GetMapCoppy() map[int]bool {
	var mapCopy = map[int]bool{}
	for key, value := range ranInstructions {
		mapCopy[key] = value
	}
	return mapCopy
}

func GetInstructionFromLine(instructionText string) Instruction {
	splitInstruction := strings.Split(instructionText, " ")
	instAction := splitInstruction[0]
	instActionOperator := splitInstruction[1][:1]

	var instOperatorValue int
	if instActionOperator == "-" {
		instOperatorValue = -1
	} else {
		instOperatorValue = 1
	}

	var instValue int
	x, err := strconv.ParseInt(string(splitInstruction[1][1:]), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	instValue = int(x)

	return Instruction{
		Action:   instAction,
		Operator: instOperatorValue,
		Value:    instValue,
	}
}
