package main

import (
	"fmt"
	// util "github.com/amiedeep/go-lang-competitive/utils"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findMaxPrice(prices []int) int {

	maxPrice := make([]int, len(prices))
	copy(maxPrice, prices)

	for i := 0; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			maxPrice[i] = max(maxPrice[i], maxPrice[j]+maxPrice[i-j-1])
		}
	}

	return maxPrice[len(maxPrice)-1]
}

func main() {

	// prices := []int{1, 5, 8, 9, 10, 17, 17, 20}
	// prices := []int{5, 4, 1}
	prices := []int{3, 5, 8, 9, 10, 17, 17, 20}

	maxValue := findMaxPrice(prices)

	// util.PrintSlice(prices)

	fmt.Println(maxValue)

}
