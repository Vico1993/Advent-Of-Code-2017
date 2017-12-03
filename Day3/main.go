package main

import "fmt"

const num = 368078

func getMiddleOfArret(i int) []int {
	var s []int
	carre := i * i
	midOfArret := i - (i / 2)

	r := carre - midOfArret + 1

	for len(s) < 4 {
		s = append(s, r)
		r = r - i + 1
	}
	return s
}

func main() {
	i := 1
	for i*i < num {
		i += 2
	}

	fmt.Println("Num:", num)

	// Distance between 1 and the carre of i
	oneToCube := (i / 2)

	// Get 4 middl of CubeOfI for each cube side
	middleOfCube := getMiddleOfArret(i)

	// Get Neariest middle of the cube
	min := 999999
	for _, m := range middleOfCube {
		x := m - num
		if x < 0 {
			x = 0 - x
		}
		if x < min {
			min = x
		}
	}

	// Distance is oneToCube + min
	fmt.Println("Disctance is : ", oneToCube+min)
}
