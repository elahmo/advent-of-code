package solutions

import (
	"aoc23/helpers"
	"fmt"
	"time"
)

type GridPos struct {
	x int
	y int
}

type GridState struct {
	x        int
	y        int
	lastX    int
	lastY    int
	steps    int
	heatLoss int
}

type Move struct {
	x   int
	y   int
	dir string
}

func day17One() int {
	lines, _ := helpers.FileToLines("day17.txt")
	LEN := len(lines)
	grid := make([]int, LEN*LEN)
	for y, line := range lines {
		for x, char := range line {
			idx := y*LEN + x
			grid[idx] = helpers.StrToInt(string(char))
		}
	}

	start, end := GridPos{0, 0}, GridPos{LEN - 1, LEN - 1}
	movesFromStart := aStarSearch(grid, start, end, LEN)
	return movesFromStart
}

func breadthFirstSearch(grid []int, start GridPos, end GridPos, LEN int) int {
	queue := []GridState{
		{x: start.x, y: start.y, steps: 0, heatLoss: 0},
	}
	minCost := 9999
	for len(queue) > 0 {
		node := queue[0]
		x, y := node.x, node.y
		queue = queue[1:]
		if end.x == x && end.y == y {
			if node.heatLoss < minCost {
				minCost = node.heatLoss
			}
			continue
		}

		idx := y*LEN + x
		// up
		if y > 0 && canMove("up", node) {
			stepCost := calculateStepCost("up", node)
			// fmt.Printf("going up from %d,%d, prev %d,%d, heatloss: %d, stepCost: %d\n", node.x, node.y, node.lastX, node.lastY, node.heatLoss, stepCost)
			newNodeHeatLoss := grid[idx-LEN]
			queue = append(queue, GridState{x: x, y: y - 1, lastX: x, lastY: y, steps: stepCost, heatLoss: node.heatLoss + newNodeHeatLoss})
		}
		// down
		if y < LEN-1 && canMove("down", node) {

			stepCost := calculateStepCost("down", node)
			// fmt.Printf("going down from %d,%d, prev %d,%d, heatloss: %d, stepCost: %d\n", node.x, node.y, node.lastX, node.lastY, node.heatLoss, stepCost)
			newNodeHeatLoss := grid[idx+LEN]
			queue = append(queue, GridState{x: x, y: y + 1, lastX: x, lastY: y, steps: stepCost, heatLoss: node.heatLoss + newNodeHeatLoss})
		}
		// left
		if x > 0 && canMove("left", node) {

			stepCost := calculateStepCost("left", node)
			// fmt.Printf("going left from %d,%d, prev %d,%d, heatloss: %d, stepCost: %d\n", node.x, node.y, node.lastX, node.lastY, node.heatLoss, stepCost)
			newNodeHeatLoss := grid[idx-1]
			queue = append(queue, GridState{x: x - 1, y: y, lastX: x, lastY: y, steps: stepCost, heatLoss: node.heatLoss + newNodeHeatLoss})
		}
		// right
		if x < LEN-1 && canMove("right", node) {

			stepCost := calculateStepCost("right", node)
			// fmt.Printf("going right from %d,%d, prev %d,%d, heatloss: %d, stepCost: %d\n", node.x, node.y, node.lastX, node.lastY, node.heatLoss, stepCost)
			newNodeHeatLoss := grid[idx+1]
			queue = append(queue, GridState{x: x + 1, y: y, lastX: x, lastY: y, steps: stepCost, heatLoss: node.heatLoss + newNodeHeatLoss})
		}
	}
	return minCost
}

type Node struct {
	x, y    int
	g, h, f int
	parent  *Node
}

func aStarSearch(grid []int, start, end GridPos, LEN int) int {
	openSet := make(map[GridPos]*Node)
	closedSet := make(map[GridPos]bool)

	startNode := &Node{x: start.x, y: start.y}
	startNode.g = 0
	startNode.h = 0
	startNode.f = startNode.g + startNode.h

	openSet[start] = startNode

	for len(openSet) > 0 {
		// Find the node with the lowest f value in openSet
		current := findLowestF(openSet)

		// If the current node is the goal, return the cost
		if current.x == end.x && current.y == end.y {
			return current.g
		}

		// Move current from openSet to closedSet
		delete(openSet, GridPos{current.x, current.y})
		closedSet[GridPos{current.x, current.y}] = true

		// Generate and consider neighbors
		neighbors := getNeighbors(current, grid, LEN)
		for _, neighbor := range neighbors {
			if closedSet[GridPos{neighbor.x, neighbor.y}] {
				continue // Skip already processed nodes
			}

			// Calculate tentative g value
			tentativeG := current.g + neighbor.steps

			// If the neighbor is not in openSet or has a lower g value
			if node, ok := openSet[GridPos{neighbor.x, neighbor.y}]; !ok || tentativeG < node.g {
				if !ok {
					node = &Node{x: neighbor.x, y: neighbor.y}
					openSet[GridPos{neighbor.x, neighbor.y}] = node
				}

				// Update node values
				node.parent = current
				node.g = tentativeG
				node.h = heuristic(neighbor.x, neighbor.y, end.x, end.y, grid, LEN)
				node.f = node.g + node.h
			}
		}
	}

	// If the search fails, return some default value
	return -1
}

