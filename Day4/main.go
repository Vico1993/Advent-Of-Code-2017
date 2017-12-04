package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isInArray(array []string, s string) bool {
	for _, d := range array {
		if d == s {
			return true
		}
	}

	return false
}

func main() {
	// Open File
	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)

	// Part 1
	var data []string
	count := 0

	for fscanner.Scan() {
		passPhraseIsRight := true
		data = data[:0]

		for _, d := range strings.Split(fscanner.Text(), " ") {
			if isInArray(data, d) {
				passPhraseIsRight = false
				break
			}

			data = append(data, d)
		}
		if passPhraseIsRight {
			count++
		}
	}

	fmt.Println(count)
}
