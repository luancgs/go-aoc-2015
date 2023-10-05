package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day08"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day08.Part1(input.Reader{})

	fmt.Println("Day 08 - Part 1 | Output: ", output1)

	output2 := day08.Part2(input.Reader{})

	fmt.Println("Day 08 - Part 2 | Output: ", output2)
}
