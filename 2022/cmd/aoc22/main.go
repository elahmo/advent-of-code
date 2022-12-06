package main

import (
	"aoc22/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	solutions.Day01()
	solutions.Day02()
	solutions.Day03()
	solutions.Day04()
	solutions.Day05()
	solutions.Day06()
	fmt.Printf("Total runtime: %s\n", time.Since(start))
}
