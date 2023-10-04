package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day05"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day05.Part1(input.Reader{})

	fmt.Println("Day 05 - Part 1 | Output: ", output1)

	output2 := day05.Part2(input.Reader{})

	fmt.Println("Day 05 - Part 2 | Output: ", output2)
}
