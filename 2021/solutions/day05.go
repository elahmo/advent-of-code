package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

// func part1() {
// 	lines, _ := fileToLines("input/day05.txt")

// 	grid := make([]int, 999*999)
// 	rex := regexp.MustCompile(`\d+`)
// 	for _, line := range lines {
// 		data := rex.FindAllStringSubmatch(line, -1)
// 		x1, _ := strconv.Atoi(data[0][0])
// 		y1, _ := strconv.Atoi(data[1][0])
// 		x2, _ := strconv.Atoi(data[2][0])
// 		y2, _ := strconv.Atoi(data[3][0])
// 		grid = updateGrid(grid, x1, y1, x2, y2)
// 	}
// 	// printGrid(grid)
// 	calculateScore(grid)
// }

func part2() {
	lines, _ := fileToLines("input/day05.txt")

	grid := make([]int, 999*999)
	rex := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		data := rex.FindAllStringSubmatch(line, -1)
		x1, _ := strconv.Atoi(data[0][0])
		y1, _ := strconv.Atoi(data[1][0])
		x2, _ := strconv.Atoi(data[2][0])
		y2, _ := strconv.Atoi(data[3][0])
		grid = updateGridDiagonal(grid, x1, y1, x2, y2)
	}
	// printGrid(grid)
	calculateScore(grid)
}

func printGrid(grid []int) {
	limit := 999
	fmt.Println("Grid:")
	for idx, item := range grid {
		if idx%limit == 0 {
			fmt.Println("")
		}
		fmt.Print(item)
	}
	fmt.Println("\n=========")
}

func calculateScore(grid []int) {
	score := 0
	for _, num := range grid {
		if num >= 2 {
			score++
		}
	}
	fmt.Println(score)
}

func updateGrid(grid []int, coords ...int) []int {
	limit := 999
	x1, y1, x2, y2 := coords[0], coords[1], coords[2], coords[3]
	if x1 == x2 {
		// row update
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for i := y1; i <= y2; i++ {
			grid[i*limit+x1]++
		}
	} else if y1 == y2 {
		// column update
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for i := x1; i <= x2; i++ {
			grid[i+limit*y1]++
		}
	}
	return grid
}

func updateGridDiagonal(grid []int, coords ...int) []int {
	limit := 999
	x1, y1, x2, y2 := coords[0], coords[1], coords[2], coords[3]
	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for i := y1; i <= y2; i++ {
			grid[i*limit+x1]++
		}
	} else if y1 == y2 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for i := x1; i <= x2; i++ {
			grid[i+limit*y1]++
		}
	} else {
		xPos, yPos := 1, 1
		rng := int(math.Abs(float64(x1 - x2)))
		if x1 > x2 {
			xPos = -1
		}
		if y1 > y2 {
			yPos = -1
		}
		for i := 0; i <= rng; i++ {
			grid[(y1+i*yPos)*limit+(x1+(i*xPos))]++
		}
	}
	return grid
}

func main() {
	start := time.Now()
	// part1()
	part2()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
