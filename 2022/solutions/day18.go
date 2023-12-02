package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"time"
)

const (
	scale18 = 23
)

func day18One() int {
	lines, _ := helpers.FileToLines("day18.txt")
	grid := make([]bool, scale18*scale18*scale18)
	for _, line := range lines {
		re := regexp.MustCompile("(\\d+),(\\d+),(\\d+)")
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		x, y, z := helpers.StrToInt(m[0]), helpers.StrToInt(m[1]), helpers.StrToInt(m[2])
		grid[z*scale18*scale18+y*scale18+x] = true
	}
	// check number of sides available
	sides := 0
	for z := 0; z < scale18; z++ {
		for y := 0; y < scale18; y++ {
			for x := 0; x < scale18; x++ {
				// only cover present cubes
				if !grid[z*scale18*scale18+y*scale18+x] {
					continue
				}
				cur := 6
				if x < scale18 && grid[(z*scale18*scale18+y*scale18+x)+1] {
					cur--
				}
				if x > 0 && grid[(z*scale18*scale18+y*scale18+x)-1] {
					cur--
				}
				if y < scale18 && grid[(z*scale18*scale18+y*scale18+x)+scale18] {
					cur--
				}
				if y > 0 && grid[(z*scale18*scale18+y*scale18+x)-scale18] {
					cur--
				}
				if z < scale18 && grid[(z*scale18*scale18+y*scale18+x)+scale18*scale18] {
					cur--
				}
				if z > 0 && grid[(z*scale18*scale18+y*scale18+x)-scale18*scale18] {
					cur--
				}
				sides += cur
			}
		}
	}
	return sides
}

func day18Two() int {
	lines, _ := helpers.FileToLines("day18.txt")
	grid := make([]bool, scale18*scale18*scale18)
	for _, line := range lines {
		re := regexp.MustCompile("(\\d+),(\\d+),(\\d+)")
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		x, y, z := helpers.StrToInt(m[0]), helpers.StrToInt(m[1]), helpers.StrToInt(m[2])
		grid[z*scale18*scale18+y*scale18+x] = true
	}
	// calculate where can air reach, starting from 0,0,0
	airGrid := propagateLava(grid)
	// check number of sides that are emtpy and have air reaching them
	sides := 0
	for x := 0; x < scale18; x++ {
		for y := 0; y < scale18; y++ {
			for z := 0; z < scale18; z++ {
				// only cover present cubes
				idx := z*scale18*scale18 + y*scale18 + x
				if !grid[idx] {
					continue
				}
				cur := 0
				if airGrid[idx+1] {
					cur++
				}
				if airGrid[idx-1] {
					cur++
				}
				if airGrid[idx+scale18] {
					cur++
				}
				if idx > scale18 && airGrid[idx-scale18] {
					// if airGrid[idx-scale18] {
					cur++
				}
				if airGrid[idx+scale18*scale18] {
					cur++
				}
				if idx > scale18*scale18-1 && airGrid[idx-scale18*scale18] {
					// if airGrid[idx-scale18*scale18] {
					cur++
				}
				sides += cur
			}
		}
	}
	return sides
}

func propagateLava(grid []bool) []bool {
	// start from zero position
	airGrid := make([]bool, scale18*scale18*scale18)
	airGrid[0] = true
	// initialize get all empty pockets next to it and add them to the list
	toExplore := getEmptyPockets(0, 0, 0, grid, airGrid)
	// iterate while there are things to explore
	for {
		newExplore := map[[3]int]bool{}
		// start from all empty air pockets that are reachable
		for p := range toExplore {
			airGrid[p[2]*scale18*scale18+p[1]*scale18+p[0]] = true
			reachableFromCurrent := getEmptyPockets(p[0], p[1], p[2], grid, airGrid)
			for newP := range reachableFromCurrent {
				newExplore[newP] = true
			}
		}
		toExplore = newExplore
		if len(toExplore) == 0 {
			break
		}
	}
	return airGrid
}

func getEmptyPockets(x, y, z int, grid []bool, airGrid []bool) map[[3]int]bool {
	emptyPockets := map[[3]int]bool{}
	// do not go out of bounds
	// if x+1 >= scale18 || y+1 >= scale18 || z+1 >= scale18 {
	// 	return emptyPockets
	// }
	idx := z*scale18*scale18 + y*scale18 + x
	grLn := len(grid)
	if x < scale18 && idx+1 < grLn && !grid[idx+1] && !airGrid[idx+1] {
		emptyPockets[[3]int{x + 1, y, z}] = true
	}
	if x > 0 && !grid[idx-1] && !airGrid[idx-1] {
		emptyPockets[[3]int{x - 1, y, z}] = true
	}
	if y < scale18 && idx+scale18 < grLn && !grid[idx+scale18] && !airGrid[idx+scale18] {
		emptyPockets[[3]int{x, y + 1, z}] = true
	}
	if y > 0 && !grid[idx-scale18] && !airGrid[idx-scale18] {
		emptyPockets[[3]int{x, y - 1, z}] = true
	}
	if z < scale18 && (idx+scale18*scale18) < grLn && !grid[idx+scale18*scale18] && !airGrid[idx+(scale18*scale18)] {
		emptyPockets[[3]int{x, y, z + 1}] = true
	}
	if z > 0 && !grid[idx-scale18*scale18] && !airGrid[idx-(scale18*scale18)] {
		emptyPockets[[3]int{x, y, z - 1}] = true
	}
	return emptyPockets
}

func Day18() {
	start := time.Now()
	one := day18One()
	two := day18Two()
	elapsed := time.Since(start)
	fmt.Printf("Day18, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
