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

type Instruction struct {
	name  string
	steps int
}

func linesToInstructions(lines []string) map[int]Instruction {
	instructions := make(map[int]Instruction)
	var rex = regexp.MustCompile("(\\w+) (.\\d+)")
	for id, line := range lines {
		data := rex.FindAllStringSubmatch(line, -1)
		inst := data[0][1]
		steps, _ := strconv.Atoi(data[0][2])
		instructions[id] = Instruction{inst, steps}
	}
	return instructions
}

func part1() {
	lines, _ := fileToLines("input/day08.txt")
	// parse the lines to instructions
	instructions := linesToInstructions(lines)
	// hold which is visited
	visited := make([]bool, len(instructions))
	// hold accumulator value
	acc := 0
	i := 0
	// iterate through instructions, with i
	for {
		// check if visited, if so break and return, else mark as visited
		if visited[i] {
			break
		} else {
			visited[i] = true
		}
		// for each, if jump go to index, else follow properly
		ins := instructions[i]
		if ins.name == "acc" {
			acc += ins.steps
			i++
		} else if ins.name == "jmp" {
			i += ins.steps
		} else {
			i++
		}
	}
	fmt.Println(acc)
	return
}

func part2() {
	lines, _ := fileToLines("input/day08.txt")
	instructions := linesToInstructions(lines)
	completed := false
	acc := 0

	for key, value := range instructions {
		visited := make([]bool, len(instructions))
		acc = 0

		// change one instruction at a time
		modifiedinstructions := make(map[int]Instruction)
		if value.name == "nop" || value.name == "jmp" {
			for k, v := range instructions {
				modifiedinstructions[k] = v
			}
			if value.name == "nop" {
				modifiedinstructions[key] = Instruction{
					"jmp",
					value.steps,
				}
			} else {
				modifiedinstructions[key] = Instruction{
					"nop",
					value.steps,
				}
			}
		} else {
			continue
		}

		// similar block as in part1, using modified instructions
		i := 0
		for {
			if i > len(modifiedinstructions)-1 {
				completed = true
				break
			}
			if visited[i] {
				break
			} else {
				visited[i] = true
			}
			ins := modifiedinstructions[i]
			if ins.name == "acc" {
				acc += ins.steps
				i++
			} else if ins.name == "jmp" {
				i += ins.steps
			} else {
				i++
			}
		}
		if completed {
			break
		}
	}
	fmt.Println(acc)
	return
}

func main() {
	part2()
}