func heuristic(x1, y1, x2, y2 int, grid []int, LEN int) int {
	// Use heat loss as the heuristic
	idx1 := y1*LEN + x1
	idx2 := y2*LEN + x2
	return abs(grid[idx1] - grid[idx2])
}

func getNeighbors(node *Node, grid []int, LEN int) []GridState {
	neighbors := []GridState{}

	// Define possible directions
	directions := []struct{ dx, dy int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for _, dir := range directions {
		newX, newY := node.x+dir.dx, node.y+dir.dy

		// Check if the new position is within bounds
		if newX >= 0 && newX < LEN && newY >= 0 && newY < LEN {
			// Check if the move is valid based on your canMove function
			if canMove(dirToString(dir.dx, dir.dy), GridState{
				x:        newX,
				y:        newY,
				lastX:    node.x,
				lastY:    node.y,
				steps:    calculateStepCost(dirToString(dir.dx, dir.dy), *node),
				heatLoss: grid[newY*LEN+newX],
				moveDb:   copyMoveMap(node.moveDb),
			}) {
				// Calculate the new position's index in the grid
				idx := newY*LEN + newX

				// Create a new neighbor node
				neighbor := GridState{
					x:        newX,
					y:        newY,
					lastX:    node.x,
					lastY:    node.y,
					steps:    calculateStepCost(dirToString(dir.dx, dir.dy), *node),
					heatLoss: grid[idx],
					moveDb:   copyMoveMap(node.moveDb),
				}

				// Add the neighbor to the list
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func dirToString(dx, dy int) string {
	// Convert direction to a string representation
	switch {
	case dx == 0 && dy == -1:
		return "up"
	case dx == 0 && dy == 1:
		return "down"
	case dx == -1 && dy == 0:
		return "left"
	case dx == 1 && dy == 0:
		return "right"
	default:
		return "unknown"
	}
}

func findLowestF(openSet map[GridPos]*Node) *Node {
	var lowestNode *Node
	var lowestF int

	first := true
	for _, node := range openSet {
		if first || node.f < lowestF {
			lowestNode = node
			lowestF = node.f
			first = false
		}
	}

	return lowestNode
}

// helper that stops movement for more than three steps in the same direction
func canMove(dir string, node GridState) bool {
	switch dir {
	case "up":
		// prevent going in a straight line for more than 3 steps
		if node.lastY-1 == node.y && node.x == node.lastX && node.steps > 2 {
			return false
		}
		// prevent going backwards
		if node.lastY+1 == node.y && node.x == node.lastX {
			return false
		}
	case "down":
		// prevent going in a straight line for more than 3 steps
		if node.lastY+1 == node.y && node.x == node.lastX && node.steps > 2 {
			return false
		}
		// prevent going backwards
		if node.lastY-1 == node.y && node.x == node.lastX {
			return false
		}

	case "left":
		// prevent going in a straight line for more than 3 steps
		if node.lastY == node.y && node.x == node.lastX-1 && node.steps > 2 {
			return false
		}
		// prevent going backwards
		if node.lastY == node.y && node.x == node.lastX+1 {
			return false
		}

	case "right":
		// prevent going in a straight line for more than 3 steps
		if node.lastY == node.y && node.x == node.lastX+1 && node.steps > 2 {
			return false
		}
		// prevent going backwards
		if node.lastY == node.y && node.x == node.lastX-1 {
			return false
		}
	}
	return true
}

// helper that adds steps if continuing in the same direction
func calculateStepCost(dir string, node GridState) int {
	switch dir {
	case "up":
		if node.lastY-1 == node.y && node.x == node.lastX {
			return node.steps + 1
		}
	case "down":
		if node.lastY+1 == node.y && node.x == node.lastX {
			return node.steps + 1
		}
	case "left":
		if node.lastY == node.y && node.x == node.lastX-1 {
			return node.steps + 1
		}
	case "right":
		if node.lastY == node.y && node.x == node.lastX+1 {
			return node.steps + 1
		}
	}
	return 1
}

func day17Two() int {
	return 0
}

func Day17() {
	start := time.Now()
	one := day17One()
	two := day17Two()
	elapsed := time.Since(start)
	fmt.Printf("Day17, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
