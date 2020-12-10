package main

import "log"
import "fmt"
import "strings"
import "regexp"
import "strconv"

func day4_1(lines []string) int {
	return solve(lines, validatePassport)
}

func day4_2(lines []string) int {
	return solve(lines, validatePassportV2)
}

func solve(lines []string, validator func(map[string]string) bool) int {
	count := 0

	passports := readRecords(lines)
	for _, passport := range passports {
		if validator(passport) {
			count++
		}
	}

	return count
}

func populateFields(passport map[string]string, line string) {
	fields := strings.Fields(line)
	for _, field := range fields {
		keyValueArray := strings.Split(field, ":")
		passport[keyValueArray[0]] = keyValueArray[1]
	}
}

func readRecords(lines []string) []map[string]string {
	var passports []map[string]string
	currentPass := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			passports = append(passports, currentPass)
			currentPass = make(map[string]string)
		}
		populateFields(currentPass, line)
	}
	passports = append(passports, currentPass)
	return passports
}

func validatePassport(passport map[string]string) bool {
	fields := mandatoryFields()
	for _, field := range fields {
		_, ok := passport[field]
		if !ok {
			return false
		}
	}
	return true
}

func validatePassportV2(passport map[string]string) bool {
	valid := true
	for field, validator := range fieldValidators() {
		value, ok := passport[field]
		valid = valid && ok && validator(value)
	}
	return valid
}

// Validator - function literal definition for passport field validation
type Validator func(string) bool

func validateRange(field string, min int, max int) bool {
	f, _ := strconv.Atoi(field)
	return f >= min && f <= max
}

func fieldValidators() map[string]Validator {
	return map[string]Validator{
		"byr": func(field string) bool {
			return validateRange(field, 1920, 2002)
		},
		"iyr": func(field string) bool {
			return validateRange(field, 2010, 2020)
		},
		"eyr": func(field string) bool {
			return validateRange(field, 2020, 2030)
		},
		"hgt": func(field string) bool {
			if strings.HasSuffix(field, "cm") {
				value := field[0:(len(field) - 2)]
				return validateRange(value, 150, 193)
			}
			if strings.HasSuffix(field, "in") {
				value := field[0:(len(field) - 2)]
				return validateRange(value, 59, 76)
			}
			return false
		},
		"hcl": func(field string) bool {
			re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			return re.MatchString(field)
		},
		"ecl": func(field string) bool {
			switch field {
			case
				"amb",
				"blu",
				"brn",
				"gry",
				"hzl",
				"grn",
				"oth":
				return true
			}
			return false
		},
		"pid": func(field string) bool {
			re := regexp.MustCompile(`^[0-9]{9}$`)
			return re.MatchString(field)
		},
		// "cid"
	}
}

func mandatoryFields() []string {
	return []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid"
	}
}

func day4Main() {
	lines, err := ReadLines("../inputs/day4")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	count := day4_1(lines)
	fmt.Println(count)

	count2 := day4_2(lines)
	fmt.Println(count2)
}
