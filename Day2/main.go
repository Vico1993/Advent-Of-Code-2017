package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitStringInToInt(s string) []int {
	var tmp = strings.Split(s, "\t")
	var digits = make([]int, len(tmp))

	for i := range tmp {
		var val, err = strconv.Atoi(tmp[i])
		if err != nil {
			log.Fatal("Error convert String to Int", err)
		}
		digits[i] = val
	}

	return digits
}

func findMaxAndMin(l []int) (int, int) {
	var min = 999999999
	var max = 0

	for _, v := range l {
		if v <= min {
			min = v
		}

		if v >= max {
			max = v
		}
	}

	return max, min
}

func findDividor(l []int) int {
	var mod = 0
	for _, v := range l {
		for _, v2 := range l {
			if v != v2 && v%v2 == 0 {
				mod = v / v2
			}
		}
	}

	return mod
}

func main() {
	// Open File
	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)

	var CheckSumPart1 = 0
	var CheckSumPart2 = 0
	for fscanner.Scan() {

		var line = splitStringInToInt(fscanner.Text())
		var max, min = findMaxAndMin(line)
		CheckSumPart1 += max - min
		CheckSumPart2 += findDividor(line)
	}

	fmt.Println("CheckSumPart1 PART 1:", CheckSumPart1)
	fmt.Println("CheckSumPart2 PART 2:", CheckSumPart2)
}
