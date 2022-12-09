package solutions

import (
	"aoc22/helpers"
	"fmt"
	"time"
)

func day08One() int {
	lines, _ := helpers.FileToLines("day08.txt")
	width := 99
	grid := make([]int, width*width)
	visible := 0
	for y, line := range lines {
		for x, h := range line {
			grid[y*width+x] = helpers.StrToInt(string(h))
		}
	}
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			if isVisible(x, y, grid, width) {
				visible++
			}
		}
	}
	return visible
}

func isVisible(x int, y int, grid []int, width int) bool {
	//store visibility
	vL, vR, vT, vB := true, true, true, true
	h := grid[y*width+x]
	if y == 0 || y == width-1 {
		return true
	}
	if x == 0 || x == width-1 {
		return true
	}
	// check left and right, x
	for i := 0; i < width; i++ {
		if i == x {
			continue
		}
		if grid[y*width+i] >= h {
			if i < x {
				vL = false
				continue
			}
			vR = false
		}
	}
	// check top and bottom, y
	for i := 0; i < width; i++ {
		if i == y {
			continue
		}
		if grid[i*width+x] >= h {
			if i < y {
				vT = false
				continue
			}
			vB = false
		}
	}
	return vL || vR || vT || vB
}

func day08Two() int {
	lines, _ := helpers.FileToLines("day08.txt")
	width := 99
	grid := make([]int, width*width)
	for y, line := range lines {
		for x, h := range line {
			grid[y*width+x] = helpers.StrToInt(string(h))
		}
	}
	scenic := 0
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			score := getScore(x, y, grid, width)
			if score > scenic {
				scenic = score
			}
		}
	}
	return scenic
}

func getScore(x int, y int, grid []int, width int) int {
	//store visibility
	scoreL, scoreR, scoreT, scoreB := 0, 0, 0, 0
	h := grid[y*width+x]
	score := 0

	// check left
	for i := x - 1; i >= 0; i-- {
		if x == 0 {
			break
		}
		if grid[y*width+i] >= h {
			score++
			break
		}
		score++
	}
	scoreL = score
	score = 0

	// check right
	for i := x + 1; i < width; i++ {
		if x == width-1 {
			break
		}
		if grid[y*width+i] >= h {
			score++
			break
		}
		score++
	}
	scoreR = score
	score = 0

	// check top
	for i := y - 1; i >= 0; i-- {
		if y == 0 {
			break
		}
		if grid[i*width+x] >= h {
			score++
			break
		}
		score++
	}
	scoreT = score
	score = 0

	// check bot
	for i := y + 1; i < width; i++ {
		if y == width-1 {
			break
		}
		if grid[i*width+x] >= h {
			score++
			break
		}
		score++
	}
	scoreB = score

	return scoreL * scoreR * scoreT * scoreB
}

func Day08() {
	start := time.Now()
	one := day08One()
	two := day08Two()
	elapsed := time.Since(start)
	fmt.Printf("Day08, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
