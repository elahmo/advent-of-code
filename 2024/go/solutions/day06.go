package solutions

import (
	"aoc24/helpers"
	"fmt"
	"time"
)

func Day06One() int {
	lines, _ := helpers.FileToLines("day06.txt")
	guardX, guardY := 0, 0
	obstacles := make(map[[2]int]bool)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				obstacles[[2]int{x, y}] = true
			} else if char == '^' {
				guardX, guardY = x, y
			}
		}
	}
	guardPositions := moveGuard(guardX, guardY, obstacles)
	return len(guardPositions)
}

func moveGuard(x, y int, obstacles map[[2]int]bool) map[[2]int]bool {
	positions := make(map[[2]int]bool)
	directions := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	dirIndex := 0
	maxX, maxY := 130, 130 // Assuming a 130x130 grid, adjust as needed

	for {
		newX, newY := x+directions[dirIndex][0], y+directions[dirIndex][1]
		if newX < 0 || newY < 0 || newX >= maxX || newY >= maxY {
			break // Exit if out of bounds
		}
		if obstacles[[2]int{newX, newY}] {
			dirIndex = (dirIndex + 1) % 4 // Turn clockwise
		} else {
			x, y = newX, newY
			positions[[2]int{x, y}] = true
		}
	}
	// printout
	// for y := 0; y < maxY; y++ {
	// 	for x := 0; x < maxX; x++ {
	// 		if obstacles[[2]int{x, y}] {
	// 			fmt.Print("#")
	// 		} else if positions[[2]int{x, y}] {
	// 			fmt.Print("X")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	return positions
}

func Day06Two() int {
	return 0
}

func Day06() {
	start := time.Now()
	one := Day06One()
	two := Day06Two()
	elapsed := time.Since(start)
	fmt.Printf("Day06, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day06", Day06)
}
