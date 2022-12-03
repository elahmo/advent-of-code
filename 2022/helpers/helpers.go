package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open("/Users/ahmet/Developer/advent-of-code/2022/input/" + filePath)
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

func LinesToInt(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for _, w := range lines {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func IntersectionSlices[T comparable](a, b []T) (c []T) {
	m := make(map[T]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func IntersectionInts(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func IntersectionStrings(a, b string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[string(item)] = true
	}

	for _, item := range b {
		if _, ok := m[string(item)]; ok {
			c = append(c, string(item))
		}
	}
	return
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
