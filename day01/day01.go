package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("day01")

	data, err := ioutil.ReadFile("input.txt")
	checkError(err)

	// Split on newlines
	splitText := strings.Split(string(data), "\n")

	complimentMap := make(map[int]int)

	for _, textVal := range splitText {
		currentNum, err := strconv.Atoi(textVal)
		if err != nil {
			// skip non-number lines
			continue
		}
		_, present := complimentMap[currentNum]
		if present {
			answer := currentNum * complimentMap[currentNum]
			fmt.Printf("part 1: %d\n", answer)
		} else {
			// Set the key to the compliment so that we can find this value later
			complimentMap[2020-currentNum] = currentNum
		}
	}

	fmt.Printf("part 2: %d\n", part2(splitText))
}

func part2(splitText []string) int {
	var nums = make([]int, 0)
	var err error
	var curVal int
	for _, val := range splitText {
		curVal, err = strconv.Atoi(val)
		if err != nil {
			// skip non-number lines
			continue
		}
		nums = append(nums, curVal)
	}

	// Yeah, I know
	var count int = 0
	for _, val1 := range nums {
		for _, val2 := range nums {
			for _, val3 := range nums {
				count++
				if (val1 + val2 + val3) == 2020 {
					return val1 * val2 * val3
				}
			}
		}
	}

	return 0
}
