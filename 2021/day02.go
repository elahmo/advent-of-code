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

type Instruction struct {
	name  string
	steps int
}

func linesToInstructions(lines []string) (instructions []Instruction) {
	var rex = regexp.MustCompile("(\\w+) (\\d+)")
	for _, line := range lines {
		data := rex.FindAllStringSubmatch(line, -1)
		inst := data[0][1]
		steps, _ := strconv.Atoi(data[0][2])
		instructions = append(instructions, Instruction{inst, steps})
	}
	return
}

func part1() {
	lines, _ := fileToLines("input/day02.txt")
	instructions := linesToInstructions(lines)

	position := 0
	depth := 0
	for _, inst := range instructions {
		if inst.name == "forward" {
			position += inst.steps
		} else if inst.name == "down" {
			depth += inst.steps
		} else {
			depth -= inst.steps
		}
	}
	fmt.Println(position * depth)
}

func part2() {
	lines, _ := fileToLines("input/day02.txt")
	instructions := linesToInstructions(lines)

	position := 0
	depth := 0
	aim := 0
	for _, inst := range instructions {
		if inst.name == "forward" {
			position += inst.steps
			depth += inst.steps * aim
		} else if inst.name == "down" {
			aim += inst.steps
		} else {
			aim -= inst.steps
		}
	}
	fmt.Println(position * depth)
}

func main() {
	part1()
	part2()
}
