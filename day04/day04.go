package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var recordKeyValueRegex *regexp.Regexp

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileRecords(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	checkError(err)

	lines := strings.Split(string(data), "\n\n")
	return lines
}

func parseRecord(recordString string) map[string]string {
	record := make(map[string]string, 0)

	pairs := recordKeyValueRegex.FindAllSubmatch([]byte(recordString), -1)
	for _, pair := range pairs {
		record[string(pair[1])] = string(pair[2])
	}

	return record
}

func recordIsValid(record map[string]string, requiredFields *[]string) bool {
	for _, requiredField := range *requiredFields {
		_, containsKey := record[requiredField]
		if !containsKey {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("day04")

	recordKeyValueRegex = regexp.MustCompile(`(\S+)\:(\S+)`)
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validRecords := 0
	records := readFileRecords("input.txt")
	// parseRecord := nil
	for _, record := range records {
		parseRecord := parseRecord(record)
		if recordIsValid(parseRecord, &requiredFields) {
			validRecords++
		}
	}

	fmt.Printf("part 1: %d\n", validRecords)
}
