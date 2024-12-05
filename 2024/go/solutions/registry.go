package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Day struct holds the day name and its associated function
type Day struct {
	Name string
	Fn   func()
}

// Registry to store days in order
var days []Day

// RegisterDay allows each day to register its function
func RegisterDay(name string, fn func()) {
	days = append(days, Day{
		Name: name,
		Fn:   fn,
	})
}

// RunAllDays executes all registered days in sorted order based on the day name
func RunAllDays() {
	// Sort days by extracting the numeric part from the name (e.g., "Day01" -> 1)
	sort.Slice(days, func(i, j int) bool {
		numI := extractDayNumber(days[i].Name)
		numJ := extractDayNumber(days[j].Name)
		return numI < numJ
	})

	// Execute each day's function
	for _, day := range days {
		day.Fn()
	}
}

// extractDayNumber extracts the numeric part from a day name like "DayXX"
func extractDayNumber(name string) int {
	// Remove the "Day" prefix and parse the remaining number
	numStr := strings.TrimPrefix(name, "Day")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid day name format: %s", name))
	}
	return num
}
