package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readMapFile(filename string) [][]bool {
	var mapData = make([][]bool, 0)
	fileData, err := ioutil.ReadFile(filename)
	checkError(err)

	fileLines := strings.Split(string(fileData), "\n")
	for index, line := range fileLines {
		if strings.TrimSpace(line) != "" {
			mapData = append(mapData, make([]bool, 0))
			for _, currentChar := range line {
				if currentChar == '.' {
					mapData[index] = append(mapData[index], false)
				} else {
					// Tree
					mapData[index] = append(mapData[index], true)
				}
			}
		}
	}

	return mapData
}

func coordinateIsTree(mapData *[][]bool, x int, y int) bool {
	// Mod x by the length of the first row
	// since the x values repeat
	x %= len((*mapData)[0])

	return (*mapData)[y][x]
}

func main() {
	fmt.Println("day03")

	mapData := readMapFile("input.txt")

	slopeHorizontal := 3
	slopeVertical := 1

	currentX := 0
	currentY := 0

	totalTrees := 0
	for currentY < len(mapData) {
		if coordinateIsTree(&mapData, currentX, currentY) {
			totalTrees++
		}
		currentX += slopeHorizontal
		currentY += slopeVertical
	}

	fmt.Printf("part 1: %d\n", totalTrees)
}
