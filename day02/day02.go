package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	min      int
	max      int
	letter   string
	password string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	checkError(err)

	// Split on newlines
	splitText := strings.Split(string(data), "\n")

	// Remove blank lines
	var lines = make([]string, 0)
	for _, s := range splitText {
		if strings.TrimSpace(s) != "" {
			lines = append(lines, s)
		}
	}

	return lines
}

func parseLine(line string, passwordRegex *regexp.Regexp) *password {
	p := password{}

	var matches = passwordRegex.FindSubmatch([]byte(line))

	p.min, _ = strconv.Atoi(string(matches[1]))
	p.max, _ = strconv.Atoi(string(matches[2]))
	p.letter = string(matches[3])
	p.password = string(matches[4])

	return &p
}

func passwordIsValid(min int, max int, letter string, password string) bool {
	count := 0
	for _, c := range password {
		if string(c) == letter {
			count++
		}
	}

	// Password validation
	return count >= min && count <= max
}

func passwordIsValidPart2(index1 int, index2 int, letter string, password string) bool {
	count := 0
	var indexes = []int{index1, index2}

	for _, i := range indexes {
		if []rune(password)[i-1] == []rune(letter)[0] {
			count++
		}
	}

	// Password validation - only 1 occurance
	return count == 1
}

func main() {
	fmt.Println("day02")
	var passwordRegex = regexp.MustCompile(`(\d+)-(\d+) (\w)\: (\w+)`)
	var lines = readFile("input.txt")
	var validLines = 0
	var p *password
	for _, line := range lines {
		p = parseLine(line, passwordRegex)
		if passwordIsValid(p.min, p.max, p.letter, p.password) {
			validLines++
		}
	}

	fmt.Printf("part 1: %d\n", validLines)

	// Part 2
	var part2ValidLines = 0
	for _, line := range lines {
		p = parseLine(line, passwordRegex)
		if passwordIsValidPart2(p.min, p.max, p.letter, p.password) {
			part2ValidLines++
		}
	}

	fmt.Printf("part 2: %d\n", part2ValidLines)
}
