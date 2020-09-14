package main

import (
	"github.com/amiedeep/go-lang-competitive/utils"
)

func next(sol [][]int, i int) bool {
	if i == len(sol) {
		return true
	}

	for k := 0; k < len(sol); k++ {
		if isSafe(sol, i, k) {
			sol[i][k] = 1
			res := next(sol, i+1)
			if res {
				return true
			}
			sol[i][k] = 0
		}
	}

	return false
}

func isSafe(sol [][]int, row int, col int) bool {
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if sol[i][j] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j < len(sol); i, j = i-1, j+1 {
		if sol[i][j] == 1 {
			return false
		}
	}

	for i := 0; i < len(sol); i++ {
		if sol[i][col] == 1 || sol[row][i] == 1 {
			return false
		}
	}

	return true
}

func main() {
	sol := make([][]int, 8)

	for row := range sol {
		sol[row] = make([]int, 8)
	}

	result := next(sol, 0)

	if result {
		utils.Print2DSlice(sol)
	}
}
