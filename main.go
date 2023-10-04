package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day01"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day01.Part1(input.Reader{})

	fmt.Println("Day 01 - Part 1 | Output: ", output1)

	output2 := day01.Part2(input.Reader{})

	fmt.Println("Day 01 - Part 2 | Output: ", output2)

}
