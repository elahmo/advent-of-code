package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func linesToBags(lines []string) map[string]map[string]int {
	var rexParentBag = regexp.MustCompile(`^(\w+ \w+)`)
	var rexBags = regexp.MustCompile(` (\d) (\w+ \w+) bags?\W`)
	var bagMap = map[string]map[string]int{}
	for _, w := range lines {
		parentBag := rexParentBag.FindStringSubmatch(w)[0]
		_, ok := bagMap[parentBag]
		if !ok {
			bagMap[parentBag] = map[string]int{}
		}
		if strings.Contains(w, "contain no other bags") {
		} else {
			bags := rexBags.FindAllStringSubmatch(w, -1)
			for _, bag := range bags {
				i, _ := strconv.Atoi(bag[1])
				bagMap[parentBag][bag[2]] += i
			}
		}
	}
	return bagMap
}

func iterateThroughBag(bag string, bagData map[string]map[string]int) []string {
	validChildrenBags := []string{}
	bags := bagData[bag]
	for childBag := range bags {
		if childBag == "shiny gold" {
			validChildrenBags = append(validChildrenBags, bag)
		} else {
			bags := iterateThroughBag(childBag, bagData)
			for _, bag := range bags {
				validChildrenBags = append(validChildrenBags, bag)
			}
		}
	}
	return validChildrenBags
}

func part1() {
	lines, _ := fileToLines("input/day07.txt")
	targetBag := "shiny gold"
	bagData := linesToBags(lines)
	validBags := map[string]bool{}
	for bag := range bagData {
		validBags[bag] = false
	}
	for parentBag, childBags := range bagData {
		for childBag := range childBags {
			if childBag == targetBag {
				validBags[parentBag] = true
			} else {
				bags := iterateThroughBag(childBag, bagData)
				if len(bags) > 0 {
					validBags[parentBag] = true
				}
				for _, bag := range bags {
					validBags[bag] = true
				}
			}
		}
	}
	counter := 0
	for _, valid := range validBags {
		if valid {
			counter++
		}
	}
	fmt.Println(counter)
	return
}

func totalForChildBag(bag string, bagData map[string]map[string]int) int {
	total := 0
	bags := bagData[bag]
	for childBag, childQty := range bags {
		total += childQty
		total += childQty * totalForChildBag(childBag, bagData)
	}
	return total
}

func part2() {
	lines, _ := fileToLines("input/day07.txt")
	targetBag := "shiny gold"
	bagData := linesToBags(lines)
	counter := 0
	for childBag, qty := range bagData[targetBag] {
		bagAmount := qty * totalForChildBag(childBag, bagData)
		counter += qty + bagAmount
	}
	fmt.Println(counter)
	return
}

func main() {
	part1()
	part2()
}
