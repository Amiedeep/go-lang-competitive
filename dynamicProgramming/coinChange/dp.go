package coinchange

func findWaysDP(input []int, sum int) int {
	sol := make([][]int, len(input)+1)

	for i := range sol {
		sol[i] = make([]int, sum+1)
	}

	for i := range sol[0] {
		sol[0][i] = 0
	}

	for i := range sol {
		sol[i][0] = 1
	}

	for i := 1; i <= len(input); i++ {
		for j := 1; j <= sum; j++ {
			if input[i-1] > j {
				sol[i][j] = sol[i-1][j]
			} else {
				sol[i][j] = sol[i-1][j] + sol[i][j-input[i-1]]
			}
		}
	}

	return sol[len(input)][sum]
}
