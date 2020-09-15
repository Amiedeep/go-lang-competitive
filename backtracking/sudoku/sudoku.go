package main

import (
	"math"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {

	sol := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	res := sudoku(0, sol)

	if res {
		utils.Print2DSlice(sol)
	}

}

func isValid(row int, col int, sol [][]int, number int) bool {
	cellRow := (row / 3) * 3
	cellCol := (col / 3) * 3

	for i := cellRow; i < cellRow+3; i++ {
		for j := cellCol; j < cellCol+3; j++ {
			if sol[i][j] == number {
				return false
			}
		}
	}

	for i := 0; i < len(sol); i++ {
		if sol[row][i] == number || sol[i][col] == number {
			return false
		}
	}

	return true
}

func sudoku(index int, sol [][]int) bool {
	if float64(index) == math.Pow(float64(len(sol)), 2) {
		return true
	}

	row := index / len(sol)
	col := index % len(sol)

	if sol[row][col] != 0 {
		return sudoku(index+1, sol)
	}

	for i := 1; i < 10; i++ {

		if isValid(row, col, sol, i) {
			sol[row][col] = i

			if sudoku(index+1, sol) {
				return true
			}
			sol[row][col] = 0
		}
	}
	return false
}
