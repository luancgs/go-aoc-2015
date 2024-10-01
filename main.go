package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luancgs/go-aoc-2015/day01"
	"github.com/luancgs/go-aoc-2015/day02"
	"github.com/luancgs/go-aoc-2015/day03"
	"github.com/luancgs/go-aoc-2015/day04"
	"github.com/luancgs/go-aoc-2015/day05"
	"github.com/luancgs/go-aoc-2015/day06"
	"github.com/luancgs/go-aoc-2015/day07"
	"github.com/luancgs/go-aoc-2015/day08"
	"github.com/luancgs/go-aoc-2015/utils/input"
)

type DayFunction func(input.InputReader) string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day provided. Please provide a valid integer.")
		return
	}

	days := map[int]struct {
		part1 DayFunction
		part2 DayFunction
	}{
		1: {day01.Part1, day01.Part2},
		2: {day02.Part1, day02.Part2},
		3: {day03.Part1, day03.Part2},
		4: {day04.Part1, day04.Part2},
		5: {day05.Part1, day05.Part2},
		6: {day06.Part1, day06.Part2},
		7: {day07.Part1, day07.Part2},
		8: {day08.Part1, day08.Part2},
	}

	if functions, exists := days[day]; exists {
		reader := input.Reader{}
		output1 := functions.part1(reader)
		fmt.Println("Day %02d - Part 1 | Output: %s\n", day, output1)
		output2 := functions.part2(reader)
		fmt.Println("Day %02d - Part 2 | Output: %s\n", day, output2)
	} else {
		fmt.Println("Solution not implemented yet!")
	}
}
