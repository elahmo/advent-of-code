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

func part1() {
	lines, _ := fileToLines("input/day06.txt")

	fish := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	rex := regexp.MustCompile(`\d+`)
	data := rex.FindAllStringSubmatch(lines[0], -1)
	for _, num := range data {
		parsedNum, _ := strconv.Atoi(num[0])
		fish[parsedNum]++
	}
	fmt.Println(fish)

	updatedFish := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for i := 0; i <= 80; i++ {
		for key, value := range fish {
			if value > 0 {
				if key == 0 {
					updatedFish[8] += value
					updatedFish[6] += value
				} else {
					updatedFish[key-1] += value
				}
			}
		}
		fish = updatedFish
		updatedFish = map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	}
	sum := 0
	for i := 0; i < 8; i++ {
		sum += fish[i]
	}
	fmt.Println(sum)
}

func printFish(fish map[int]int) {
	for i := 0; i < 9; i++ {
		fmt.Print(fish[i], ",")
	}
	fmt.Println()
}

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
