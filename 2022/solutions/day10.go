package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"time"
)

const ()

func day10One() int {
	lines, _ := helpers.FileToLines("day10.txt")
	signal := 0
	register := 1
	cycle := 1
	queued := map[int]int{}
	screen := make([]string, 40*6)
	for _, line := range lines {
		noop := false
		re := regexp.MustCompile(`^(\w+)( ([-+]?\d+))?`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]

		if m[0] == "noop" {
			noop = true
		} else {
			queued[cycle+1] = helpers.StrToInt(m[2])
		}

		if noop {
			if cycle == 20 || (cycle-20)%40 == 0 {
				signal += cycle * register
			}
			// check if queued for execution in this cycle
			inst, exists := queued[cycle]
			if exists {
				register += inst
				delete(queued, cycle)
			}
			if shouldDraw(cycle, register) {
				screen[cycle-1] = "#"
			} else {
				screen[cycle-1] = "."
			}
			cycle++
			continue
		}
		for i := 0; i < 2; i++ {
			if cycle == 20 || (cycle-20)%40 == 0 {
				signal += cycle * register
			}
			// check if queued for execution in this cycle
			inst, exists := queued[cycle]
			if exists {
				register += inst
				delete(queued, cycle)
			}
			if shouldDraw(cycle, register) {
				screen[cycle-1] = "#"
			} else {
				screen[cycle-1] = "."
			}
			cycle++
		}

	}
	// printScreen(screen)
	return signal
}

func shouldDraw(cycle int, register int) bool {
	if (cycle%40) >= register-1 && (cycle%40) <= register+1 {
		return true
	}
	return false
}

func printScreen(screen []string) {
	for i, letter := range screen {
		if (i+1)%40 == 0 {
			fmt.Println(letter)
			continue
		}
		fmt.Print(letter)
	}
}

func day10Two() int {
	return 0
}

func Day10() {
	start := time.Now()
	one := day10One()
	two := day10Two()
	elapsed := time.Since(start)
	fmt.Printf("Day10, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
