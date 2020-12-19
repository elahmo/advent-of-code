package main

import (
	"bufio"
	"fmt"
	"os"
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

func updateSeat(seatX, seatY int, seatmap [][]string) (newLetter string, updated bool) {
	seat := seatmap[seatX][seatY]
	fieldsToCheck := [3][3]bool{{true, true, true}, {true, false, true}, {true, true, true}}
	if seatX == 0 {
		fieldsToCheck[0][0] = false
		fieldsToCheck[0][1] = false
		fieldsToCheck[0][2] = false
	} else if seatX == len(seatmap)-1 {
		fieldsToCheck[2][0] = false
		fieldsToCheck[2][1] = false
		fieldsToCheck[2][2] = false
	}
	if seatY == 0 {
		fieldsToCheck[0][0] = false
		fieldsToCheck[1][0] = false
		fieldsToCheck[2][0] = false
	} else if seatY == len(seatmap[0])-1 {
		fieldsToCheck[0][2] = false
		fieldsToCheck[1][2] = false
		fieldsToCheck[2][2] = false
	}
	occupiedCount := 0
	for idx, fields := range fieldsToCheck {
		for idy, check := range fields {
			if check {
				if seatmap[seatX+(idx-1)][seatY+(idy-1)] == "#" {
					occupiedCount++
				}
			}
		}
	}

	switch seat {
	case "L":
		if occupiedCount == 0 {
			newLetter = "#"
			updated = true
		}
	case "#":
		if occupiedCount >= 4 {
			newLetter = "L"
			updated = true
		}
	}
	return
}

func updateSeatVisible(seatX, seatY int, seatmap [][]string) (newLetter string, updated bool) {
	seat := seatmap[seatX][seatY]
	maxX := len(seatmap) - 1
	maxY := len(seatmap[0]) - 1
	movement := make(map[int][]int)
	movement[0] = []int{-1, -1}
	movement[1] = []int{-1, 0}
	movement[2] = []int{-1, 1}
	movement[3] = []int{0, -1}
	movement[4] = []int{0, 0}
	movement[5] = []int{0, 1}
	movement[6] = []int{1, -1}
	movement[7] = []int{1, 0}
	movement[8] = []int{1, 1}
	occupiedCount := 0
	for idx := 0; idx < 3; idx++ {
		for idy := 0; idy < 3; idy++ {
			mvmt := movement[idx*3+idy]
			x := seatX
			y := seatY
			validMovement := true
			for validMovement {
				x += mvmt[0]
				y += mvmt[1]
				if x < 0 || y < 0 || x > maxX || y > maxY || (mvmt[0] == 0 && mvmt[1] == 0) {
					validMovement = false
					break
				}
				visibleSeat := seatmap[x][y]
				switch visibleSeat {
				case ".":
					continue
				case "#":
					occupiedCount++
					validMovement = false
				case "L":
					validMovement = false
				}

			}
		}
	}

	switch seat {
	case "L":
		if occupiedCount == 0 {
			newLetter = "#"
			updated = true
		}
	case "#":
		if occupiedCount >= 5 {
			newLetter = "L"
			updated = true
		}
	}
	return
}

func generateSeatMap(seatMap []string) [][]string {
	board := make([][]string, len(seatMap))
	for row, seat := range seatMap {
		board[row] = make([]string, len(seat))
		for i, c := range seat {
			board[row][i] = string(c)
		}
	}
	return board
}

func countSeats(seatMap [][]string, targetSeat string) (count int) {
	for _, seat := range seatMap {
		for _, c := range seat {
			if string(c) == targetSeat {
				count++
			}
		}
	}
	return
}

func part1() {
	lines, _ := fileToLines("input/day11.txt")
	seatMap := generateSeatMap(lines)
	iterations := 0
	for {
		madeChanges := false
		updatedSeatMap := make([][]string, len(seatMap))
		for i := range updatedSeatMap {
			updatedSeatMap[i] = make([]string, len(seatMap[i]))
			copy(updatedSeatMap[i], seatMap[i])
		}
		for row, seats := range seatMap {
			for i := range seats {
				newSeat, updated := updateSeat(row, i, seatMap)
				if updated {
					madeChanges = true
					updatedSeatMap[row][i] = newSeat
				}
			}
		}
		if !madeChanges {
			break
		}
		iterations++
		copy(seatMap, updatedSeatMap)
		seatMap := updatedSeatMap
		_ = seatMap //really hate that I have to do this in go
	}
	fmt.Println(iterations, countSeats(seatMap, "#"))
}

func part2() {
	lines, _ := fileToLines("input/day11.txt")
	seatMap := generateSeatMap(lines)
	iterations := 0
	for {
		madeChanges := false
		updatedSeatMap := make([][]string, len(seatMap))
		for i := range updatedSeatMap {
			updatedSeatMap[i] = make([]string, len(seatMap[i]))
			copy(updatedSeatMap[i], seatMap[i])
		}
		for row, seats := range seatMap {
			for i := range seats {
				newSeat, updated := updateSeatVisible(row, i, seatMap)
				if updated {
					madeChanges = true
					updatedSeatMap[row][i] = newSeat
				}
			}
		}
		if !madeChanges {
			break
		}
		iterations++
		copy(seatMap, updatedSeatMap)
		seatMap := updatedSeatMap
		_ = seatMap //really hate that I have to do this in go
		// for _, line := range seatMap {
		// 	fmt.Println(line)
		// }
		// fmt.Println("-------")
	}
	fmt.Println(iterations, countSeats(seatMap, "#"))
}

func main() {
	// part1()
	part2()
}
