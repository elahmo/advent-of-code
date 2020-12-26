package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func linesToInstructions(lines []string) ([]string, map[string]map[uint64]uint64, map[string][]uint64) {
	var instructions = map[string]map[uint64]uint64{}
	var order = []string{}
	var orderIns = map[string][]uint64{}
	currentKey := ""
	rex := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	for _, line := range lines {
		maskIndex := strings.Index(line, "mask")
		if maskIndex == 0 {
			currentKey = string(line[7:])
			instructions[currentKey] = map[uint64]uint64{}
			orderIns[currentKey] = []uint64{}
			order = append(order, currentKey)
			continue
		}
		m := rex.FindAllStringSubmatch(line, -1)
		mem, err := strconv.ParseUint(m[0][1], 10, 64)
		if err != nil {
			panic(err)
		}
		val, err := strconv.ParseUint(m[0][2], 10, 64)
		if err != nil {
			panic(err)
		}
		instructions[currentKey][mem] = val
		orderIns[currentKey] = append(orderIns[currentKey], mem)
	}
	return order, instructions, orderIns

}

func applyMask(s string, mask string) uint64 {
	tempNum := ""
	for i := 0; i < len(s); i++ {
		if string(mask[i]) == "X" {
			tempNum += string(s[i])
		} else {
			tempNum += string(mask[i])
		}
	}
	num, err := strconv.ParseUint(tempNum, 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func applyMaskV2(s string, mask string) (addresses []uint64) {
	tempNum := ""
	for i := 0; i < len(s); i++ {
		switch string(mask[i]) {
		case "X":
			tempNum += "X"
		case "0":
			tempNum += string(s[i])
		case "1":
			tempNum += "1"
		}
	}
	tempAddresses := []string{tempNum}
	done := false
	for {
		done = true
		updatedAddresses := []string{}
		for _, address := range tempAddresses {
			if strings.Index(address, "X") == -1 {
				continue
			}
			addressZero := strings.Replace(address, "X", "0", 1)
			addressOne := strings.Replace(address, "X", "1", 1)
			updatedAddresses = append(updatedAddresses, addressZero)
			updatedAddresses = append(updatedAddresses, addressOne)
			done = false
		}
		if done {
			break
		}
		tempAddresses = updatedAddresses
		_ = tempAddresses
	}
	for _, address := range tempAddresses {
		num, err := strconv.ParseUint(address, 2, 64)
		if err != nil {
			panic(err)
		}
		addresses = append(addresses, num)
	}
	return
}

func part1() {
	lines, _ := fileToLines("input/day14.txt")
	order, instructions, _ := linesToInstructions(lines)
	memory := make(map[uint64]uint64)
	for _, mask := range order {
		ins := instructions[mask]
		for address, num := range ins {
			s := fmt.Sprintf("%036b", num)
			memory[address] = applyMask(s, mask)
		}
	}
	var sum uint64
	for _, val := range memory {
		sum += uint64(val)
	}
	fmt.Println(sum)
}

func part2() {
	lines, _ := fileToLines("input/day14.txt")
	order, instructions, orderIns := linesToInstructions(lines)
	memory := make(map[uint64]uint64)
	for _, mask := range order {
		ins := instructions[mask]
		for _, address := range orderIns[mask] {
			value := ins[address]
			s := fmt.Sprintf("%036b", address)
			for _, address := range applyMaskV2(s, mask) {
				memory[address] = value
			}

		}
	}
	var sum uint64
	for _, val := range memory {
		sum += uint64(val)
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
