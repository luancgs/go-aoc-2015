package input

import (
	"fmt"
	"os"
)

const year = 2015

type InputReader interface {
	GetInput(day int, isTest ...bool) string
}

type Reader struct{}

// Receives the day and returns the input as a string
func (Reader) GetInput(day int, isTest ...bool) string {
	var inputPath string

	if len(isTest) > 0 && isTest[0] {
		inputPath = fmt.Sprintf("../inputs/%d/sample/day%02d.txt", year, day)
	} else {
		inputPath = fmt.Sprintf("./inputs/%d/day%02d.txt", year, day)
	}

	content, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return string(content)
}
