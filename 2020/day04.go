package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func fileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func linesToInt(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for _, w := range lines {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func linesToPassports(lines []string) map[int]map[string]string {
	var rex = regexp.MustCompile("(\\w+):(\\W*\\w+)")
	var passports = map[int]map[string]string{}
	currentPass := 0
	for _, w := range lines {
		_, ok := passports[currentPass]
		if !ok {
			passports[currentPass] = map[string]string{}
		}
		if w != "" {
			data := rex.FindAllStringSubmatch(w, -1)
			for _, kv := range data {
				passports[currentPass][kv[1]] = kv[2]
			}
		} else {
			currentPass++
		}
	}
	return passports
}

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func ValidatePassport(pass map[string]string) bool {
	success := true
	for key, val := range pass {
		switch key {
		case "byr":
			matched, _ := regexp.MatchString(`^(19[2-8][0-9]|199[0-9]|200[0-2])$`, val)
			if !matched {
				success = false
			}
		case "iyr":
			matched, _ := regexp.MatchString(`^(201[0-9]|2020)$`, val)
			if !matched {
				success = false
			}
		case "eyr":
			matched, _ := regexp.MatchString(`^(202[0-9]|2030)$`, val)
			if !matched {
				success = false
			}
		case "hgt":
			matched, _ := regexp.MatchString(`^(1[5-8][0-9]|19[0-3])cm|^(59|6[0-9]|7[0-6])in$`, val)
			if !matched {
				success = false
			}
		case "hcl":
			matched, _ := regexp.MatchString(`^#\w{6}$`, val)
			if !matched {
				success = false
			}
		case "ecl":
			matched, _ := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, val)
			if !matched {
				success = false
			}
		case "pid":
			matched, _ := regexp.MatchString(`^\d{9}$`, val)
			if !matched {
				success = false
			}
		case "cid":
		default:
			success = false
		}
	}
	return success
}

func parts1and2() {
	lines, _ := fileToLines("input/day04.txt")
	passData := linesToPassports(lines)
	requiredData := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	matchesPartOne := 0
	matchesPartTwo := 0
	for _, pass := range passData {
		fields := make([]string, 0, len(pass))
		for k := range pass {
			fields = append(fields, k)
		}
		fieldsMatching := Intersection(requiredData, fields)
		if len(fieldsMatching) == 7 {
			matchesPartOne++
			if ValidatePassport(pass) {
				matchesPartTwo++
			}
		}
	}
	fmt.Println(matchesPartOne, matchesPartTwo)
	return
}

func main() {
	parts1and2()
}
