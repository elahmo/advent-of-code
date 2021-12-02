package main

import (
	"bufio"
	"fmt"
	"os"
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

type Hex struct {
	x int
	y int
}

func linesToGrid(lines []string) map[Hex]bool {
	grid := map[Hex]bool{}
	for i := -200; i < 200; i++ {
		for y := -200; y < 200; y++ {
			grid[Hex{i, y}] = false
		}
	}
	for _, line := range lines {
		countNW := strings.Count(line, "nw")
		countNE := strings.Count(line, "ne")
		countSW := strings.Count(line, "sw")
		countSE := strings.Count(line, "se")
		countE := strings.Count(line, "e") - countSE - countNE
		countW := strings.Count(line, "w") - countSW - countNW
		x := countNE + countE - countSW - countW
		y := countSW + countSE - countNW - countNE
		tile := Hex{x, y}
		grid[tile] = !grid[tile]
	}
	return grid
}

func updateNeighbours(grid map[Hex]bool) map[Hex]bool {
	var directions = map[string][]int{
		"nw": {0, -1},
		"ne": {1, -1},
		"sw": {-1, 1},
		"se": {0, 1},
		"e":  {1, 0},
		"w":  {-1, 0},
	}
	updatedGrid := map[Hex]bool{}
	for tile, isBlack := range grid {
		var blackCount int
		var black bool
		for _, xy := range directions {
			if grid[Hex{tile.x + xy[0], tile.y + xy[1]}] {
				blackCount++
			}
		}
		if isBlack {
			if blackCount == 1 || blackCount == 2 {
				black = true
			}
		} else {
			if blackCount == 2 {
				black = true
			}
		}
		updatedGrid[tile] = black
	}
	return updatedGrid
}

func part1() {
	lines, _ := fileToLines("input/day24.txt")
	gridLines := linesToGrid(lines)
	var counter int
	for _, value := range gridLines {
		if value {
			counter++
		}
	}
	fmt.Println(counter)
}

func part2() {
	lines, _ := fileToLines("input/day24.txt")
	gridLines := linesToGrid(lines)
	var counter int
	for i := 0; i < 100; i++ {
		gridLines = updateNeighbours(gridLines)
	}
	for _, value := range gridLines {
		if value {
			counter++
		}
	}
	fmt.Println(counter)
}

func main() {
	part2()
}
