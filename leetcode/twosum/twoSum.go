package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, element := range nums {
		required := target - element
		if value, ok := m[required]; ok {
			return []int{value, index}
		}
		m[element] = index
	}
	return nil
}
