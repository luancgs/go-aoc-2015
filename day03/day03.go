package day03

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

const day = 3

func formatInput(input string) []rune {
	return []rune(input)
}

func Part1(reader input.InputReader) int {
	directions := formatInput(reader.GetInput(day))

	set := make(map[string]bool)
	set["0,0"] = true

	x := 0
	y := 0

	for _, direction := range directions {
		switch direction {
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		case '<':
			x--
		}

		set[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return len(set)
}

func Part2(reader input.InputReader) int {
	directions := formatInput(reader.GetInput(day))

	set := make(map[string]bool)
	set["0,0"] = true

	santaX := 0
	santaY := 0
	robotX := 0
	robotY := 0

	isSanta := true

	for _, direction := range directions {
		if isSanta {
			switch direction {
			case '^':
				santaY++
			case '>':
				santaX++
			case 'v':
				santaY--
			case '<':
				santaX--
			}
			set[fmt.Sprintf("%d,%d", santaX, santaY)] = true
		} else {
			switch direction {
			case '^':
				robotY++
			case '>':
				robotX++
			case 'v':
				robotY--
			case '<':
				robotX--
			}
			set[fmt.Sprintf("%d,%d", robotX, robotY)] = true
		}
		isSanta = !isSanta
	}

	return len(set)
}
