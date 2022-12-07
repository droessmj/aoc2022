package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	size   int
	name   string
	parent *Directory
}

type Directory struct {
	files       []File
	directories []*Directory
	parent      *Directory
	name        string
	size        int
}

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

func solve(inputs []string) int {

	rootDir := Directory{files: []File{}, directories: []*Directory{}, parent: nil, name: "root"}
	var curDir *Directory = &rootDir

	for _, el := range inputs {
		//fmt.Println(el)

		pieces := strings.Split(el, " ")
		switch pieces[0] {
		case "$":
			// command
			if pieces[1] == "cd" {
				if pieces[2] != ".." {
					newDir := Directory{files: []File{}, directories: []*Directory{}, parent: curDir, name: pieces[2]}
					curDir.directories = append(curDir.directories, &newDir)
					curDir = &newDir
				} else {
					curDir = curDir.parent
				}
			}
			//fmt.Println(curDir)

		case "dir":
			//fmt.Println("dir", el)
			//new dir
			continue
		default:
			//file
			//fmt.Println("file", el)
			fileSize, _ := strconv.Atoi(pieces[0])
			file := File{size: fileSize, name: pieces[1], parent: curDir}
			curDir.files = append(curDir.files, file)
		}
	}

	// calculate Directory size recursively
	calcSetDirSize(&rootDir)
	//fmt.Println(rootDir)

	//walk dirs to identify those in scope, add their sizes to score
	score := identifyDirs(&rootDir)

	return score
}

func identifyDirs(dir *Directory) int {
	var score int = 0
	for _, d := range dir.directories {
		if d.size <= 100000 {
			score += d.size
			score += identifyDirs(d)
		} else {
			if len(d.directories) > 0 {
				score += identifyDirs(d)
			}
		}
	}

	return score
}

func calcSetDirSize(dir *Directory) int {
	var size int = 0

	for _, d := range dir.directories {
		size += calcSetDirSize(d)
	}

	for _, f := range dir.files {
		size += f.size
	}

	dir.size = size

	return size
}

func main() {

	inputs := parseInput()
	//fmt.Println(inputs)

	resultPt1 := solve(inputs)
	fmt.Println(resultPt1)

	// resultPt2 := solve(inputs, 14)
	// fmt.Println(resultPt2)
}
