package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
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

func linesToInstructions(lines []string) ([]string, []int) {
	ins := make([]string, len(lines))
	steps := make([]int, len(lines))
	var rex = regexp.MustCompile("(\\w)(\\d+)")
	for idx, w := range lines {
		data := rex.FindAllStringSubmatch(w, -1)
		ins[idx] = data[0][1]
		i, err := strconv.Atoi(data[0][2])
		if err == nil {
			steps[idx] = i
		}
	}
	return ins, steps
}

func updateBearing(bearing string, ins string, steps int) (newBearing string) {
	right := "ESWN"
	left := "ENWS"
	switch ins {
	case "R":
		idx := strings.Index(right, bearing)
		shift := (idx + (steps / 90)) % 4
		newBearing = string(right[shift])
	case "L":
		idx := strings.Index(left, bearing)
		shift := (idx + (steps / 90)) % 4
		newBearing = string(left[shift])
	}
	return
}

func updateWaypoint(wx, wy int, ins string, steps int) (newwx, newwy int) {
	shift := (steps / 90) % 4
	switch ins {
	case "R":
		switch shift {
		case 1:
			newwx = wy
			newwy = -wx
		case 2:
			newwx = -wx
			newwy = -wy
		case 3:
			newwx = -wy
			newwy = wx
		}
	case "L":
		switch shift {
		case 1:
			newwx = -wy
			newwy = wx
		case 2:
			newwx = -wx
			newwy = -wy
		case 3:
			newwx = wy
			newwy = -wx
		}
	}
	return
}

func part1() {
	lines, _ := fileToLines("input/day12.txt")
	ins, steps := linesToInstructions(lines)
	bearing := "E"
	x, y := 0, 0
	for i := 0; i < len(lines); i++ {
		steps := steps[i]
		switch ins[i] {
		case "F":
			switch bearing {
			case "E":
				x += steps
			case "W":
				x -= steps
			case "N":
				y += steps
			case "S":
				y -= steps
			}
		case "E":
			x += steps
		case "W":
			x -= steps
		case "N":
			y += steps
		case "S":
			y -= steps
		case "R":
			bearing = updateBearing(bearing, "R", steps)
		case "L":
			bearing = updateBearing(bearing, "L", steps)
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func part2() {
	lines, _ := fileToLines("input/day12.txt")
	ins, steps := linesToInstructions(lines)
	wx := 10
	wy := 1
	x, y := 0, 0
	for i := 0; i < len(lines); i++ {
		steps := steps[i]
		switch ins[i] {
		case "F":
			x += wx * steps
			y += wy * steps
		case "E":
			wx += steps
		case "W":
			wx -= steps
		case "N":
			wy += steps
		case "S":
			wy -= steps
		case "R":
			wx, wy = updateWaypoint(wx, wy, "R", steps)
		case "L":
			wx, wy = updateWaypoint(wx, wy, "L", steps)
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	part1()
	part2()
}
