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

const (
	Win int = 6
	Draw int = 3
	Lose int = 0
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

func buildEvalRound2(r []string) int {
	var p1Throw Throw
	var p2Throw Throw
	var outcome int
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
			outcome = Lose
		case "Y":
			outcome = Draw
		case "Z": 
			outcome = Win
	}

	switch {
		case outcome == Draw:
			score += 3
			switch p1Throw {
				case Rock:
					score += 1
				case Paper:
					score += 2
				case Scissors:
					score += 3
			}
		case p1Throw == Rock && outcome == Lose:
			score += 3
		case p1Throw == Rock && outcome == Win:
			score += 8
		case p1Throw == Paper && outcome == Lose:
			score += 1
		case p1Throw == Paper && outcome == Win:
			score += 9
		case p1Throw == Scissors && outcome == Lose:
			score += 2
		case p1Throw == Scissors && outcome == Win:
			score += 7
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

func solvePart2(input [][]string) int {

	var score int = 0 

	for _, e := range input {
		score += buildEvalRound2(e)
	}

	return score
}

func main() {

	inputs := parseInput()

	resultPt1 := solvePart1(inputs)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(inputs)
	fmt.Println(resultPt2)
}