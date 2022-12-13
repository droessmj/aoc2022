package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"github.com/emirpasic/gods/stacks/arraystack"
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

func AddValsFromString(packet *[]interface{}, s string) {
	vals := strings.Split(strings.Trim(s, ","), ",")
	for _, val := range vals {
		// should all be ints?
		i, err := strconv.Atoi(val)
		if err != nil {
			i = -1 //math.MinInt
		}
		*packet = append(*packet, i) // add to front!
	}

}

func ParsePacketFromListString(s string) []interface{} {

	var packet []interface{}

	if strings.Contains(s, "[") || strings.Contains(s, "]") {
		// case 2
		listLeftIdx := strings.Index(s, "[") + 1
		listRightIdx := strings.LastIndex(s, "]")

		if listLeftIdx != 0 {
			// sub case -- multiple packets in one layer
			packets := strings.Split(s, "],")

			prePiece := s[0 : listLeftIdx-1]
			if prePiece != "" {
				AddValsFromString(&packet, prePiece)

				closeBracketIdx2 := strings.LastIndex(s, "]")
				if closeBracketIdx2 > len(s)-1 {
					closeBracketIdx2 = len(s) - 1
				}

				packet = append(packet, ParsePacketFromListString(s[listLeftIdx:closeBracketIdx2]))

				if closeBracketIdx2 > 0 && closeBracketIdx2 < len(s) {
					ints2 := strings.Trim(s[closeBracketIdx2:], "]")
					if ints2 != "" {
						AddValsFromString(&packet, ints2)
					}
				}

			} else {
				for _, p := range packets {
					if strings.ContainsAny(p, "[]") {
						// need to trim last one only, not ALL
						closeBracketIdx := strings.LastIndex(p, "]") - 1
						openBracketIdx := strings.Index(p, "[")

						if openBracketIdx == 0 {
							p = p[1:] //trim
						} else if openBracketIdx > 0 {
							// prune front, then recurse
							ints := p[0:openBracketIdx]
							AddValsFromString(&packet, ints)
							p = p[openBracketIdx:]
						}

						if closeBracketIdx > -1 {
							p = p[0:closeBracketIdx]
						}
						//listPiece := strings.Trim(strings.Trim(p, "]"), "[")
						packet = append(packet, ParsePacketFromListString(p))
					} else {
						AddValsFromString(&packet, p)
					}
				}
			}
		} else {
			packet = append(packet, ParsePacketFromListString(s[listLeftIdx:listRightIdx]))
		}
	} else {
		// case 1
		AddValsFromString(&packet, s)
	}

	return packet
}

func NewPacket(s string) Packet {

	listLeftIdx := strings.Index(s, "[") + 1
	listRightIdx := strings.LastIndex(s, "]")
	packet := ParsePacketFromListString(s[listLeftIdx:listRightIdx])

	return Packet{values: packet}
}

func (left Packet) LessThanEqualTo(right Packet) bool {

	return true
	// for leftIdx, leftVal := range left.values {

	// 	if len(right.values) < leftIdx {
	// 		return false
	// 	}

	// 	rightVal := right.values[leftIdx]

	// 	switch leftVal.(type) {
	// 	case int:
	// 		if rightVal.(int) < leftVal.(int) {
	// 			return false
	// 		}

	// 	case []int:
	// 	}
	// }

	// return true
}

func solvePart1(input []string) int {
	var correctCount int = 0

	for i := 0; i < len(input); i += 2 {
		left := NewPacket(input[i])
		right := NewPacket(input[i+1])

		fmt.Println(left)
		fmt.Println(right)
		fmt.Println("")
		if left.LessThanEqualTo(right) {
			correctCount += (i + 1)

			//fmt.Println(left, right, i+1)
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
