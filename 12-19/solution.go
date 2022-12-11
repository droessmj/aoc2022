package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() []int {
	var input []int

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
	return 0
}

func solvePart2(input []int) int {

	return 0
}

func main() {

	input := parseInput()

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(input)
	fmt.Println(resultPt2)

}
