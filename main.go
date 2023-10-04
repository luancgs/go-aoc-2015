package main

import (
	"fmt"

	"github.com/luancgs/go-aoc-2015/day02"
	input "github.com/luancgs/go-aoc-2015/utils"
)

func main() {

	output1 := day02.Part1(input.Reader{})

	fmt.Println("Day 02 - Part 1 | Output: ", output1)

	output2 := day02.Part2(input.Reader{})

	fmt.Println("Day 02 - Part 2 | Output: ", output2)

}
