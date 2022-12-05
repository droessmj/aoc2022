package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

func parseInput() []string {
	var input []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			input = append(input, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func processMove(line string, stacks []arraystack.Stack) {
	// move 1 from 2 to 1
	pieces := strings.Split(line, " ")
	count, _ := strconv.Atoi(pieces[1])
	startIdx, _ := strconv.Atoi(pieces[3])
	endIdx, _ := strconv.Atoi(pieces[5])
	startIdx-- //offset index
	endIdx--

	for i := 0; i < count; i++ {
		val, _ := stacks[startIdx].Pop()
		stacks[endIdx].Push(val)
	}
}

func processMovePt2(line string, stacks []arraystack.Stack) {
	// move 1 from 2 to 1
	pieces := strings.Split(line, " ")
	count, _ := strconv.Atoi(pieces[1])
	startIdx, _ := strconv.Atoi(pieces[3])
	endIdx, _ := strconv.Atoi(pieces[5])
	startIdx-- //offset index
	endIdx--

	// move as unit
	if count == 1 {
		val, _ := stacks[startIdx].Pop()
		stacks[endIdx].Push(val)
	} else {
		tempStack := arraystack.New()
		for i := 0; i < count; i++ {
			val, _ := stacks[startIdx].Pop()
			tempStack.Push(val)
		}
		p, _ := tempStack.Peek()
		for p != nil {
			val, _ := tempStack.Pop()
			stacks[endIdx].Push(val)
			p, _ = tempStack.Peek()
		}
	}

}

func solvePart2(inputs []string, stacks []arraystack.Stack) string {
	for _, l := range inputs {
		processMovePt2(l, stacks)
	}

	//final String
	var result string
	for _, s := range stacks {
		val, _ := s.Peek()
		result += val.(string)
	}

	return result

}

func solvePart1(inputs []string, stacks []arraystack.Stack) string {

	for _, l := range inputs {
		processMove(l, stacks)
	}

	//final String
	var result string
	for _, s := range stacks {
		val, _ := s.Peek()
		result += val.(string)
	}

	return result
}

func main() {

	inputs := parseInput()

	/*
			[C]             [L]         [T]
			[V] [R] [M]     [T]         [B]
			[F] [G] [H] [Q] [Q]         [H]
			[W] [L] [P] [V] [M] [V]     [F]
			[P] [C] [W] [S] [Z] [B] [S] [P]
		[G] [R] [M] [B] [F] [J] [S] [Z] [D]
		[J] [L] [P] [F] [C] [H] [F] [J] [C]
		[Z] [Q] [F] [L] [G] [W] [H] [F] [M]
		1   2   3   4   5   6   7   8   9

		0   1   2   3   4   5   6   7   8
	*/
	stacks := make([]arraystack.Stack, 9)
	stacks[0] = *arraystack.New()
	stacks[0].Push("Z")
	stacks[0].Push("J")
	stacks[0].Push("G")

	stacks[1] = *arraystack.New()
	stacks[1].Push("Q")
	stacks[1].Push("L")
	stacks[1].Push("R")
	stacks[1].Push("P")
	stacks[1].Push("W")
	stacks[1].Push("F")
	stacks[1].Push("V")
	stacks[1].Push("C")

	stacks[2] = *arraystack.New()
	stacks[2].Push("F")
	stacks[2].Push("P")
	stacks[2].Push("M")
	stacks[2].Push("C")
	stacks[2].Push("L")
	stacks[2].Push("G")
	stacks[2].Push("R")

	stacks[3] = *arraystack.New()
	stacks[3].Push("L")
	stacks[3].Push("F")
	stacks[3].Push("B")
	stacks[3].Push("W")
	stacks[3].Push("P")
	stacks[3].Push("H")
	stacks[3].Push("M")

	stacks[4] = *arraystack.New()
	stacks[4].Push("G")
	stacks[4].Push("C")
	stacks[4].Push("F")
	stacks[4].Push("S")
	stacks[4].Push("V")
	stacks[4].Push("Q")

	stacks[5] = *arraystack.New()
	stacks[5].Push("W")
	stacks[5].Push("H")
	stacks[5].Push("J")
	stacks[5].Push("Z")
	stacks[5].Push("M")
	stacks[5].Push("Q")
	stacks[5].Push("T")
	stacks[5].Push("L")

	stacks[6] = *arraystack.New()
	stacks[6].Push("H")
	stacks[6].Push("F")
	stacks[6].Push("S")
	stacks[6].Push("B")
	stacks[6].Push("V")

	stacks[7] = *arraystack.New()
	stacks[7].Push("F")
	stacks[7].Push("J")
	stacks[7].Push("Z")
	stacks[7].Push("S")

	stacks[8] = *arraystack.New()
	stacks[8].Push("M")
	stacks[8].Push("C")
	stacks[8].Push("D")
	stacks[8].Push("P")
	stacks[8].Push("F")
	stacks[8].Push("H")
	stacks[8].Push("B")
	stacks[8].Push("T")

	/*
				[D]
			[N] [C]
			[Z] [M] [P]
			1   2   3

		stacks := make([]arraystack.Stack, 3)
		stacks[0] = *arraystack.New()
		stacks[0].Push("Z")
		stacks[0].Push("N")

		stacks[1] = *arraystack.New()
		stacks[1].Push("M")
		stacks[1].Push("C")
		stacks[1].Push("D")

		stacks[2] = *arraystack.New()
		stacks[2].Push("P")
	*/
	//resultPt1 := solvePart1(inputs, stacks)
	//fmt.Println(resultPt1)

	resultPt2 := solvePart2(inputs, stacks)
	fmt.Println(resultPt2)
}
