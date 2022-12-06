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
		if !(len(scanner.Text()) == 0) {
			input = append(input, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func solve(inputs []string, uniqChars int) int {
	var input string = inputs[0]
	var seenRune = make([]rune, uniqChars)
	runeMap := make(map[rune]int)
	var count int

	for idx, el := range input {

		if idx < uniqChars {
			seenRune[idx] = el
			runeMap[el] += 1
			continue
		} else {
			var dupes bool = false
			for _, r := range seenRune {
				count = runeMap[r]
				if count > 1 {
					dupes = true
				}
			}
			if dupes == false {
				return idx
			} else {
				runeMap[seenRune[0]]-- //remove from head

				//shift to left
				for each := range seenRune {
					if each == 0 {
						continue
					}

					seenRune[each-1] = seenRune[each]
				}

				//add new
				seenRune[uniqChars-1] = el
				runeMap[el]++
			}
		}
	}

	return -1
}

func main() {

	inputs := parseInput()
	//fmt.Println(inputs)

	resultPt1 := solve(inputs, 4)
	fmt.Println(resultPt1)

	resultPt2 := solve(inputs, 14)
	fmt.Println(resultPt2)
}
