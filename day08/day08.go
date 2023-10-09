package day08

import (
	"regexp"
	"strings"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

const day = 8

func formatInput(input string) []string {
	return strings.Split(input, "\n")
}

func Part1(reader input.InputReader) int {
	lines := formatInput(reader.GetInput(day))

	codeTotal := 0
	literalTotal := 0

	for _, line := range lines {

		if line == "" {
			continue
		}
		codeTotal += len(line)

		line := strings.TrimPrefix(line, "\"")
		line = strings.TrimSuffix(line, "\"")
		line = strings.ReplaceAll(line, "\\\\", "-")
		line = strings.ReplaceAll(line, "\\\"", "-")
		m := regexp.MustCompile(`(\\x[0-9a-fA-F]{2})`)
		line = m.ReplaceAllString(line, "-")

		literalTotal += len(line)
	}

	return codeTotal - literalTotal
}

func Part2(reader input.InputReader) int {
	lines := formatInput(reader.GetInput(day))

	codeTotal := 0
	literalTotal := 0

	for _, line := range lines {

		if line == "" {
			continue
		}

		codeTotal += len(line)

		line = strings.ReplaceAll(line, "\\", "\\\\")
		line := strings.ReplaceAll(line, "\"", "\\\"")

		literalTotal += len(line) + 2
	}

	return literalTotal - codeTotal
}
