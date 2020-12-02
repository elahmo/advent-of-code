package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1() {
	lines, _ := fileToLines("input/day01.txt")
	numbers := linesToInt(lines)

	differences := make([]int, 0, len(lines))
	for _, num := range numbers {
		diff := 2020 - num
		differences = append(differences, diff)
	}

	m := make(map[int]int)
	for _, k := range numbers {
		m[k]++
	}
	for _, k := range differences {
		m[k]++
	}

	for k, v := range m {
		if v == 2 {
			fmt.Println(k, 2020-k, k*(2020-k))
			break
		}
	}
}

func part2() {
	lines, _ := fileToLines("input/day01.txt")
	numbers := linesToInt(lines)

	for _, number := range numbers {
		// seemd a bit crude, but do the same as in previous
		differences := make([]int, 0, len(lines))
		for _, num := range numbers {
			diff := 2020 - num - number
			differences = append(differences, diff)
		}

		m := make(map[int]int)
		for _, k := range numbers {
			m[k]++
		}
		for _, k := range differences {
			m[k]++
		}

		found := false
		for k, v := range m {
			if v == 2 {
				fmt.Println(k, number, 2020-k-number, k*number*(2020-k-number))
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func main() {
	part1()
	part2()
}
