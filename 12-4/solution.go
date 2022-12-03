package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/emirpasic/gods/sets/hashset"
)

func parseInput() []string {
	var input []string
	
	file, err := os.Open("input.test")
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

		set := hashset.New()

		for _, r := range first {
			set.Add(r)
		}
		
		// this would be so much easier if there were built-in sets...
		seen := hashset.New()
		for _, r := range second {

			// if we've already seen this rune, continue
			// if _, ok := seen[r]; ok  {
			// 	continue
			// } else {
			// 	seen[r] = exists
			// }
			if seen.Contains(r) {
				continue
			} else {
				seen.Add(r)
			}

			if set.Contains(r)  {
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

func main() {

	inputs := parseInput()

	resultPt1 := solvePart1(inputs)
	fmt.Println(resultPt1)
}