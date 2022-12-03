package main

import (
	"bufio"
	"fmt"
	"os"
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
		if ! (len(scanner.Text()) == 0) {
			input = append(input, scanner.Text())
		} 
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func solvePart1(inputs []string) int {
	var score int = 0 
	var row int = 0 

	//lowercase a == 97 (1) == -96 offset
	//upercase A == 65 (27) == -38 offset
	const lowerOffset int = -96
	const upperOffset int = -38

	for _, e := range inputs {
		//var print bool = true

		midpoint := len(e) / 2
		runeArr := []rune(e)

		first := runeArr[0:midpoint]
		second := runeArr[midpoint:]

		set := make(map[rune]struct{})
		var exists = struct{}{}

		for _, r := range first {
			set[r] = exists
		}
		
		// this would be so much easier if there were built-in sets...
		seen := make(map[rune]struct{})
		for _, r := range second {

			// if we've already seen this rune, continue
			if _, ok := seen[r]; ok  {
				continue
			} else {
				seen[r] = exists
			}

			if _, ok := set[r]; ok  {
				//var t int
				if r < 95 {
					//t = (int(r) + upperOffset)
					score += (int(r) + upperOffset)
				} else {
					//t = (int(r) + lowerOffset)
					score += (int(r) + lowerOffset)
				}
				//fmt.Println("row:", row, " ", string(r), ":", t)
				//print = false
			}
		}
		
		row++
	}

	return score
}

func solvePart2(inputs []string) int {
	var score int = 0 
	var groupIdx int = 0 

	//lowercase a == 97 (1) == -96 offset
	//upercase A == 65 (27) == -38 offset
	const lowerOffset int = -96
	const upperOffset int = -38

	set := make(map[rune]int)
	var exists = struct{}{}

	for _, e := range inputs {
		seen := make(map[rune]struct{})

		for _, r := range e {
			// if we've already seen this rune, continue
			if _, ok := seen[r]; ok  {
				continue
			} else {
				seen[r] = exists
			}

			if groupIdx== 0 {
				set[r] = 1
			} else {
				if _, ok := set[r]; ok  {
					set[r]++
				} 
			}
		}

		groupIdx++
		if groupIdx > 2 {
			for r, val := range set {
				if val > 2 {
					//fmt.Println(r, val)
					if r < 95 {
						score += (int(r) + upperOffset)
					} else {
						score += (int(r) + lowerOffset)
					}
				}
			}

			groupIdx = 0
			// set meta-set to null
			set = make(map[rune]int)
		}
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