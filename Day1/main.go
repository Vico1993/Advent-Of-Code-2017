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
		println(val)
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
	parseToInt(dataStr)

}
