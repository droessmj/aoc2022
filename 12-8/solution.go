package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() [][]int {
	var input [][]int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		if !(len(scanner.Text()) == 0) {
			for _, el := range scanner.Text() {
				i, _ := strconv.Atoi(string(el))
				row = append(row, i)
			}
		}
		input = append(input, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func evalSubSlice(inputs []int, val int) bool {
	for _, el := range inputs {
		if el >= val {
			return false
		}
	}
	return true
}

func solve(inputs [][]int) int {
	var count int = 0

	for rowIdx, rEl := range inputs {
		for colIdx, colEl := range rEl {

			var up bool = false
			var down bool = false
			var left bool = false
			var right bool = false

			if rowIdx == 0 || rowIdx == len(inputs)-1 {
				count++
				continue
			} else if colIdx == 0 || colIdx == len(rEl)-1 {
				count++
				continue
			}

			colSliceUp := []int{}
			colSliceDown := []int{}
			for i := 0; i < len(inputs); i++ {
				if i < rowIdx {
					colSliceUp = append(colSliceUp, inputs[i][colIdx])
				} else if i > rowIdx {
					colSliceDown = append(colSliceDown, inputs[i][colIdx])
				}
			}
			up = evalSubSlice(colSliceUp, colEl)
			if up {
				count++
				continue
			}
			down = evalSubSlice(colSliceDown, colEl) //eval down
			if down {
				count++
				continue
			}

			left = evalSubSlice(inputs[rowIdx][0:colIdx], colEl) //eval left
			if left {
				count++
				continue
			}

			if colIdx+1 < len(rEl) {
				right = evalSubSlice(inputs[rowIdx][colIdx+1:], colEl) //eval right
				if right {
					count++
					continue
				}
			}
		}
	}

	return count
}

func main() {

	inputs := parseInput()
	//fmt.Println(inputs)

	result1 := solve(inputs)
	fmt.Println(result1)

	// result2 := solve2(inputs)
	// fmt.Println(result2)
}
