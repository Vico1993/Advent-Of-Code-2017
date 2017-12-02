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

func main() {
	// Open File
	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)

	var CheckSumPart1 = 0
	for fscanner.Scan() {
		var min = 999999999
		var max = 0

		var line = splitStringInToInt(fscanner.Text())
		for _, v := range line {
			if v <= min {
				min = v
			}

			if v >= max {
				max = v
			}
		}

		CheckSumPart1 += max - min
	}

	fmt.Println("CheckSumPart1 PART 1:", CheckSumPart1)
}
