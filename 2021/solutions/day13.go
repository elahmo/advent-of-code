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
	lines, _ := fileToLines("input/day13.txt")

	width, length, foldIdx := findMax(lines)

	grid := make([]bool, (length+1)*width)
	rex := regexp.MustCompile(`(\d+),(\d+)`)
	for i := 0; i < foldIdx-1; i++ {
		line := lines[i]
		data := rex.FindAllStringSubmatch(line, -1)
		x, _ := strconv.Atoi(data[0][1])
		y, _ := strconv.Atoi(data[0][2])
		grid[x+y*width] = true
	}

	grid = foldY(grid, width, length, 7)

	fmt.Println(length, width, foldIdx)
}

func foldY(grid []bool, w, l, foldLine int) []bool {
	newGrid := grid[:(l-foldLine)*w-1]
	distFromFold := 1
	for i := (l - foldLine + 1) * w; i < len(grid); i++ {
		newGrid[i-2*distFromFold*w] = grid[i]
		if (i-(l-foldLine)*w+1)%w == 0 {
			distFromFold++
		}
	}
	return newGrid
}

func findMax(lines []string) (int, int, int) {
	x := 0
	y := 0
	fold := 0
	rex := regexp.MustCompile(`(\d+),(\d+)`)
	for idx, line := range lines {
		data := rex.FindAllStringSubmatch(line, -1)
		if data == nil {
			fold = idx + 1
			break
		}
		tempx, _ := strconv.Atoi(data[0][1])
		tempy, _ := strconv.Atoi(data[0][2])
		if tempx > x {
			x = tempx
		}
		if tempy > y {
			y = tempy
		}
	}
	return x, y, fold
}

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
