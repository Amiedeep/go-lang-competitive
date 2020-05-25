package main

import "fmt"

func hasSum(input []int, sum int, index int) bool {

	if sum == 0 {
		return true
	}
	if sum < 0 || len(input) == index {
		return false
	}

	return hasSum(input, sum-input[index], index+1) || hasSum(input, sum, index+1)
}
func main() {
	input := []int{3, 34, 4, 12, 5, 2}

	// var sum int
	// fmt.Println("Enter sum: ")
	// fmt.Scan(&sum)

	fmt.Println(hasSum(input, 9, 0))
}
