package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func linesToInt(lines []string) []int {
	ints := make([]int, 0, len(lines)*len(lines[0]))
	for _, w := range lines {
		for i := 0; i < len(w); i++ {
			i, err := strconv.Atoi(string(w[i]))
			if err == nil {
				ints = append(ints, i)
			}
		}
	}
	return ints
}

func unique(arr []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func part1() {
	lines, _ := fileToLines("input/day09.txt")

	ints := linesToInt(lines)
	w := len(lines[0])
	var lowPoints []int
	var lowPointsIdxes []int

	for idx, num := range ints {
		if idx == 0 {
			if ints[idx+1] > num && ints[w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx == w-1 {
			if ints[idx-1] > num && ints[2*w-1] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx == w*(len(lines)-1) {
			if ints[idx+1] > num && ints[w*(len(lines)-1)] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx == len(ints)-1 {
			if ints[idx-1] > num && ints[w*len(lines)-1] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx < w {
			if ints[idx-1] > num && ints[idx+1] > num && ints[idx+w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx > w*(len(lines)-1) {
			if ints[idx-1] > num && ints[idx+1] > num && ints[idx-w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx%w == 0 {
			if ints[idx+1] > num && ints[idx+w] > num && ints[idx-w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else if idx%w == w-1 {
			if ints[idx-1] > num && ints[idx+w] > num && ints[idx-w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		} else {
			if ints[idx-1] > num && ints[idx+1] > num && ints[idx+w] > num && ints[idx-w] > num {
				lowPoints = append(lowPoints, num)
				lowPointsIdxes = append(lowPointsIdxes, idx)
			}
		}
	}

	// fmt.Println(len(lowPoints))

	basins := map[int][]int{}
	for _, lowPointIdx := range lowPointsIdxes {
		basinNums := checkNeighbours(ints, lowPointIdx, w, len(lines))
		// fmt.Println(basinNums)
		basins[lowPointIdx] = unique(basinNums)
	}

	lengths := []int{}
	for _, values := range basins {
		sort.Ints(values)
		// fmt.Println(ints[key], len(values), values)
		lengths = append(lengths, len(values)+1)
	}
	sort.Ints(lengths)
	fmt.Println(lengths[len(lengths)-3:])
}

func checkNeighbours(ints []int, idx int, w int, h int) []int {
	var valid []int
	low := ints[idx]
	if low == 8 {
		return valid
	}
	if idx == 0 {
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[w] == low+1 {
			valid = append(valid, w)
		}
	} else if idx == w-1 {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[2*w-1] == low+1 {
			valid = append(valid, 2*w-1)
		}
	} else if idx == w*(h-1) {
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[w*(h-1)] == low+1 {
			valid = append(valid, w*(h-1))
		}
	} else if idx == len(ints)-1 {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[w*h-1] == low+1 {
			valid = append(valid, w*h-1)
		}
	} else if idx < w {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[idx+w] == low+1 {
			valid = append(valid, idx+w)
		}
	} else if idx > w*(h-1) {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[idx-w] == low+1 {
			valid = append(valid, idx-w)
		}
	} else if idx%w == 0 {
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[idx+w] == low+1 {
			valid = append(valid, idx+w)
		}
		if ints[idx-w] == low+1 {
			valid = append(valid, idx-w)
		}
	} else if idx%w == w-1 {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[idx+w] == low+1 {
			valid = append(valid, idx+w)
		}
		if ints[idx-w] == low+1 {
			valid = append(valid, idx-w)
		}
	} else {
		if ints[idx-1] == low+1 {
			valid = append(valid, idx-1)
		}
		if ints[idx+1] == low+1 {
			valid = append(valid, idx+1)
		}
		if ints[idx+w] == low+1 {
			valid = append(valid, idx+w)
		}
		if ints[idx-w] == low+1 {
			valid = append(valid, idx-w)
		}
	}
	if len(valid) > 0 {
		for i := 0; i < len(valid); i++ {
			newValid := checkNeighbours(ints, valid[i], w, h)
			valid = append(valid, newValid...)
		}
	}
	return valid
}

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
