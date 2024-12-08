package solutions

import (
	"aoc24/helpers"
	"fmt"
	"time"
)

func Day04One() int {
	lines, _ := helpers.FileToLines("day04.txt")
	rows := len(lines)
	cols := len(lines[0])
	var grid []string
	var start_chars []int
	// find all locations of letter X
	for idy, line := range lines {
		for idx, subline := range line {
			grid = append(grid, string(subline))
			if subline == 'X' {
				start_chars = append(start_chars, idx+cols*idy)
			}
		}
	}
	count := 0
	for _, location := range start_chars {
		x, y := toCoordinates(location, cols)
		visited := make(map[int]bool)
		found_words := 0
		for _, dir := range directions {
			search(grid, x, y, rows, cols, 1, dir, visited, &found_words)
		}
		count += found_words
	}
	return count
}

// Depth-limited search for spelling "XMAS" on the grid in a single direction
func search(grid []string, x, y, rows, cols, depth int, dir [2]int, visited map[int]bool, counter *int) {
	// Define the word to spell
	targetWord := "XMAS"

	// Stop recursion if we go beyond the word length
	if depth > len(targetWord) {
		return
	}

	// Check if the current letter matches the target at this depth
	targetValue := string(targetWord[depth-1])
	index := toIndex(x, y, cols)

	if grid[index] != targetValue {
		return
	}

	// Mark the current position as visited
	visited[index] = true

	// If we reach the last letter, increment the counter
	if depth == len(targetWord) {
		*counter++
		visited[index] = false // Unmark for backtracking
		return
	}

	// Move in the current direction
	newX, newY := x+dir[0], y+dir[1]
	if isValid(newX, newY, rows, cols) && !visited[toIndex(newX, newY, cols)] {
		search(grid, newX, newY, rows, cols, depth+1, dir, visited, counter)
	}

	// Unmark the current position for backtracking
	visited[index] = false
}

func toCoordinates(index, cols int) (int, int) {
	x := index / cols // Row
	y := index % cols // Column
	return x, y
}

var directions = [][2]int{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Top-left
	{-1, 1},  // Top-right
	{1, -1},  // Bottom-left
	{1, 1},   // Bottom-right
}

// Convert 2D coordinates to a 1D index
func toIndex(x, y, cols int) int {
	return x*cols + y
}

// Check if a position is valid
func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func Day04Two() int {
	lines, _ := helpers.FileToLines("day04.txt")
	rows := len(lines)
	cols := len(lines[0])
	var grid []string
	var start_chars []int
	// find all locations of letter X
	for idy, line := range lines {
		for idx, subline := range line {
			grid = append(grid, string(subline))
			if subline == 'A' {
				start_chars = append(start_chars, idx+cols*idy)
			}
		}
	}
	count := 0
	for _, location := range start_chars {
		x, y := toCoordinates(location, cols)
		// ignore first and last row, and first and last col
		if x == 0 || y == 0 || x == cols-1 || y == rows-1 {
			continue
		}
		if check_xmas(x, y, cols, grid) {
			count += 1
		}
	}
	return count
}

func check_xmas(x, y, cols int, grid []string) bool {
	var left, right bool
	if (grid[toIndex(x-1, y-1, cols)] == "M" && grid[toIndex(x+1, y+1, cols)] == "S") ||
		(grid[toIndex(x-1, y-1, cols)] == "S" && grid[toIndex(x+1, y+1, cols)] == "M") {
		left = true
	}

	if (grid[toIndex(x-1, y+1, cols)] == "M" && grid[toIndex(x+1, y-1, cols)] == "S") ||
		(grid[toIndex(x-1, y+1, cols)] == "S" && grid[toIndex(x+1, y-1, cols)] == "M") {
		right = true
	}

	return (left && right)
}

func Day04() {
	start := time.Now()
	one := Day04One()
	two := Day04Two()
	elapsed := time.Since(start)
	fmt.Printf("Day04, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day04", Day04)
}
