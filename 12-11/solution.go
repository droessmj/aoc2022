package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []int64
	opString        string
	opPieces        []string
	opOperation     string
	testString      string
	testVal         int
	testTrueTarget  int
	testFalseTarget int
	inspections     int
}

func NewMonkeyFromLines(lines []string) Monkey {
	//var m Monkey

	var items []int64
	var opString string
	var opOperation string
	var opPieces []string
	var testString string
	var testVal int
	var testTrueTarget int
	var testFalseTarget int

	// parse items
	pieces := strings.Split(strings.Split(strings.Trim(lines[1], " "), ":")[1], ",")
	for _, e := range pieces {
		intVal, _ := strconv.Atoi(strings.Trim(e, " "))
		items = append(items, int64(intVal))
	}

	//parse operation
	opString = strings.Trim(lines[2], " ")
	opPieces = strings.Split(opString, ":")
	opOperation = opPieces[1]
	opPieces = strings.Split(opOperation, " ")

	//parse test condition
	testString = strings.Trim(lines[3], " ")
	testVal, _ = strconv.Atoi(strings.Split(testString, " ")[3])
	testTrueTargetString := strings.Trim(lines[4], " ")
	testTrueTarget, _ = strconv.Atoi(strings.Split(testTrueTargetString, " ")[5])
	testFalseTargetString := strings.Trim(lines[5], " ")
	testFalseTarget, _ = strconv.Atoi(strings.Split(testFalseTargetString, " ")[5])

	return Monkey{
		items:           items,
		opString:        opString,
		opPieces:        opPieces,
		opOperation:     opOperation,
		testString:      testString,
		testVal:         testVal,
		testTrueTarget:  testTrueTarget,
		testFalseTarget: testFalseTarget,
	}
}

func Toss(source *Monkey, target *Monkey, newWorryLevel int64) {
	_, source.items = source.items[0], source.items[1:] // pop item off source
	target.items = append(target.items, newWorryLevel)
}

func CalcNewWorryLevel(m *Monkey, curWorryLevel int64, modulo int64) int64 {
	/*
		- First, you need to multiply all the monkeys' mod values together.
		- After every operation, mod that value with M.
	*/

	var result int64 = curWorryLevel
	var intVal int64
	if m.opPieces[5] == "old" {
		intVal = curWorryLevel
	} else {
		i, _ := strconv.Atoi(m.opPieces[5])
		intVal = int64(i)
	}

	switch m.opPieces[4] {
	case "*":
		result *= intVal
	case "+":
		result += intVal
	}

	remainder := result % modulo

	return remainder
}

func parseInput() []*Monkey {
	var inputs []*Monkey
	var lines []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			lines = append(lines, scanner.Text())
		} else {
			m := NewMonkeyFromLines(lines)
			inputs = append(inputs, &m)
			lines = make([]string, 0)
			continue
		}
	}
	m := NewMonkeyFromLines(lines)
	inputs = append(inputs, &m) //capture last guy

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return inputs
}

func solve(inputMonkeys []*Monkey, rounds int, reduceWorryLevel bool) int {

	var modulo int64 = 1
	for _, m := range inputMonkeys {
		modulo *= int64(m.testVal)
	}

	for i := 0; i < rounds; i++ {
		for _, m := range inputMonkeys {
			//inspect item
			for _, item := range m.items {
				m.inspections++

				//worry level multiplied
				newWorryLevel := CalcNewWorryLevel(m, item, modulo)

				//woryy level divided by 3
				if reduceWorryLevel {
					newWorryLevel = newWorryLevel / 3
				}

				//perform test
				if newWorryLevel%int64(m.testVal) == 0 {
					//newWorryLevel = newWorryLevel / int64(m.testVal)
					// toss to new monkey with new value
					Toss(m, inputMonkeys[m.testTrueTarget], newWorryLevel)
				} else {
					Toss(m, inputMonkeys[m.testFalseTarget], newWorryLevel)
				}

			} // else turn passes
		}
	}

	// get most active monkeys
	var topInspectionCount = 0
	var topInspectionIdx = 0

	var secondInspectionCount = 0
	var secondInspectionIdx = 0
	for idx, m := range inputMonkeys {
		if m.inspections > topInspectionCount {
			tempCount := topInspectionCount
			tempIdx := topInspectionIdx

			topInspectionCount = m.inspections
			topInspectionIdx = idx

			if tempCount > secondInspectionCount { // should always be true?
				secondInspectionCount = tempCount
				secondInspectionIdx = tempIdx
			}
		} else if m.inspections > secondInspectionCount {
			secondInspectionCount = m.inspections
			secondInspectionIdx = idx
		}
	}

	return inputMonkeys[topInspectionIdx].inspections * inputMonkeys[secondInspectionIdx].inspections
}

func solvePart2(input []int) int {

	return 0
}

func main() {
	input := parseInput()
	//fmt.Println(input)

	resultPt1 := solve(input, 20, true)
	fmt.Println(resultPt1)

	input = parseInput()
	resultPt2 := solve(input, 10000, false)
	fmt.Println(resultPt2)
}
