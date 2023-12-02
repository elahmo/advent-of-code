package solutions

import (
	"aoc22/helpers"
	"fmt"
	"strconv"
	"time"
)

const ()

func day25One() int {
	lines, _ := helpers.FileToLines("day25.txt")
	// numMap := map[string]int{
	// 	"2": 2,
	// 	"1": 1,
	// 	"0": 0,
	// 	"-": -1,
	// 	"=": -2,
	// }
	sum := 0
	// for _, line := range lines {
	// 	num := 0
	// 	for i := 0; i < len(line); i++ {
	// 		base := int(math.Pow(5, float64(len(line)-1-i)))
	// 		mult := numMap[string(line[i])]
	// 		num += base * mult
	// 	}
	// 	sum += num
	// }
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		fmt.Println(negabinary(num))
	}
	return sum
}

func toSnafu(number int) []int {
	// numMap := map[int]string{
	// 	2:  "2",
	// 	1:  "1",
	// 	0:  "0",
	// 	-1: "-",
	// 	-2: "=",
	// }
	// find the largest power
	// pow := 19

	// sum := int(math.Pow(5, float64(19)))
	// remainder := number - sum
	// numbers := []int{}
	// for i := 0; i < pow; i++ {
	// 	// go one power down
	// 	base := int(math.Pow(5, float64(19-1-i)))
	// 	// try top is less than next power top, fine
	// 	if base*2 < remainder {
	// 		numbers = append(numbers, 2)
	// 		continue
	// 	}
	// 	// try top  -1 is less than next power top, fine
	// 	if base*1 < remainder {
	// 		numbers = append(numbers, 1)
	// 		continue
	// 	}
	// 	// try bot is less than next power top, fine
	// 	if base*1 < remainder {
	// 		numbers = append(numbers, 1)
	// 		continue
	// 	}
	// 	// try bot +1  is less than next power top, fine
	// 	// otherwise 0
	// 	// num += base * mult
	// }

	// bruteforce ftw
	numbers := []int{}

	// start from the right,

	return numbers
}

func negabinary(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n != 0 {
		remainder := n % -4
		n = n / -4
		if remainder < 0 {
			remainder += 4
			n++
		}
		result = string(remainder+'0') + result
	}

	return result
}

func day25Two() int {
	return 0
}

func Day25() {
	start := time.Now()
	one := day25One()
	two := day25Two()
	elapsed := time.Since(start)
	fmt.Printf("Day25, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
