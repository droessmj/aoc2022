package main

import (
	"bufio"
	"fmt"
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

func evalSample(cycleIndex int, sampleCount int) bool {
	if cycleIndex%(20*sampleCount) == 0 {
		return true
	}
	return false
}

func solvePart1(input []Move) int {
	var x int = 1
	var sampleCount int = 1
	var cycleIndex int = 1

	samples := make([]int, 0)
	var result bool = false
	var sample int = 0
	var finalScore int = 0

	for _, m := range input {
		switch m.operation {
		case noop:
			cycleIndex++
			result = evalSample(cycleIndex, sampleCount)
			if result {
				samples = append(samples, sample)
				sampleCount += 2
				finalScore += (cycleIndex * x)
				fmt.Println("Index", cycleIndex, "X", x, "final", finalScore)
			}
			continue
		case addx:
			cycleIndex++
			result = evalSample(cycleIndex, sampleCount)
			if result {
				samples = append(samples, sample)
				sampleCount += 2
				finalScore += (cycleIndex * x)
				fmt.Println("Index", cycleIndex, "X", x, "final", finalScore)
			}

			x += m.val

			cycleIndex++
			result = evalSample(cycleIndex, sampleCount)
			if result {
				samples = append(samples, sample)
				sampleCount += 2
				finalScore += (cycleIndex * x)
				fmt.Println("Index", cycleIndex, "X", x, "final", finalScore)
			}
		}
	}

	return finalScore
}

func solvePart2(input []int) int {

	return 0
}

func main() {

	input := parseInput()
	//fmt.Println(input)

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	//resultPt2 := solvePart2(input)
	//fmt.Println(resultPt2)

}
