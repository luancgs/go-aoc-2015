package day01

import input "github.com/luancgs/go-aoc-2015/utils"

const day = 1

func formatInput(input string) []rune {
	return []rune(input)
}

func Part1(reader input.InputReader) int {
	content := reader.GetInput(day)

	var floor int = 0
	for _, char := range formatInput(content) {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}

	return floor
}

func Part2(reader input.InputReader) int {
	content := reader.GetInput(day)

	var floor int = 0
	for index, char := range formatInput(content) {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}

		if floor == -1 {
			return index + 1
		}
	}

	return 0
}
