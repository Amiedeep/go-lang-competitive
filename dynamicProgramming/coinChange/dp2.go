package coinchange

func findWaysDP2(input []int, sum int) int {
	sol := make([]int, sum+1)
	sol[0] = 1

	for i := range input {
		for j := input[i]; j < len(sol); j++ {
			sol[j] += sol[j-input[i]]
		}

	}

	return sol[sum]
}
