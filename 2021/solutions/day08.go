package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
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

func part1() {
	lines, _ := fileToLines("input/day08.txt")

	rex := regexp.MustCompile(`\w+`)
	counts := 0
	for _, line := range lines {
		data := rex.FindAllStringSubmatch(line, -1)
		for _, matches := range data[10:] {
			if len(matches[0]) < 5 || len(matches[0]) == 7 {
				counts++
			}
		}
	}
	fmt.Println(counts)
}

func sortString(input string) string {
	s := []rune(input)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func removeDuplicates(input string) string {
	s := []rune(input)
	return string(removeDuplicateStr(s))
}

func removeDuplicateStr(strSlice []rune) []rune {
	allKeys := make(map[rune]bool)
	list := []rune{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func difference(one, two string) string {
	a := []rune(one)
	b := []rune(two)
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []rune
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return string(diff)
}

func contains(a, b string) bool {
	contains := true
	for i := 0; i < len(b); i++ {
		if !strings.Contains(a, string(b[i])) {
			contains = false
		}
	}
	return contains

}

func part2() {
	lines, _ := fileToLines("input/day08.txt")

	rex := regexp.MustCompile(`\w+`)
	counts := 0
	for _, line := range lines {
		chars := map[string]int{}
		nums := map[int]string{}
		data := rex.FindAllStringSubmatch(line, -1)
		for _, matches := range data {
			word := sortString(matches[0])
			switch len(word) {
			case 2:
				chars[word] = 1
				nums[1] = word
			case 3:
				chars[word] = 7
				nums[7] = word
			case 4:
				chars[word] = 4
				nums[4] = word
			case 7:
				chars[word] = 8
				nums[8] = word
			}
		}
		easyNums := removeDuplicates(nums[1] + nums[7])
		diff := difference(nums[4], easyNums)
		for _, matches := range data {
			word := sortString(matches[0])
			switch len(word) {
			case 5:
				// 235
				if contains(word, diff) {
					chars[word] = 5
					nums[5] = word
				} else if contains(word, nums[1]) {
					chars[word] = 3
					nums[3] = word
				} else {
					chars[word] = 2
					nums[2] = word
				}
			case 6:
				// 960
				if contains(word, diff+nums[1]) {
					chars[word] = 9
					nums[9] = word
				} else if contains(word, diff) {
					chars[word] = 6
					nums[6] = word
				} else {
					chars[word] = 0
					nums[0] = word
				}
			}
		}
		counts += chars[sortString(data[10][0])]*1000 + chars[sortString(data[11][0])]*100 + chars[sortString(data[12][0])]*10 + chars[sortString(data[13][0])]
		// fmt.Println(nums)
	}
	fmt.Println(counts)
}

func main() {
	start := time.Now()
	part2()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
