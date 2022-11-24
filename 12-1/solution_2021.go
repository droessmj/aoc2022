package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
		scannerInt, _ := strconv.Atoi(scanner.Text())
		input = append(input, scannerInt)
		//fmt.Println(scannerInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func main() {

	input := parseInput()

	//var allReadings := make([]int,0)
	var priorReading int = math.MinInt
	var increasedCount int = 0

	for _, currentReading := range input {
		//allReadings = append(allReadings, reading)
		if priorReading != math.MinInt && currentReading > priorReading {
			increasedCount += 1
		}
		priorReading = currentReading
	}

	fmt.Println(increasedCount)
}
