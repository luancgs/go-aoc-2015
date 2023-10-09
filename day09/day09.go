package day09

import (
	"strings"

	"github.com/luancgs/go-aoc-2015/utils/input"
)

type Node struct {
	Local     string
	Neighbors []Node
	Visited   bool
}

type Edge struct {
	From     Node
	To       Node
	Distance int
}

type Graph struct {
	Nodes []Node
	Edges []Edge
}

const day = 9

type Nodes map[string]Node

var locals Nodes = make(Nodes)

func formatInput(input string) {

	directions := strings.Split(input, "\n")

	for _, direction := range directions {
		if direction == "" {
			continue
		}

		firstSplit := strings.Split(direction, " = ")
		cities := strings.Split(firstSplit[0], " to ")

		from := cities[0]
		to := cities[1]

		// distance, err := strconv.Atoi(firstSplit[1])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if _, exists := locals[from]; !exists {
			locals[from] = Node{from, []Node{}, false}
		}

		if _, exists := locals[to]; !exists {
			locals[to] = Node{to, []Node{}, false}
		}

		locals[from].Neighbors = append(locals[from].Neighbors, locals[to])
		locals[to].Neighbors = append(locals[to].Neighbors, locals[from])

	}
}

func Part1(reader input.InputReader) int {
	formatInput(reader.GetInput(day))
}

func Part2(reader input.InputReader) int {
	return 0
}
