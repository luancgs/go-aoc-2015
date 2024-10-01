package day05

import (
	"fmt"
	"strings"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

const day = 5

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func formatInput(input string) []string {
	allStrings := strings.Split(input, "\n")
	return allStrings
}

func containsThreeVowels(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	total := 0

	for _, vowel := range vowels {
		total += strings.Count(s, vowel)

		if total >= 3 {
			return true
		}
	}

	return false
}

func containsDoubleLetters(s string) bool {
	for _, letter := range alphabet {
		if strings.Contains(s, fmt.Sprintf("%s%s", letter, letter)) {
			return true
		}
	}

	return false
}

func containsProhibitedStrings(s string) bool {
	prohibitedStrings := []string{"ab", "cd", "pq", "xy"}

	for _, prohibitedString := range prohibitedStrings {
		if strings.Contains(s, prohibitedString) {
			return true
		}
	}

	return false
}

func containsPairOfTwo(s string) bool {
	for _, letter1 := range alphabet {
		for _, letter2 := range alphabet {
			pairs := strings.Count(s, fmt.Sprintf("%s%s", letter1, letter2))

			if pairs >= 2 {
				return true
			}
		}
	}

	return false
}

func containsDoubleWithSpace(s string) bool {
	for _, letter := range alphabet {
		for _, space := range alphabet {
			if strings.Contains(s, fmt.Sprintf("%s%s%s", letter, space, letter)) {
				return true
			}
		}
	}

	return false
}

func Part1(reader input.InputReader) string {
	allStrings := formatInput(reader.GetInput(day))
	total := 0

	for _, s := range allStrings {
		if s == "" {
			continue
		}

		if containsThreeVowels(s) && containsDoubleLetters(s) && !containsProhibitedStrings(s) {
			total++
		}
	}

	return fmt.Sprintf("%d", total)
}

func Part2(reader input.InputReader) string {
	allStrings := formatInput(reader.GetInput(day))
	total := 0

	for _, s := range allStrings {
		if s == "" {
			continue
		}

		if containsPairOfTwo(s) && containsDoubleWithSpace(s) {
			total++
		}
	}

	return fmt.Sprintf("%d", total)
}
