package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func linesToInt(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for _, w := range lines {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}


func parts1and2() {
	lines, _ := fileToLines("input/day05.txt")
	
	var largestSeatId int64
	var seenIds [1024]bool
	largestSeatId = 0
	for _, seat := range lines {
		binarySeat := ""
		for _, char := range seat{
			switch string(char) {
			case "F":
				binarySeat += "0"
			case "B":
				binarySeat += "1"
			case "L":
				binarySeat += "0"
			case "R":
				binarySeat += "1"
			}
		}
		row, _ := strconv.ParseInt(binarySeat[:7], 2, 64)
		column, _ := strconv.ParseInt(binarySeat[7:], 2, 64)
		seatId := (row*8 + column)
		seenIds[seatId] = true
		if seatId > largestSeatId {
			largestSeatId = row*8 + column
		}
	}
	frontSeats := true
	for seatId, occupied := range seenIds {
		if !frontSeats && !occupied {
			fmt.Println(seatId)
			break
		}
		if frontSeats && occupied {
			frontSeats = false
		}
	}
	fmt.Println(largestSeatId)
	return
}

func main() {
	parts1and2()
}
