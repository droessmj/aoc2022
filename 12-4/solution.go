package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ElfWorkload struct {
	start int
	end   int
}

func parseInput() [][]string {
	var input [][]string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			roundStrings := strings.Split(scanner.Text(), ",")
			input = append(input, roundStrings)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func compareWorkloadOverlap(e1 ElfWorkload, e2 ElfWorkload) bool {

	if e1.start >= e2.start && e1.end <= e2.end {
		return true
	} else if e2.start >= e1.start && e2.end <= e1.end {
		return true
	}

	return false
}

func compareWorkloadOverlapPart2(e1 ElfWorkload, e2 ElfWorkload) bool {
	a := math.Max(float64(e1.start), float64(e2.start))
	b := math.Min(float64(e1.end), float64(e2.end))

	return a <= b
}

func solvePart1(inputs [][]string) int {
	var count int
	for _, e := range inputs {

		e1Parts := strings.Split(e[0], "-")
		start, _ := strconv.Atoi(e1Parts[0])
		end, _ := strconv.Atoi(e1Parts[1])
		elf1 := ElfWorkload{start: start, end: end}

		e2Parts := strings.Split(e[1], "-")
		start, _ = strconv.Atoi(e2Parts[0])
		end, _ = strconv.Atoi(e2Parts[1])
		elf2 := ElfWorkload{start: start, end: end}

		if result := compareWorkloadOverlap(elf1, elf2); result == true {
			count++
		}
	}

	return count
}

func solvePart2(inputs [][]string) int {
	var count int
	for _, e := range inputs {

		e1Parts := strings.Split(e[0], "-")
		start, _ := strconv.Atoi(e1Parts[0])
		end, _ := strconv.Atoi(e1Parts[1])
		elf1 := ElfWorkload{start: start, end: end}

		e2Parts := strings.Split(e[1], "-")
		start, _ = strconv.Atoi(e2Parts[0])
		end, _ = strconv.Atoi(e2Parts[1])
		elf2 := ElfWorkload{start: start, end: end}

		if result := compareWorkloadOverlapPart2(elf1, elf2); result == true {
			count++
		}
	}

	return count
}

func main() {

	inputs := parseInput()

	resultPt1 := solvePart1(inputs)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(inputs)
	fmt.Println(resultPt2)
}
