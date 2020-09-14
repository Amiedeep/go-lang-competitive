package main

import (
	"math"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {

	sol := make([][]int, 8)

	for i := range sol {
		sol[i] = make([]int, 8)
	}

	nextPosI := []int{2, 1, -1, -2, -2, -1, 1, 2}
	nextPosJ := []int{1, 2, 2, 1, -1, -2, -2, -1}

	res := kT(sol, 0, 0, 1, nextPosI, nextPosJ)

	if res {
		utils.Print2DSlice(sol)
	}

}

func kT(sol [][]int, i int, j int, num int, nextPosI []int, nextPosJ []int) bool {
	// fmt.Printf("%d-%d", i, j)
	// fmt.Println()
	if num == int(math.Pow(float64(len(sol)), 2)) {
		return true
	}

	sol[i][j] = num

	for k := 0; k < len(nextPosI); k++ {
		nextI := i + nextPosI[k]
		nextJ := j + nextPosJ[k]

		if isValid(nextI, nextJ, sol) {
			res := kT(sol, nextI, nextJ, num+1, nextPosI, nextPosJ)

			if res {
				return true
			}
		}
	}

	sol[i][j] = 0

	return false
}

func isValid(i int, j int, sol [][]int) bool {
	if i < 0 || i >= len(sol) || j < 0 || j >= len(sol[i]) {
		return false
	}

	return sol[i][j] == 0
}
