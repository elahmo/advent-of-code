package solutions

import (
	"aoc24/helpers"
	"fmt"
	"time"
)

func Day25One() int {
	lines, _ := helpers.FileToLines("day25.txt")
	var keys [][]int
	var locks [][]int
	schemaChange := true
	var isLock bool
	currentSchematic := make([]int, 5)
	for _, line := range lines {
		if line == "" {
			schemaChange = true
			if isLock {
				locks = append(locks, currentSchematic)
			} else {
				keySchematic := make([]int, 5)
				for i := 0; i < len(currentSchematic); i++ {
					keySchematic[i] = currentSchematic[i] - 1
					if keySchematic[i] < 0 {
						keySchematic[i] = 0
					}
				}
				keys = append(keys, keySchematic)
			}
			currentSchematic = []int{0, 0, 0, 0, 0}
			continue
		}
		if schemaChange {
			schemaChange = false
			if line == "#####" {
				isLock = true
			} else {
				isLock = false
			}
			continue
		}

		for idx, element := range line {
			if element == '#' {
				currentSchematic[idx] += 1
			}
		}
	}
	// add last
	if isLock {
		locks = append(locks, currentSchematic)
	} else {
		keySchematic := make([]int, 5)
		for i := 0; i < len(currentSchematic); i++ {
			keySchematic[i] = currentSchematic[i] - 1
			if keySchematic[i] < 0 {
				keySchematic[i] = 0
			}
		}
		keys = append(keys, keySchematic)
	}

	// fmt.Println("locks", locks)
	// fmt.Println("keys", keys)
	count := 0
	uniquePairs := make(map[string]bool)
	for _, lock := range locks {
		for _, key := range keys {
			match := true
			for i := 0; i < len(key); i++ {
				if key[i]+lock[i] >= 6 {
					match = false
					break
				}
			}
			if match {
				pair := fmt.Sprintf("%v-%v", lock, key)
				if !uniquePairs[pair] {
					// fmt.Println("Matching lock and key:", lock, key)
					uniquePairs[pair] = true
					count += 1
				} else {
					// fmt.Println("found duplicate", lock, key)
				}
			}
		}
	}
	return count
}

func Day25Two() int {
	return 0
}

func Day25() {
	start := time.Now()
	one := Day25One()
	two := Day25Two()
	elapsed := time.Since(start)
	fmt.Printf("Day25, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day25", Day25)
}
