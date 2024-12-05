package main

import (
	"aoc24/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	solutions.RunAllDays()
	fmt.Printf("Total runtime: %s\n", time.Since(start))
}
