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

func LessThanEqualTo(left []interface{}, right []interface{}, index int) bool {
	// do we need a depth?

	if len(right) < index {
		//false if right is smaller, and if we can't check it...
		return false
	}

	switch left[index].(type) {
	case int:
		rightInt := GetNextInt(right, index)
		if rightInt == -1 {
			return false
		}

		if left[index].(int) < rightInt {
			return true
		} else if left[index].(int) == rightInt {
			if index+1 < len(left) {
				return LessThanEqualTo(left, right, index+1)
			} else {
				return true
			}

		} else {
			return false
		}

	case []interface{}:
		left := left[index].([]interface{})
		right := right[index]
		switch right.(type) {
		case int:
			return false
		case []interface{}:
			return LessThanEqualTo(left, right.([]interface{}), index)
		default:
			panic("unplanned type")
		}

	default:
		panic("unhandled case")
	}
}

func GetNextInt(val []interface{}, index int) int {

	if len(val) <= index {
		//false if right is smaller, and if we can't check it...
		return -1
	}

	switch val[index].(type) {
	case []interface{}:
		return GetNextInt(val[index].([]interface{}), index)
	case int:
		return val[index].(int)
	default:
		return -1
	}
}

func solvePart1(input []string) int {
	var correctCount int = 0

	for i := 0; i < len(input); i += 2 {
		left := NewPacket(input[i])
		right := NewPacket(input[i+1])

		if LessThanEqualTo(left.values, right.values, 0) {
			correctCount += (i / 2) + 1
			fmt.Println(left, right, (i/2)+1)
			fmt.Println()
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
