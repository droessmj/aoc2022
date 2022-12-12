package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourbasic/graph"
)

type Point struct {
	row int
	col int
}

func parseInput() [][]byte {
	var input [][]byte

	var name string = "input.test"
	if len(os.Args) > 1 {
		name = "input.txt"
	}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			input = append(input, []byte(scanner.Text()))

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func getAdjacentSteps(input [][]byte, rowIdx int, colIdx int) []Point {
	var adjacent []Point

	// bias left/right!
	if colIdx == 0 {
		// no left
		adjacent = append(adjacent, Point{row: rowIdx, col: colIdx + 1})
	} else if colIdx == len(input[0])-1 {
		// no right
		adjacent = append(adjacent, Point{row: rowIdx, col: colIdx - 1})
	} else {
		// both
		adjacent = append(adjacent, Point{row: rowIdx, col: colIdx + 1})
		adjacent = append(adjacent, Point{row: rowIdx, col: colIdx - 1})
	}

	if rowIdx == 0 {
		// no up
		adjacent = append(adjacent, Point{row: rowIdx + 1, col: colIdx})
	} else if rowIdx == len(input)-1 {
		// no down
		adjacent = append(adjacent, Point{row: rowIdx - 1, col: colIdx})
	} else {
		// both
		adjacent = append(adjacent, Point{row: rowIdx + 1, col: colIdx})
		adjacent = append(adjacent, Point{row: rowIdx - 1, col: colIdx})
	}

	return adjacent
}

func getPointForValue(input [][]byte, target byte) Point {
	for rIdx, row := range input {
		for cIdx, el := range row {
			if el == target {
				return Point{row: rIdx, col: cIdx}
			}
		}
	}
	panic("Failure to find target byte!")
}

func getGraphIndexFromPoint(input [][]byte, p Point) int {
	return getGraphIndexFromRowCol(input, p.row, p.col)
}

func getGraphIndexFromRowCol(input [][]byte, row int, col int) int {
	colWidth := len(input[0])
	return col + (row * colWidth)
}

func solvePart1(input [][]byte) int {
	g := graph.New(len(input) * len(input[0]))

	// scan for start value
	startPoint := getPointForValue(input, []byte("S")[0])
	targetPoint := getPointForValue(input, []byte("E")[0])

	input[startPoint.row][startPoint.col] = byte('a') // set for comparisons
	input[targetPoint.row][targetPoint.col] = byte('z')

	for rowIdx, row := range input {
		for colIdx, bVal := range row {
			adjacencies := getAdjacentSteps(input, rowIdx, colIdx)
			for _, adjacency := range adjacencies {
				if bVal >= input[adjacency.row][adjacency.col] ||
					input[adjacency.row][adjacency.col]-bVal == 1 {
					g.AddCost(getGraphIndexFromRowCol(input, rowIdx, colIdx), getGraphIndexFromPoint(input, adjacency), 1)
				}
			}
		}
	}

	startIdx := getGraphIndexFromPoint(input, startPoint)
	targetIdx := getGraphIndexFromPoint(input, targetPoint)

	path, dist := graph.ShortestPath(g, startIdx, targetIdx)
	fmt.Println("Path", path, "Distance", dist)

	return int(dist)
}

func main() {

	input := parseInput()
	resultPt1 := solvePart1(input)

	fmt.Println(resultPt1)
}
