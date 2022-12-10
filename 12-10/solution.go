package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	operation Op
	cycles    int
	val       int
}

type Op string

const (
	noop Op = "noop"
	addx    = "addx"
)

func parseInput() []Move {
	var input []Move

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			pieces := strings.Split(scanner.Text(), " ")
			var m Move
			if len(pieces) > 1 {
				scannerInt, _ := strconv.Atoi(pieces[1])
				m = Move{cycles: 2, val: scannerInt, operation: addx}
			} else {
				m = Move{cycles: 1, val: 0, operation: noop}
			}

			input = append(input, m)

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func evalSample(cycleIndex int, modOffset int) bool {
	if cycleIndex%(20*modOffset) == 0 {
		return true
	}
	return false
}

func solve(input []Move) int {
	var x int = 1

	var modOffset int = 1
	var cycleIndex int = 1
	var finalScore int = 0

	var screen [][]string
	row := make([]string, 40)
	var rowX int = 0

	for _, m := range input {
		switch m.operation {

		case noop:
			tick(&row, &screen, &cycleIndex, &x, &modOffset, &rowX, &finalScore, 0, m.val)

			continue

		case addx:
			tick(&row, &screen, &cycleIndex, &x, &modOffset, &rowX, &finalScore, 0, m.val)
			tick(&row, &screen, &cycleIndex, &x, &modOffset, &rowX, &finalScore, 1, m.val)
		}
	}

	return finalScore
}

func tick(row *[]string, screen *[][]string, cycleIndex *int, x *int, modOffset *int, rowX *int, finalScore *int, iteration int, mVal int) {
	var t int = 0
	if *cycleIndex > 5 {
		t++
	}
	(*row)[*rowX] = evalScreen(*rowX, *x)
	if (*cycleIndex)%40 == 0 {
		fmt.Println(strings.Trim(fmt.Sprint(*row), "[]"))
		*row = make([]string, 40)
	}

	(*cycleIndex)++
	(*rowX)++
	if *rowX > 39 {
		*rowX = 0
	}

	if iteration > 0 {
		(*x) += mVal
	}

	var result bool
	result = evalSample(*cycleIndex, *modOffset)
	if result {
		(*modOffset) += 2
		(*finalScore) += ((*cycleIndex) * (*x))
	}

}

func evalScreen(cycleIndex, x int) string {
	if math.Abs(float64(cycleIndex-x)) <= 1 {
		return "#"
	} else {
		return "."
	}
}

func main() {

	input := parseInput()

	result := solve(input)
	fmt.Println(result)

}
