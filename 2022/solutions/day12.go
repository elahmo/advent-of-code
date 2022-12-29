package solutions

import (
	"aoc22/helpers"
	"fmt"
	"time"
)

const (
	width12  = 144
	height12 = 41
)

type GridPos struct {
	x int
	y int
}

func day12OneAndTwo() (int, int) {
	lines, _ := helpers.FileToLines("day12.txt")
	grid := make([]int, height12*width12)
	start, end := GridPos{}, GridPos{}
	startPositions := []GridPos{}
	//make grid
	for y, line := range lines {
		for x, char := range line {
			idx := y*width12 + x
			if string(char) == "S" {
				start.x = x
				start.y = y
				grid[idx] = 1
				startPositions = append(startPositions, GridPos{x, y})
				continue
			}
			if string(char) == "E" {
				end.x = x
				end.y = y
				grid[idx] = 26
				continue
			}
			height := charToHeight(char)
			grid[idx] = height
			if height == 1 {
				startPositions = append(startPositions, GridPos{x, y})
			}
		}
	}
	movesFromStart := breadthFirstSearch(grid, start, end)
	leastMoves := 99999999
	for _, startPos := range startPositions {
		moves := breadthFirstSearch(grid, startPos, end)
		if moves < leastMoves {
			leastMoves = moves
		}
	}
	return movesFromStart, leastMoves
}

type GridState struct {
	x         int
	y         int
	moveCount int
}

func breadthFirstSearch(grid []int, start GridPos, end GridPos) int {
	visited := map[int]bool{
		start.y*width12 + start.x: false,
	}
	queue := []GridState{
		{x: start.x, y: start.y, moveCount: 0},
	}
	for len(queue) > 0 {
		node := queue[0]
		x, y := node.x, node.y

		if end.x == x && end.y == y {
			return node.moveCount
		}
		queue = queue[1:]

		idx := y*width12 + x
		if !visited[idx] {
			visited[idx] = true

			// add all unvisited neighbours to the queue if possible
			// up
			if y > 0 && !visited[idx-width12] && grid[idx]+1 >= grid[idx-width12] {
				queue = append(queue, GridState{x: x, y: y - 1, moveCount: node.moveCount + 1})
			}
			// down
			if y < height12-1 && !visited[idx+width12] && grid[idx]+1 >= grid[idx+width12] {
				queue = append(queue, GridState{x: x, y: y + 1, moveCount: node.moveCount + 1})
			}
			// left
			if x > 0 && !visited[idx-1] && grid[idx]+1 >= grid[idx-1] {
				queue = append(queue, GridState{x: x - 1, y: y, moveCount: node.moveCount + 1})
			}
			// right
			if x < width12-1 && !visited[idx+1] && grid[idx]+1 >= grid[idx+1] {
				queue = append(queue, GridState{x: x + 1, y: y, moveCount: node.moveCount + 1})
			}
		}
	}
	return 9999
}

func charToHeight(c rune) int {
	runes := []rune{c}
	asciiValue := int(runes[0])
	return asciiValue - 96
}

func Day12() {
	start := time.Now()
	one, two := day12OneAndTwo()
	elapsed := time.Since(start)
	fmt.Printf("Day12, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
