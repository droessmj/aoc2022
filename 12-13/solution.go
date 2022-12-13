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
	s = strings.Trim(strings.Trim(strings.Trim(s, "["), "]"), ",")
	if s != "" && s != "]" {
		vals := strings.Split(s, ",")
		for _, val := range vals {
			// should all be ints?
			i, err := strconv.Atoi(val)
			if err != nil {
				i = -1 //math.MinInt
			}
			*packet = append(*packet, i) // add to front!
		}
	}

}

func SplitSameLevel(s string) []string {
	// 1],[2,3,4
	// TODO - Round this out?
	pieces := strings.Split(s, "],[")
	for idx, p := range pieces {
		openIndex := strings.Index(p, "[")
		if openIndex > -1 {
			p = p[openIndex+1:]
		}

		closeIndex := strings.LastIndex(p, "]")
		if closeIndex > -1 {
			p = p[:closeIndex]
		}
		pieces[idx] = p
	}

	return pieces

}

func ParsePacketFromListString(s string) []interface{} {

	var packet []interface{}

	if strings.Contains(s, "[") {
		openIndex := strings.Index(s, "[")
		closeIndex := strings.LastIndex(s, "]")

		//need to do a split of items at the same level...then do the below on vals + packets
		packets := SplitSameLevel(s[openIndex : closeIndex+1])

		AddValsFromString(&packet, s[0:openIndex])
		for _, p := range packets {
			packet = append(packet, ParsePacketFromListString(p))
		}
		AddValsFromString(&packet, s[closeIndex:])

	} else {
		AddValsFromString(&packet, s)
	}

	return packet
}

func NewPacket(s string) Packet {
	openIndex := strings.Index(s, "[")
	closeIndex := strings.LastIndex(s, "]")
	packet := ParsePacketFromListString(s[openIndex+1 : closeIndex])

	return Packet{values: packet}
}

func LessThanEqualTo(left []interface{}, right []interface{}, index int) bool {
	// do we need a depth?

	if len(right) <= index {
		//false if right is smaller, and if we can't check it...
		return false
	} else if len(left) <= index && len(right) > index {
		return true
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

		fmt.Println(left)
		fmt.Println(right)
		fmt.Println()

		if LessThanEqualTo(left.values, right.values, 0) {
			correctCount += (i / 2) + 1
		}

		if i > 10 {
			break
		}
	}

	return correctCount
}

func solvePart2(input []string) int {

	return 0
}

func main() {
	input := parseInput()

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	resultPt2 := solvePart2(input)
	fmt.Println(resultPt2)
}
