package day07

import (
	"fmt"
	"strconv"
	"strings"

	input "github.com/luancgs/go-aoc-2015/utils"
)

const day = 7

var variables = make(map[string]uint16)

func getOperation(input string) int {
	if strings.Contains(input, "AND") {
		return 1
	} else if strings.Contains(input, "OR") {
		return 2
	} else if strings.Contains(input, "NOT") {
		return 3
	} else if strings.Contains(input, "RSHIFT") {
		return 4
	} else if strings.Contains(input, "LSHIFT") {
		return 5
	} else {
		return 0
	}
}

func attribution(input string) bool {
	split := strings.Split(input, " -> ")

	var definedOperand bool
	var firstOperand uint16
	returnOperand := split[1]

	if num, err := strconv.Atoi(split[0]); err == nil {
		firstOperand, definedOperand = uint16(num), true
	} else {
		firstOperand, definedOperand = variables[split[0]]
	}

	if !definedOperand {
		return false
	}

	variables[returnOperand] = firstOperand

	return true
}

func bitwiseAnd(input string) bool {
	firstSplit := strings.Split(input, " AND ")
	secondSplit := strings.Split(firstSplit[1], " -> ")

	var definedFirstOperand, definedSecondOperand bool
	var firstOperand, secondOperand uint16
	returnOperand := secondSplit[1]

	if num, err := strconv.Atoi(firstSplit[0]); err == nil {
		firstOperand, definedFirstOperand = uint16(num), true
	} else {
		firstOperand, definedFirstOperand = variables[firstSplit[0]]
	}

	if num, err := strconv.Atoi(secondSplit[0]); err == nil {
		secondOperand, definedSecondOperand = uint16(num), true
	} else {
		secondOperand, definedSecondOperand = variables[secondSplit[0]]
	}

	if !definedFirstOperand || !definedSecondOperand {
		return false
	}

	variables[returnOperand] = firstOperand & secondOperand

	return true
}

func bitwiseOr(input string) bool {
	firstSplit := strings.Split(input, " OR ")
	secondSplit := strings.Split(firstSplit[1], " -> ")

	var definedFirstOperand, definedSecondOperand bool
	var firstOperand, secondOperand uint16
	returnOperand := secondSplit[1]

	if num, err := strconv.Atoi(firstSplit[0]); err == nil {
		firstOperand, definedFirstOperand = uint16(num), true
	} else {
		firstOperand, definedFirstOperand = variables[firstSplit[0]]
	}

	if num, err := strconv.Atoi(secondSplit[0]); err == nil {
		secondOperand, definedSecondOperand = uint16(num), true
	} else {
		secondOperand, definedSecondOperand = variables[secondSplit[0]]
	}

	if !definedFirstOperand || !definedSecondOperand {
		return false
	}

	variables[returnOperand] = firstOperand | secondOperand

	return true
}

func bitwiseNot(input string) bool {
	firstSplit := strings.Split(input, "NOT ")
	secondSplit := strings.Split(firstSplit[1], " -> ")

	var definedOperand bool
	var firstOperand uint16
	returnOperand := secondSplit[1]

	if num, err := strconv.Atoi(secondSplit[0]); err == nil {
		firstOperand, definedOperand = uint16(num), true
	} else {
		firstOperand, definedOperand = variables[secondSplit[0]]
	}

	if !definedOperand {
		return false
	}

	variables[returnOperand] = ^firstOperand

	return true
}

func bitwiseRightShift(input string) bool {
	firstSplit := strings.Split(input, " RSHIFT ")
	secondSplit := strings.Split(firstSplit[1], " -> ")

	var definedFirstOperand, definedSecondOperand bool
	var firstOperand, secondOperand uint16
	returnOperand := secondSplit[1]

	if num, err := strconv.Atoi(firstSplit[0]); err == nil {
		firstOperand, definedFirstOperand = uint16(num), true
	} else {
		firstOperand, definedFirstOperand = variables[firstSplit[0]]
	}

	if num, err := strconv.Atoi(secondSplit[0]); err == nil {
		secondOperand, definedSecondOperand = uint16(num), true
	} else {
		secondOperand, definedSecondOperand = variables[secondSplit[0]]
	}

	if !definedFirstOperand || !definedSecondOperand {
		return false
	}

	variables[returnOperand] = firstOperand >> secondOperand

	return true
}

func bitwiseLeftShift(input string) bool {
	firstSplit := strings.Split(input, " LSHIFT ")
	secondSplit := strings.Split(firstSplit[1], " -> ")

	var definedFirstOperand, definedSecondOperand bool
	var firstOperand, secondOperand uint16
	returnOperand := secondSplit[1]

	if num, err := strconv.Atoi(firstSplit[0]); err == nil {
		firstOperand, definedFirstOperand = uint16(num), true
	} else {
		firstOperand, definedFirstOperand = variables[firstSplit[0]]
	}

	if num, err := strconv.Atoi(secondSplit[0]); err == nil {
		secondOperand, definedSecondOperand = uint16(num), true
	} else {
		secondOperand, definedSecondOperand = variables[secondSplit[0]]
	}

	if !definedFirstOperand || !definedSecondOperand {
		return false
	}

	variables[returnOperand] = firstOperand << secondOperand

	return true
}

func runLines(lines []string) uint16 {
	for len(lines) > 0 {
		for index, line := range lines {
			if line == "" {
				lines = append(lines[:index], lines[index+1:]...)
				break
			}

			operation := getOperation(line)
			var success bool

			switch operation {
			case 0:
				success = attribution(line)
			case 1:
				success = bitwiseAnd(line)
			case 2:
				success = bitwiseOr(line)
			case 3:
				success = bitwiseNot(line)
			case 4:
				success = bitwiseRightShift(line)
			case 5:
				success = bitwiseLeftShift(line)
			}

			if success {
				lines = append(lines[:index], lines[index+1:]...)
				break
			}
		}
	}

	return variables["a"]
}

func formatInput(input string) []string {
	return strings.Split(input, "\n")
}

func Part1(reader input.InputReader) uint16 {
	lines := formatInput(reader.GetInput(day))

	output := runLines(lines)

	return output
}

func Part2(reader input.InputReader) uint16 {
	a := Part1(reader)
	override := fmt.Sprintf("%d -> b", a)

	lines := formatInput(reader.GetInput(day))
	lines = append([]string{override}, lines...)

	variables = make(map[string]uint16)

	output := runLines(lines)

	return output
}
