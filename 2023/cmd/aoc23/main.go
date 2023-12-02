package main

import (
	"aoc23/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// solutions.Day01()
	solutions.Day02()
	fmt.Printf("Total runtime: %s\n", time.Since(start))
}
