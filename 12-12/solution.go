package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Point struct {
	row int
	col int
}

type Node struct {
	p     Point
	steps int
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

func solvePart1(input [][]byte) int {
	var minSteps int = math.MaxInt

	stack := arraystack.New()
	visitedSet := hashset.New()

	/*
	   procedure DFS_iterative(G, v) is
	       let S be a stack
	       S.push(v)
	       while S is not empty do
	           v = S.pop()
	           if v is not labeled as discovered then
	               label v as discovered
	               for all edges from v to w in G.adjacentEdges(v) do
	                   S.push(w)
	*/

	// scan for start value
	startPoint := getPointForValue(input, []byte("S")[0])
	stack.Push(Node{p: startPoint, steps: 0})

	targetPoint := getPointForValue(input, []byte("E")[0])
	input[startPoint.row][startPoint.col] = byte('a') // set for comparisons
	input[targetPoint.row][targetPoint.col] = byte('z')

	//loop
	for {
		temp, _ := stack.Pop()
		v := temp.(Node)

		if !visitedSet.Contains(v.p) {

			visitedSet.Add(v.p)

			if v.p.row == targetPoint.row && v.p.col == targetPoint.col {
				minSteps = v.steps
				continue
			}

			availableSteps := getAdjacentSteps(input, v.p.row, v.p.col)
			for _, e := range availableSteps {
				if math.Abs(float64(input[e.row][e.col]-input[v.p.row][v.p.col])) <= 1 &&
					v.steps+1 < minSteps {
					stack.Push(Node{p: e, steps: v.steps + 1})
				}
			}
		}

		if _, ok := stack.Peek(); !ok {
			break
		}
	}

	return minSteps
}

func main() {

	input := parseInput()

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	// resultPt2 := solvePart2(input)
	// fmt.Println(resultPt2)

}
