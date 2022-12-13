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

	var name string = "input.test"
	if len(os.Args) > 1 {
		name = "input.txt"
	}
	file, err := os.Open(name)
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

type Packet struct {
	values []interface{}
}

func NewPacket(s string) Packet {
	var packetValues []interface{}
	var curList []interface{}
	headerStack := arraystack.New()

	stringPieces := strings.Split(s, ",")
	for _, piece := range stringPieces {
		i, ok := strconv.Atoi(string(piece))

		switch {

		case ok == nil:
			curList = append(curList, i)

		case strings.Contains()'[':
			curList = make([]interface{}, 0)

		case r[0] == ']':
			//Pop?
			packetValues = append(packetValues, curList...)

		default:
			panic("default hit")
		}
	}

	return Packet{values: packetValues}
}

func (left Packet) LessThanEqualTo(right Packet) bool {
	for leftIdx, leftVal := range left.values {

		if len(right.values) < leftIdx {
			return false
		}

		rightVal := right.values[leftIdx]

		switch leftVal.(type) {
		case int:
			if rightVal.(int) < leftVal.(int) {
				return false
			}

		case []int:
		}
	}

	return true
}

func solvePart1(input []string) int {
	var correctCount int = 0

	for i := 0; i < len(input); i += 2 {
		left := NewPacket(input[i])
		right := NewPacket(input[i+1])

		if left.LessThanEqualTo(right) {
			correctCount += (i + 1)

			fmt.Println(left, right, i+1)
		}
	}

	return correctCount
}

func solvePart2(input []string) int {

	return 0
}

func main() {

	//fmt.Println(os.Args[1])
	input := parseInput()
	//fmt.Println(input)

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(input)
	fmt.Println(resultPt2)

}
