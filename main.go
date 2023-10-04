package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day03"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day03.Part1(input.Reader{})

	fmt.Println("Day 03 - Part 1 | Output: ", output1)

	output2 := day03.Part2(input.Reader{})

	fmt.Println("Day 03 - Part 2 | Output: ", output2)
}
