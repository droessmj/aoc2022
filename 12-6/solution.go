package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
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
	var t string
	var seenRune []rune

	for idx1, el := range input {
		t += string(el)
		if idx1 < uniqChars {
			continue
		} else {
			t = t[1:]
		}

		// could swap this to be a sliding window rather than a for - for
		for idx2, e := range t {
			if idx2 == 0 {
				seenRune = seenRune[:0]
			}
			//fmt.Println(seenRune)

			if result := slices.Contains(seenRune, e); result {
				break
			}
			seenRune = append(seenRune, e)

			if idx2 == (uniqChars - 1) {
				return idx1
			}

		}
	}

	return -1
}

func main() {

	inputs := parseInput()
	//fmt.Println(inputs)

	resultPt1 := solve(inputs, 4) + 1
	fmt.Println(resultPt1)

	resultPt2 := solve(inputs, 14) + 1
	fmt.Println(resultPt2)
}
