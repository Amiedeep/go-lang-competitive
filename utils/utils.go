package utils

import "fmt"

func PrintSlice(input []int) {
	for _, element := range input {
		fmt.Printf("%d  ", element)
		fmt.Println()
	}
}
