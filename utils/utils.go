package utils

import "fmt"

func PrintSlice(input []int) {
	for _, element := range input {
		fmt.Printf("%d  ", element)
		fmt.Println()
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
