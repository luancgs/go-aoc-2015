package day04

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

const day = 4

func hash(key string, prefix string) int {
	hash := ""
	i := 1

	for !strings.HasPrefix(hash, prefix) {
		data := fmt.Sprintf("%s%d", key, i)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(data)))
		i++
	}

	return i - 1
}

func Part1(reader input.InputReader) string {
	secretKey := reader.GetInput(day)
	output := hash(secretKey, "00000")
	return fmt.Sprintf("%d", output)
}

func Part2(reader input.InputReader) string {
	secretKey := reader.GetInput(day)
	output := hash(secretKey, "000000")
	return fmt.Sprintf("%d", output)
}
