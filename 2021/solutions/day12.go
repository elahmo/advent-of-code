package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
	"unicode"
)

// type node struct {
// 	bracket  string
// 	closed   bool
// 	children []node
// }

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

func part1() {
	lines, _ := fileToLines("input/day12.txt")
	caveMap := buildCaveMap(lines)

	paths := iterateThroughCaves([]string{"start"}, caveMap)
	fmt.Println(paths, len(paths))
	// fmt.Println(caveMap)
}

// start,A,c,A,b,A,end

func iterateThroughCaves(path []string, caveMap map[string][]string) [][]string {
	validPaths := [][]string{}
	cave := path[len(path)-1]
	caves := caveMap[cave]
	for _, targetCave := range caves {
		if targetCave == "end" {
			finalPath := []string{}
			for _, pth := range path {
				finalPath = append(finalPath, pth)
			}
			finalPath = append(finalPath, targetCave)
			validPaths = append(validPaths, finalPath)
		} else {
			// if small cave and in path already, do not visit it
			if IsUpper(targetCave) || !contains(path, targetCave) {
				tempPath := append(path, targetCave)
				paths := iterateThroughCaves(tempPath, caveMap)
				validPaths = append(validPaths, paths...)
			}
		}
	}
	return validPaths
}

func part2() {
	lines, _ := fileToLines("input/day12.txt")
	caveMap := buildCaveMap(lines)

	paths := iterateThroughCavesTwo([]string{"start"}, caveMap)
	paths = cleanPaths(paths)
	fmt.Println(paths, len(paths))
	// fmt.Println(caveMap)
}

func cleanPaths(paths [][]string) [][]string {
	validPaths := [][]string{}
	for _, validPath := range paths {
		counter := map[string]int{}
		for _, item := range validPath {
			if IsUpper(item) {
				continue
			}
			_, ok := counter[item]
			if !ok {
				counter[item] = 0
			}
			counter[item]++
		}
		start, ok := counter["start"]
		if ok && start > 1 {
			continue
		}
		end, ok := counter["end"]
		if ok && end > 1 {
			continue
		}
		doubleCount := 0
		for key, count := range counter {
			if key == "start" || key == "end" {
				continue
			}
			if count > 1 {
				doubleCount++
			}
		}
		if doubleCount < 2 {
			validPaths = append(validPaths, validPath)
		}

	}

	return validPaths
}

func iterateThroughCavesTwo(path []string, caveMap map[string][]string) [][]string {
	validPaths := [][]string{}
	cave := path[len(path)-1]
	caves := caveMap[cave]
	for _, targetCave := range caves {
		if targetCave == "end" {
			finalPath := []string{}
			for _, pth := range path {
				finalPath = append(finalPath, pth)
			}
			finalPath = append(finalPath, targetCave)
			validPaths = append(validPaths, finalPath)
		} else {
			// if small cave and in path already, do not visit it
			if IsUpper(targetCave) || !containsTwo(path, targetCave) {
				tempPath := append(path, targetCave)
				paths := iterateThroughCavesTwo(tempPath, caveMap)
				validPaths = append(validPaths, paths...)
			}
		}
	}
	return validPaths
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func containsTwo(arr []string, s string) bool {
	count := 0
	for _, a := range arr {
		if a == s {
			count++
		}
	}
	return count > 1
}

func contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func buildCaveMap(lines []string) map[string][]string {
	caveMap := map[string][]string{}
	rex := regexp.MustCompile(`\w+`)
	for _, line := range lines {
		data := rex.FindAllString(line, -1)
		_, ok := caveMap[data[0]]
		if !ok {
			caveMap[data[0]] = []string{}
		}
		_, ok = caveMap[data[1]]
		if !ok {
			caveMap[data[1]] = []string{}
		}
		caveMap[data[0]] = append(caveMap[data[0]], data[1])
		caveMap[data[1]] = append(caveMap[data[1]], data[0])
	}
	return caveMap
}

func main() {
	start := time.Now()
	part2()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
