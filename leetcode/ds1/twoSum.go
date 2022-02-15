package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, e := range nums {
		if _, ok := hm[target-e]; ok {
			return []int{i, hm[target-e]}
		}
		hm[e] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
