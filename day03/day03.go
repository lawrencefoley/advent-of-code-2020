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

func getTreesHit(mapData *[][]bool, slopeX int, slopeY int) int {
	currentX := 0
	currentY := 0
	totalTrees := 0

	for currentY < len(*mapData) {
		if coordinateIsTree(mapData, currentX, currentY) {
			totalTrees++
		}
		currentX += slopeX
		currentY += slopeY
	}

	return totalTrees
}

func main() {
	fmt.Println("day03")

	mapData := readMapFile("input.txt")
	fmt.Printf("part 1: %d\n", getTreesHit(&mapData, 3, 1))

	var partTwoSlopes [][]int = make([][]int, 5)
	partTwoSlopes[0] = []int{1, 1}
	partTwoSlopes[1] = []int{3, 1}
	partTwoSlopes[2] = []int{5, 1}
	partTwoSlopes[3] = []int{7, 1}
	partTwoSlopes[4] = []int{1, 2}

	partTwoTotal := 1
	for _, slope := range partTwoSlopes {
		partTwoTotal *= getTreesHit(&mapData, slope[0], slope[1])
	}

	fmt.Printf("part 2: %d\n", partTwoTotal)
}
