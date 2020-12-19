package main

import (
	"bufio"
	"fmt"
	"os"
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

func linesToTimeline(lines []string) (time int, ids []int) {
	i, err := strconv.Atoi(lines[0])
	if err == nil {
		time = i
	}
	busses := strings.Split(lines[1], ",")
	for _, bus := range busses {
		i, err := strconv.Atoi(bus)
		if err == nil {
			ids = append(ids, i)
		}
	}
	return

}

func linesToTimelineWithOffsets(lines []string) (ids []int, offsets []int) {
	busses := strings.Split(lines[1], ",")
	for off, bus := range busses {
		i, err := strconv.Atoi(bus)
		if err == nil {
			ids = append(ids, i)
			offsets = append(offsets, off)
		}
	}
	return

}

func part1() {
	lines, _ := fileToLines("input/day13.txt")
	time, busses := linesToTimeline(lines)
	delay := 999999999
	busID := 0
	for _, bus := range busses {
		div, rem := time/bus, time%bus
		if rem == 0 {
			busID = bus
			delay = 0
			break
		}
		diff := (div+1)*bus - time
		if diff < delay {
			delay = diff
			busID = bus
		}
	}
	fmt.Println(delay * busID)
}

func part2() {
	lines, _ := fileToLines("input/day13.txt")
	busses, offsets := linesToTimelineWithOffsets(lines)
	timestamp := 0
	increment := busses[0]
	for i := 1; i < len(busses); i++ {
		for {
			timestamp += increment
			if (timestamp+offsets[i])%busses[i] == 0 {
				increment *= busses[i]
				break
			}
		}
	}
	fmt.Println(timestamp)
}

func main() {
	part2()
}
