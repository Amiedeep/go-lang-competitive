package main

func numberOfArithmeticSlices(nums []int) int {
	total := 0
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		start, end := i, i
		for j := i; j < len(nums)-1 && nums[j+1]-nums[j] == diff; j++ {
			end = j + 1
			i = j
		}

		if end >= start+2 {
			for k := 1; k < end-start; k++ {
				total += k
			}
		}
	}
	return total
}
