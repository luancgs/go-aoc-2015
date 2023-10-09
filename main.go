package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day09"
	"github.com/luancgs/go-aoc-2015/utils/input"
)

func main() {

	output1 := day09.Part1(input.Reader{})

	fmt.Println("Day 09 - Part 1 | Output: ", output1)

	output2 := day09.Part2(input.Reader{})

	fmt.Println("Day 09 - Part 2 | Output: ", output2)
}
