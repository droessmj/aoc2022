package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]string {
	var input [][]string
	
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ! (len(scanner.Text()) == 0) {
			roundStrings := strings.Split(scanner.Text(), " ")
			input = append(input, roundStrings)
		} 
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

type Throw int
const (
	Rock Throw = 1
	Paper Throw = 2
	Scissors Throw = 3
)

func buildEvalRound(r []string) int {
	var p1Throw Throw
	var p2Throw Throw
	var score int = 0 

	switch r[0] {
		case "A":
			p1Throw = Rock
		case "B":
			p1Throw = Paper
		case "C": 
			p1Throw = Scissors
	}

	switch r[1] {
		case "X":
			p2Throw = Rock
		case "Y":
			p2Throw = Paper
		case "Z": 
			p2Throw = Scissors
	}

	switch {
		case p1Throw == p2Throw:
			score += 3
		case p1Throw == Rock && p2Throw == Scissors:
			score += 0
		case p1Throw == Rock && p2Throw == Paper:
			score += 6
		case p1Throw == Paper && p2Throw == Rock:
			score += 0 
		case p1Throw == Paper && p2Throw == Scissors:
			score += 6
		case p1Throw == Scissors && p2Throw == Paper:
			score += 0
		case p1Throw == Scissors && p2Throw == Rock:
			score += 6
	}

	score += int(p2Throw)

	return score
}

func solvePart1(input [][]string) int {

	var score int = 0 

	for _, e := range input {
		score += buildEvalRound(e)
	}

	return score
}


func main() {

	inputs := parseInput()

	resultPt1 := solvePart1(inputs)
	fmt.Println(resultPt1)

	// resultPt2 := solvePart2(input)
	// fmt.Println(resultPt2)

}