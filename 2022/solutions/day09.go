package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"time"
)

type Knot struct {
	x int
	y int
}

type GridMove struct {
	x     int
	y     int
	steps int
}

const (
	center = 500
	width  = 2*center + 1
)

func day09One() int {
	lines, _ := helpers.FileToLines("day09.txt")
	grid := make([]bool, width*width)
	head, tail := Knot{center, center}, Knot{center, center}
	movesMap := map[string]GridMove{
		"L": {-1, 0, 0},
		"R": {1, 0, 0},
		"U": {0, 1, 0},
		"D": {0, -1, 0},
	}
	for _, line := range lines {
		re := regexp.MustCompile(`^(\w) (\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		move := movesMap[m[0]]
		move.steps = helpers.StrToInt(m[1])
		moveHeadAndTail(&head, &tail, &grid, move)
	}
	visitedCount := 0
	for _, visited := range grid {
		if visited {
			visitedCount++
		}
	}
	return visitedCount
}

func moveHeadAndTail(head *Knot, tail *Knot, grid *[]bool, move GridMove) {
	for i := 0; i < move.steps; i++ {
		// move head
		head.x += move.x
		head.y += move.y

		// move tail
		// diagonal touching
		if helpers.Abs(head.y-tail.y) == 1 && helpers.Abs(head.x-tail.x) == 1 {
			continue
		}
		// same row
		if head.y == tail.y {
			// touching
			if helpers.Abs(head.x-tail.x) <= 1 {
				continue
			}
			if head.x > tail.x {
				tail.x++
			} else {
				tail.x--
			}
			(*grid)[tail.y*width+tail.x] = true
			continue
		}
		// same column
		if head.x == tail.x {
			// touching
			if helpers.Abs(head.y-tail.y) <= 1 {
				continue
			}
			if head.y > tail.y {
				tail.y++
			} else {
				tail.y--
			}
			(*grid)[tail.y*width+tail.x] = true
			continue
		}
		// diagonal catchup
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
		(*grid)[tail.y*width+tail.x] = true
	}
}

func day09Two() int {
	lines, _ := helpers.FileToLines("day09.txt")
	grid := make([]bool, width*width)
	rope := make([]Knot, 10)

	//set all knots to be in center
	for i := 0; i < len(rope); i++ {
		rope[i].x = center
		rope[i].y = center
	}

	movesMap := map[string]GridMove{
		"L": {-1, 0, 0},
		"R": {1, 0, 0},
		"U": {0, 1, 0},
		"D": {0, -1, 0},
	}

	for _, line := range lines {
		re := regexp.MustCompile(`^(\w) (\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		move := movesMap[m[0]]
		move.steps = helpers.StrToInt(m[1])
		moveRope(&rope, &grid, move)
	}
	visitedCount := 0
	for _, visited := range grid {
		if visited {
			visitedCount++
		}
	}
	return visitedCount
}

func moveRope(rope *[]Knot, grid *[]bool, move GridMove) {
	for i := 0; i < move.steps; i++ {
		// move head according to the rule
		head := &(*rope)[0]
		head.x += move.x
		head.y += move.y

		// move segments individualy
		for knot := 0; knot < len((*rope))-1; knot++ {
			segOne := &(*rope)[knot]
			segTwo := &(*rope)[knot+1]

			stopMoves := moveSegment(segOne, segTwo, grid, knot == 8)
			if stopMoves {
				break
			}
		}
	}
}

func moveSegment(head *Knot, tail *Knot, grid *[]bool, tailMove bool) bool {
	// move head
	// diagonal touching
	if helpers.Abs(head.y-tail.y) == 1 && helpers.Abs(head.x-tail.x) == 1 {
		return true
	}
	// same row
	if head.y == tail.y {
		// touching
		if helpers.Abs(head.x-tail.x) <= 1 {
			return true
		}
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
		if tailMove {
			(*grid)[tail.y*width+tail.x] = true
		}
		return false
	}
	// same column
	if head.x == tail.x {
		// touching
		if helpers.Abs(head.y-tail.y) <= 1 {
			return true
		}
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
		if tailMove {
			(*grid)[tail.y*width+tail.x] = true
		}
		return false
	}
	// diagonal catchup
	if head.y > tail.y {
		tail.y++
	} else {
		tail.y--
	}
	if head.x > tail.x {
		tail.x++
	} else {
		tail.x--
	}
	if tailMove {
		(*grid)[tail.y*width+tail.x] = true
	}
	return false
}

func Day09() {
	start := time.Now()
	one := day09One()
	two := day09Two()
	elapsed := time.Since(start)
	fmt.Printf("Day09, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
