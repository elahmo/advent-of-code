package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	lines, _ := fileToLines("input/day07.txt")

	fuelCost := 99999999
	rng := 0
	var numbers []int
	rex := regexp.MustCompile(`\d+`)
	data := rex.FindAllStringSubmatch(lines[0], -1)
	for _, num := range data {
		parsedNum, _ := strconv.Atoi(num[0])
		numbers = append(numbers, parsedNum)
		if parsedNum > rng {
			rng = parsedNum
		}
	}

	for i := 0; i <= rng; i++ {
		cost := 0
		for _, num := range numbers {
			cost += Abs(i - num)
		}
		if cost < fuelCost {
			fuelCost = cost
		}
	}
	fmt.Println(fuelCost)
}

func Mul(x int) int {
	result := 0
	for i := 1; i <= x; i++ {
		result += i
	}
	return result
}

func part2() {
	lines, _ := fileToLines("input/day07.txt")

	fuelCost := 99999999
	rng := 0
	var numbers []int
	rex := regexp.MustCompile(`\d+`)
	data := rex.FindAllStringSubmatch(lines[0], -1)
	for _, num := range data {
		parsedNum, _ := strconv.Atoi(num[0])
		numbers = append(numbers, parsedNum)
		if parsedNum > rng {
			rng = parsedNum
		}
	}

	for i := 0; i <= rng; i++ {
		cost := 0
		for _, num := range numbers {
			cost += Mul(Abs(i - num))
		}
		if cost < fuelCost {
			fuelCost = cost
		}
	}
	fmt.Println(fuelCost)
}

func main() {
	start := time.Now()
	part2()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
