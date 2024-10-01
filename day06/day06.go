package day06

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

const day = 6

func formatInput(input string) ([]int, [][]int) {
	lines := strings.Split(input, "\n")

	var actions []int
	var coordinates [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}

		var action int
		var actionString string

		if strings.HasPrefix(line, "turn off") {
			action = 0
			actionString = "turn off"
		} else if strings.HasPrefix(line, "turn on") {
			action = 1
			actionString = "turn on"
		} else {
			action = 2
			actionString = "toggle"
		}

		actions = append(actions, action)
		coordStrings := strings.ReplaceAll(strings.Split(line, fmt.Sprintf("%s ", actionString))[1], " through ", ",")

		var coords []int

		for _, coordinate := range strings.Split(coordStrings, ",") {
			num, err := strconv.Atoi(coordinate)
			if err != nil {
				log.Fatal(err)
			}

			coords = append(coords, num)
		}

		coordinates = append(coordinates, coords)
	}

	return actions, coordinates
}

func Part1(reader input.InputReader) string {
	actions, coordinates := formatInput(reader.GetInput(day))

	var grid [1000][1000]bool

	for index, coordinate := range coordinates {
		startX := coordinate[0]
		startY := coordinate[1]
		endX := coordinate[2]
		endY := coordinate[3]

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				switch actions[index] {
				case 0:
					grid[i][j] = false
				case 1:
					grid[i][j] = true
				case 2:
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}

	lit := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				lit++
			}
		}
	}

	return fmt.Sprintf("%d", lit)
}

func Part2(reader input.InputReader) string {
	actions, coordinates := formatInput(reader.GetInput(day))

	var grid [1000][1000]int

	for index, coordinate := range coordinates {
		startX := coordinate[0]
		startY := coordinate[1]
		endX := coordinate[2]
		endY := coordinate[3]

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				switch actions[index] {
				case 0:
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				case 1:
					grid[i][j]++
				case 2:
					grid[i][j] += 2
				}
			}
		}
	}

	brightness := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			brightness += grid[x][y]
		}
	}

	return fmt.Sprintf("%d", brightness)
}
