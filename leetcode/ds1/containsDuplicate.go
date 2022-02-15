package main

// func main() {
// 	var input string

// 	// fmt.Scanln(&input)
// 	// fmt.Scanf("%s", &input)
// 	fmt.Println(input)

// }

func containsDuplicate(nums []int) bool {
	hm := make(map[int]bool)

	for _, n := range nums {
		if _, ok := hm[n]; ok {
			return true
		}
		hm[n] = true
	}
	return false
}
