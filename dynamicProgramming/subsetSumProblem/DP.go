package main

import "fmt"

func initializeSlice(input []int, sum int) [][]bool {
	dp := make([][]bool, len(input))
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	return dp
}
func hasSum(input []int, sum int) bool {
	dp := initializeSlice(input, sum)

	dp[0][0] = true
	if input[0] < sum {
		dp[0][input[0]] = true
	}

	for row := 1; row < len(dp); row++ {
		for column := 0; column < len(dp[row]); column++ {
			// fmt.Printf("row is: %d , column is: %d ", row, column)
			if input[row] > column {
				dp[row][column] = dp[row-1][column]
			} else {
				dp[row][column] = dp[row-1][column] || dp[row-1][column-input[row]]
			}
		}
	}
	return dp[len(input)-1][sum]
}

func printArray(input [][]bool) {
	for row, element := range input {
		for column := range element {
			fmt.Printf("%t  ", input[row][column])
		}
		fmt.Println()
	}
}

func main() {
	input := []int{3, 34, 4, 12, 5, 2}

	// for i := 40; i < 70; i++ {

	// 	fmt.Println(i, ": ", hasSum(input, i))
	// }
	fmt.Println(hasSum(input, 9))
}
