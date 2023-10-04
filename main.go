package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day04"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day04.Part1(input.Reader{})

	fmt.Println("Day 04 - Part 1 | Output: ", output1)

	output2 := day04.Part2(input.Reader{})

	fmt.Println("Day 04 - Part 2 | Output: ", output2)
}
