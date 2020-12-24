package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var recordKeyValueRegex *regexp.Regexp
var heightRegex *regexp.Regexp
var hairColorRegex *regexp.Regexp
var eyeColorRegex *regexp.Regexp
var passportIDRegex *regexp.Regexp

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileRecords(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	checkError(err)

	lines := strings.Split(string(data), "\r\n\r\n")
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

func recordIsValidStrict(record map[string]string, requiredFields *[]string) bool {
	if !recordIsValid(record, requiredFields) {
		return false
	}

	for k, v := range record {
		switch k {
		case "byr":
			if !validBirthYear(v) {
				return false
			}
		case "iyr":
			if !validIssueYear(v) {
				return false
			}
		case "eyr":
			if !validExpirationYear(v) {
				return false
			}
		case "hgt":
			if !validHeight(v) {
				return false
			}
		case "hcl":
			if !validHairColor(v) {
				return false
			}
		case "ecl":
			if !validEyeColor(v) {
				return false
			}
		case "pid":
			if !validPassportID(v) {
				return false
			}
		}
	}

	return true
}

func validBirthYear(val string) bool {
	year, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return year >= 1920 && year <= 2002
}

func validIssueYear(val string) bool {
	year, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return year >= 2010 && year <= 2020
}

func validExpirationYear(val string) bool {
	year, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return year >= 2020 && year <= 2030
}

func validHeight(val string) bool {
	matches := heightRegex.FindAllSubmatch([]byte(val), -1)
	if len(matches) < 1 || len(matches[0]) < 2 {
		return false
	}

	height, _ := strconv.Atoi(string(matches[0][1]))
	if string(matches[0][2]) == "cm" {
		return height >= 150 && height <= 193
	} else {
		return height >= 59 && height <= 76
	}
}

func validHairColor(val string) bool {
	return hairColorRegex.Match([]byte(val))
}

func validEyeColor(val string) bool {
	return eyeColorRegex.Match([]byte(val))
}

func validPassportID(val string) bool {
	return passportIDRegex.Match([]byte(val))
}

func main() {
	fmt.Println("day04")

	recordKeyValueRegex = regexp.MustCompile(`(\S+)\:(\S+)`)
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validRecords := 0
	records := readFileRecords("input.txt")
	fmt.Printf("%d records\n", len(records))

	for _, record := range records {
		parseRecord := parseRecord(record)
		if recordIsValid(parseRecord, &requiredFields) {
			validRecords++
		}
	}

	fmt.Printf("part 1: %d\n", validRecords)

	// -------------------------------------
	// Part 2
	validRecordsPartTwo := 0
	heightRegex = regexp.MustCompile(`(\d+)(cm|in)`)
	hairColorRegex = regexp.MustCompile(`#[0-9a-f]{6}`)
	eyeColorRegex = regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	passportIDRegex = regexp.MustCompile(`^\d{9}$`)

	for _, record := range records {
		parseRecord := parseRecord(record)
		if recordIsValidStrict(parseRecord, &requiredFields) {
			validRecordsPartTwo++
			delete(parseRecord, "cid")
			fmt.Printf("%s\n", parseRecord)
		}
	}

	fmt.Printf("part 2: %d\n", validRecordsPartTwo)

}
