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

func linesToKeys(lines []string) (card int, door int) {
	card, _ = strconv.Atoi(lines[0])
	door, _ = strconv.Atoi(lines[1])
	return
}

func part1() {
	lines, _ := fileToLines("input/day25.txt")
	card, door := linesToKeys(lines)
	value := 1
	subject := 7
	loopSizeCard := 0
	for {
		value *= subject
		value = value % 20201227
		loopSizeCard++
		if value == card {
			break
		}
	}
	value = 1
	loopSizeDoor := 0
	for {
		value *= subject
		value = value % 20201227
		loopSizeDoor++
		if value == door {
			break
		}
	}
	value = 1
	for i := 0; i < loopSizeDoor; i++ {
		value *= card
		value = value % 20201227
	}
	fmt.Println(value)
}

func main() {
	part1()
	// no part 2 :()
}
