package coinchange

func findWaysRecursive(input []int, sum int, index int) int {
	if sum == 0 {
		return 1
	}
	if sum < 0 || index >= len(input) {
		return 0
	}

	return findWaysRecursive(input, sum, index+1) + findWaysRecursive(input, sum-input[index], index)
}
