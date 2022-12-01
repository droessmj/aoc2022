package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func parseInput() []int {
	var input []int
	
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ! (len(scanner.Text()) == 0) {
			scannerInt, _ := strconv.Atoi(scanner.Text())
			input = append(input, scannerInt)
			
		} else {
			input = append(input, -1)

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func solvePart1(input []int) int {
	var runningTotal int = 0 
	var result int = 0

	for _, e := range input {
		if e != -1 {
			runningTotal += e
		} else {
			if runningTotal > result {
				result = runningTotal
			}
			runningTotal = 0
		}
	}

	if runningTotal > result {
		result = runningTotal
	}

	return result
}

func solvePart2(input []int) int {
	var runningTotal int = 0 
	var elfResults []int = make([]int, 0)

	for _, e := range input {
		if e != -1 {
			runningTotal += e
		} else {
			elfResults = append(elfResults, runningTotal)
			runningTotal = 0
		}

	}

	if runningTotal != 0 {
		elfResults = append(elfResults, runningTotal)
	}

	sort.Slice(elfResults, func(i, j int) bool {
		return elfResults[i] > elfResults[j]
	})

	var resultSum int = 0 
	for i := 0; i < 3; i++ {
		resultSum += elfResults[i]
	}

	return resultSum
}

func main() {

	input := parseInput()

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(input)
	fmt.Println(resultPt2)

}