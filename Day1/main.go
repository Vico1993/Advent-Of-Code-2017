package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func parseToInt(s string) []int {
	digits := make([]int, len(s))
	for i := range s {
		val, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			log.Fatal("Error")
		}
		digits[i] = val
	}
	return digits
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error on the file")
	}

	// Convert Bytes to String
	dataStr := string(data[:])
	// Convert String to Array []Int
	AInt := parseToInt(dataStr)

	var num int
	for i := range AInt {
		next := 0
		if i+1 < len(AInt) {
			next = i + 1
		}
		if AInt[i] == AInt[next] {
			num += AInt[i]
		}
	}

	println(num)
}
