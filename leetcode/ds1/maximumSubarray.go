package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max, currentMax := math.MinInt32, 0

	fmt.Println("5")
	for i := 0; i < len(nums); i++ {
		currentMax = currentMax + nums[i]

		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

	}
	return max
}

// func main() {
// 	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
// }
