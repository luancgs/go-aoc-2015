package day02

import (
	"log"
	"sort"
	"strconv"
	"strings"

	input "github.com/luancgs/go-aoc-2015/utils"
)

const day = 2

func formatInput(input string) [][]int {
	var output [][]int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		stringNumbers := strings.Split(line, "x")

		if len(stringNumbers) != 3 {
			continue
		}

		var dimensions []int
		for _, stringNumber := range stringNumbers {

			if stringNumber == "" {
				continue
			}

			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Fatal(err)
			}

			dimensions = append(dimensions, number)
		}

		output = append(output, dimensions)
	}

	return output
}

func Part1(reader input.InputReader) int {
	allPresents := formatInput(reader.GetInput(day))

	total := 0

	for _, presentDimensions := range allPresents {
		l := presentDimensions[0]
		w := presentDimensions[1]
		h := presentDimensions[2]

		sides := []int{l * w, w * h, h * l}
		smallestSide := sides[0]

		for _, side := range sides {
			if side < smallestSide {
				smallestSide = side
			}

			total += 2 * side
		}

		total += smallestSide
	}

	return total
}

func Part2(reader input.InputReader) int {
	allPresents := formatInput(reader.GetInput(day))

	total := 0

	for _, presentDimensions := range allPresents {

		sort.Ints(presentDimensions)

		total += (2 * presentDimensions[0]) + (2 * presentDimensions[1]) + (presentDimensions[0] * presentDimensions[1] * presentDimensions[2])
	}

	return total
}
