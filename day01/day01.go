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

	for i := 0; i < len(splitText); i++ {
		currentNum, err := strconv.Atoi(splitText[i])
		if err != nil {
			// skip non-number lines
			continue
		}
		_, present := complimentMap[currentNum]
		if present {
			answer := currentNum * complimentMap[currentNum]
			fmt.Println(answer)
		} else {
			// Set the key to the compliment so that we can find this value later
			complimentMap[2020-currentNum] = currentNum
		}
	}
}
