package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type Direction string

const (
	Up    Direction = "U"
	Right           = "R"
	Left            = "L"
	Down            = "D"
)

type Move struct {
	direction Direction
	steps     int
}

type Location struct {
	x int
	y int
}

func parseInput() []Move {
	var input []Move

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !(len(scanner.Text()) == 0) {
			pieces := strings.Split(scanner.Text(), " ")
			scannerInt, _ := strconv.Atoi(pieces[1])

			var direction Direction
			switch pieces[0] {
			case "U":
				direction = Up
			case "D":
				direction = Down
			case "R":
				direction = Right
			case "L":
				direction = Left
			}

			input = append(input, Move{direction: direction, steps: scannerInt})

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return input
}

func moveT(h Location, hPrev Location, t *Location, visitedLocations *hashset.Set, dir Direction) {
	//math on whether or not t should move

	//case - head one above
	//case - head one to the right
	//case - head one to left
	//case - head one below
	//case - diagonal any of the 4 directions

	if math.Abs(float64(h.x-t.x)) < 2 && math.Abs(float64(h.y-t.y)) < 2 {
		//do nothing
		return
	} else if math.Abs(float64(h.x-t.x)) > 1 && math.Abs(float64(h.y-t.y)) == 1 ||
		math.Abs(float64(h.x-t.x)) == 1 && math.Abs(float64(h.y-t.y)) > 1 {
		t.y = hPrev.y
		t.x = hPrev.x
		visitedLocations.Add(Location{x: t.x, y: t.y})
	} else {
		//case -- not touching
		switch dir {
		case Up:
			t.y++
			//diagonal case
		case Down:
			t.y--
			//diagonal case
		case Right:
			t.x++
			//diagonal case
		case Left:
			t.x--
			//diagonal
		}

		visitedLocations.Add(Location{x: t.x, y: t.y})
	}

}

func moveH(move Move, h *Location, t *Location, visitedLocations *hashset.Set) {
	for i := 0; i < move.steps; i++ {
		hPrev := Location{x: h.x, y: h.y}
		switch move.direction {
		case Up:
			h.y++
			moveT(*h, hPrev, t, visitedLocations, Up)
		case Down:
			h.y--
			moveT(*h, hPrev, t, visitedLocations, Down)
		case Left:
			h.x--
			moveT(*h, hPrev, t, visitedLocations, Left)
		case Right:
			h.x++
			moveT(*h, hPrev, t, visitedLocations, Right)
		}
		//fmt.Println("   h", h, "t", t, "direction", move.direction)
	}
}

func solvePart1(input []Move) int {

	visitedLocations := hashset.New()

	//initial state is 0,0 on grid, len(inputs) - 1, 0 in 2d array
	h := Location{x: 0, y: 0}
	t := Location{x: 0, y: 0}
	visitedLocations.Add(Location{x: t.x, y: t.y})

	for _, move := range input {
		//fmt.Println("i", i, "h", h, "t", t)
		moveH(move, &h, &t, visitedLocations)
	}

	return visitedLocations.Size()
}

func solvePart2(input []Move) int {

	return 0
}

func main() {
	input := parseInput()
	//fmt.Println(input)

	resultPt1 := solvePart1(input)
	fmt.Println(resultPt1)

	// resultPt2 := solvePart2(input)
	// fmt.Println(resultPt2)
}
