package solutions

import (
	"aoc24/helpers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Day05One() int {
	lines, _ := helpers.FileToLines("day05.txt")
	var rules [][2]int
	var updates [][]int
	doingUpdates := false
	for _, line := range lines {
		if line == "" {
			doingUpdates = true
			continue
		}
		if !doingUpdates {
			var num1, num2 int
			_, err := fmt.Sscanf(line, "%d|%d", &num1, &num2)
			if err != nil {
				panic(err)
			}
			rules = append(rules, [2]int{num1, num2})
		} else {
			parts := strings.Split(line, ",")
			integers := make([]int, len(parts))

			// Iterate over the parts and convert to integers
			for i, part := range parts {
				// Trim whitespace and convert to integer
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err != nil {
					panic(err)
				}
				integers[i] = num
			}
			updates = append(updates, integers)
		}
	}
	sum := 0 // Ensure `sum` is initialized
	for _, update := range updates {
		// Create a map to store the index of each number in the current update
		indexMap := make(map[int]int)
		for i, num := range update {
			indexMap[num] = i
		}

		// Assume rules are satisfied initially
		satisfied := true
		for _, rule := range rules {
			// Check if both elements in the rule exist in the map
			if idx1, ok1 := indexMap[rule[0]]; ok1 {
				if idx2, ok2 := indexMap[rule[1]]; ok2 && idx1 > idx2 {
					// Rule is violated
					satisfied = false
					break
				}
			}
		}

		// If rules are satisfied, add the midpoint value to the sum
		if satisfied {
			if len(update) == 0 {
				continue // Safeguard: skip empty updates
			}
			midpoint := len(update) / 2
			sum += update[midpoint]
		}
	}
	return sum
}

func Day05Two() int {
	lines, _ := helpers.FileToLines("day05.txt")
	var rules [][2]int
	var updates [][]int
	doingUpdates := false
	for _, line := range lines {
		if line == "" {
			doingUpdates = true
			continue
		}
		if !doingUpdates {
			var num1, num2 int
			_, err := fmt.Sscanf(line, "%d|%d", &num1, &num2)
			if err != nil {
				panic(err)
			}
			rules = append(rules, [2]int{num1, num2})
		} else {
			parts := strings.Split(line, ",")
			integers := make([]int, len(parts))

			// Iterate over the parts and convert to integers
			for i, part := range parts {
				// Trim whitespace and convert to integer
				num, err := strconv.Atoi(strings.TrimSpace(part))
				if err != nil {
					panic(err)
				}
				integers[i] = num
			}
			updates = append(updates, integers)
		}
	}
	sum := 0 // Ensure `sum` is initialized
	for _, update := range updates {
		// Create a map to store the index of each number in the current update
		indexMap := make(map[int]int)
		for i, num := range update {
			indexMap[num] = i
		}

		// Assume rules are satisfied initially
		satisfied := true
		for _, rule := range rules {
			// Check if both elements in the rule exist in the map
			if idx1, ok1 := indexMap[rule[0]]; ok1 {
				if idx2, ok2 := indexMap[rule[1]]; ok2 && idx1 > idx2 {
					// Rule is violated
					satisfied = false
					break
				}
			}
		}

		// If rules are satisfied, add the midpoint value to the sum
		if !satisfied {
			reordered_update, err := reorderToSatisfyRules(update, rules)
			if err != nil {
				panic(err)
			}
			if len(reordered_update) == 0 {
				continue // Safeguard: skip empty updates
			}
			midpoint := len(reordered_update) / 2
			sum += reordered_update[midpoint]
		}
	}
	return sum
}

func reorderToSatisfyRules(update []int, rules [][2]int) ([]int, error) {
	// Build the graph and in-degree map
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	elements := make(map[int]bool)

	// Initialize graph and in-degree map
	for _, num := range update {
		graph[num] = []int{}
		inDegree[num] = 0
		elements[num] = true
	}

	// Populate the graph with rules
	for _, rule := range rules {
		from, to := rule[0], rule[1]
		if elements[from] && elements[to] {
			graph[from] = append(graph[from], to)
			inDegree[to]++
		}
	}

	// Perform Kahn's Algorithm for Topological Sorting
	var sorted []int
	queue := []int{}

	// Add nodes with zero in-degree to the queue
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}

	// Process the queue
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue
		sorted = append(sorted, current)

		// Reduce in-degree of neighbors
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if topological sorting is possible
	if len(sorted) != len(update) {
		return nil, fmt.Errorf("cycle detected, cannot satisfy all rules")
	}

	// Reorder the update slice according to the sorted order
	position := make(map[int]int)
	for i, num := range sorted {
		position[num] = i
	}

	// Sort the update slice based on the topological order
	reordered := make([]int, len(update))
	for _, num := range update {
		reordered[position[num]] = num
	}

	return reordered, nil
}

func Day05() {
	start := time.Now()
	one := Day05One()
	two := Day05Two()
	elapsed := time.Since(start)
	fmt.Printf("Day05, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day05", Day05)
}
