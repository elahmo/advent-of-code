package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func checkSums(numbers []int, sum int) (addsUp bool) {
	for _, numOne := range numbers {
		for _, numTwo := range numbers {
			if numOne == numTwo {
				continue
			}
			if numOne+numTwo == sum {
				addsUp = true
				break
			}
		}
		if addsUp {
			break
		}
	}
	return addsUp
}

func part1() {
	lines, _ := fileToLines("input/day10.txt")
	numbers := linesToInt(lines)
	sort.Ints(numbers)
	differenceMap := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
	for idx, num := range numbers[1:] {
		differenceMap[num-numbers[idx]]++
	}
	differenceMap[numbers[0]]++
	differenceMap[3]++
	fmt.Println(differenceMap[1] * differenceMap[3])
}

func part2() {
	lines, _ := fileToLines("input/day10.txt")
	numbers := linesToInt(lines)
	sort.Ints(numbers)
	numbers = append(numbers[:1], numbers[0:]...)
	numbers[0] = 0
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	arrangements := map[int]int{}
	arrangements[0] = 1
	for _, number := range numbers {
		arrangements[number+1] += arrangements[number]
		arrangements[number+2] += arrangements[number]
		arrangements[number+3] += arrangements[number]
	}

	fmt.Println(arrangements[numbers[len(numbers)-1]])
}

func main() {
	part1()
	part2()
}
