package utils

import "fmt"

func PrintSlice(input []int) {
	for _, element := range input {
		fmt.Printf("%d  ", element)
		fmt.Println()
	}
}

func Print2DSlice(input [][]int) {
	for _, row := range input {
		for _, column := range row {
			fmt.Printf("%d  ", column)
		}
		fmt.Println()
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Create2DIntSlice creates and returns a 2d slice.
func Create2DIntSlice(length int) [][]int {
	slice := make([][]int, length)
	for i := range slice {
		slice[i] = make([]int, length)
	}

	return slice
}
