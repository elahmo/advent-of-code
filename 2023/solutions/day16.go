package solutions

import (
	"aoc23/helpers"
	"fmt"
	"time"
)

type Position struct {
	x   int
	y   int
	mov [2]int // movement vector, 0 for x, 1 for y
	// 10 right, -10 left, 01 up, 0-1 down
}

func day16One() int {
	lines, _ := helpers.FileToLines("day16.txt")
	LEN := len(lines)
	grid := make([]bool, LEN*LEN)
	signs := make([]string, LEN*LEN)
	for row, line := range lines {
		// add chars from line
		for col, char := range line {
			signs[row*LEN+col] = string(char)
		}
	}

	rayDb := make(map[Position]bool)
	pos := Position{0, 0, [2]int{1, 0}}
	calculateRays(&grid, &signs, &rayDb, LEN, pos)

	// calculate score
	sum := 0
	for _, v := range grid {
		// visual output is optional
		if v {
			// fmt.Print("#")
			sum++
		} else {
			// fmt.Print(".")
		}
		// if (idx+1)%LEN == 0 {
		// 	fmt.Println()
		// }
	}
	return sum
}

func calculateRays(grid *[]bool, signs *[]string, rayDb *map[Position]bool, LEN int, pos Position) {
	curPos := pos
	rayComplete := false
	for {
		//if out of bounds, stop movement
		if curPos.x > LEN-1 || curPos.y > LEN-1 || curPos.x < 0 || curPos.y < 0 || rayComplete {
			rayComplete = true
			break
		}
		idx := curPos.x + curPos.y*LEN
		symbol := (*signs)[idx]
		(*grid)[idx] = true
		switch symbol {
		case ".":
			// continue movement
			curPos.x += curPos.mov[0]
			curPos.y += curPos.mov[1]
		case "/":
			// mirror
			// going down, go left
			if curPos.mov[1] == 1 {
				curPos = Position{curPos.x - 1, curPos.y, [2]int{-1, 0}}
				break
			}
			// going up, go right
			if curPos.mov[1] == -1 {
				curPos = Position{curPos.x + 1, curPos.y, [2]int{1, 0}}
				break
			}
			// going left, go down
			if curPos.mov[0] == -1 {
				curPos = Position{curPos.x, curPos.y + 1, [2]int{0, 1}}
				break
			}
			// going right, go up
			if curPos.mov[0] == 1 {
				curPos = Position{curPos.x, curPos.y - 1, [2]int{0, -1}}
				break
			}
		case "\\":
			// mirror
			// going down, go right
			if curPos.mov[1] == 1 {
				curPos = Position{curPos.x + 1, curPos.y, [2]int{1, 0}}
				break
			}
			// going up, go left
			if curPos.mov[1] == -1 {
				curPos = Position{curPos.x - 1, curPos.y, [2]int{-1, 0}}
				break
			}
			// going left, go up
			if curPos.mov[0] == -1 {
				curPos = Position{curPos.x, curPos.y - 1, [2]int{0, -1}}
				break
			}
			// going right, go down
			if curPos.mov[0] == 1 {
				curPos = Position{curPos.x, curPos.y + 1, [2]int{0, 1}}
				break
			}
		case "-":
			// split or stop
			// going left or right, continue movement
			if curPos.mov[1] == 0 {
				curPos.x += curPos.mov[0]
				curPos.y += curPos.mov[1]
				break
			}
			// going up or down, split left and right
			if curPos.mov[0] == 0 {
				posA := Position{curPos.x - 1, curPos.y, [2]int{-1, 0}}
				if _, ok := (*rayDb)[posA]; !ok {
					(*rayDb)[posA] = true
					calculateRays(grid, signs, rayDb, LEN, posA)
				}
				posB := Position{curPos.x + 1, curPos.y, [2]int{1, 0}}
				if _, ok := (*rayDb)[posB]; !ok {
					(*rayDb)[posB] = true
					calculateRays(grid, signs, rayDb, LEN, posB)
				}
				rayComplete = true
				break
			}
		case "|":
			// split or stop
			// going up or down, continue movement
			if curPos.mov[0] == 0 {
				curPos.x += curPos.mov[0]
				curPos.y += curPos.mov[1]
				break
			}
			// going left or right, split up and down
			if curPos.mov[1] == 0 {
				posA := Position{curPos.x, curPos.y - 1, [2]int{0, -1}}
				if _, ok := (*rayDb)[posA]; !ok {
					(*rayDb)[posA] = true
					calculateRays(grid, signs, rayDb, LEN, posA)
				}
				posB := Position{curPos.x, curPos.y + 1, [2]int{0, 1}}
				if _, ok := (*rayDb)[posB]; !ok {
					(*rayDb)[posB] = true
					calculateRays(grid, signs, rayDb, LEN, posB)
				}
				rayComplete = true
				break
			}
		}
	}
}

func day16Two() int {
	lines, _ := helpers.FileToLines("day16.txt")
	LEN := len(lines)
	signs := make([]string, LEN*LEN)
	for row, line := range lines {
		// add chars from line
		for col, char := range line {
			signs[row*LEN+col] = string(char)
		}
	}

	// do all combos
	maxSum := 0
	for x := 0; x < LEN; x++ {
		for y := 0; y < LEN; y++ {
			for version := 0; version < 2; version++ {
				if x == 0 || y == 0 || x == LEN-1 || y == LEN-1 {
					grid := make([]bool, LEN*LEN)
					rayDb := make(map[Position]bool)
					// chose a direction
					dir := choseDirection(x, y, version, LEN)
					pos := Position{x, y, dir}
					calculateRays(&grid, &signs, &rayDb, LEN, pos)
					sum := 0
					for _, v := range grid {
						if v {
							sum++
						}
					}
					if sum > maxSum {
						maxSum = sum
					}
				}
			}
		}
	}
	return maxSum
}

func choseDirection(x, y, version, LEN int) [2]int {
	// top left
	if x == 0 && y == 0 {
		if version == 0 { // right
			return [2]int{1, 0}
		}
		return [2]int{0, 1} // down
	}
	// top right
	if x == LEN-1 && y == 0 {
		if version == 0 {
			return [2]int{-1, 0} // left
		}
		return [2]int{0, 1} // down
	}
	// bottom left
	if x == 0 && y == LEN-1 {
		if version == 0 {
			return [2]int{1, 0} // right
		}
		return [2]int{0, -1} // up
	}
	// bottom right
	if x == LEN-1 && y == LEN-1 {
		if version == 0 {
			return [2]int{-1, 0} // left
		}
		return [2]int{0, -1} // up
	}
	if x == 0 {
		return [2]int{1, 0}
	}
	if x == LEN-1 {
		return [2]int{-1, 0}
	}
	if y == 0 {
		return [2]int{0, 1}
	}
	if y == LEN-1 {
		return [2]int{0, -1}
	}
	panic(0) // should never happen
}

func Day16() {
	start := time.Now()
	one := day16One()
	two := day16Two()
	elapsed := time.Since(start)
	fmt.Printf("Day16, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
