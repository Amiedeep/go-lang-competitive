package main


import (
	"fmt"

	"github.com/amiedeep/go-lang-competitive/utils"
)
func main() {

	maze := [][]int {
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 1},
	}

	sol := make([][]int, len(maze))

	for row := range sol {
		sol[row] = make([]int, len(maze))
	}

	fmt.Println(next(sol, maze, 0, 0))
	fmt.Println("Path is: ")
	utils.Print2DSlice(sol)
}

func next(sol [][]int, maze [][]int, i int, j int) bool {

	if(i == len(maze)-1 && j == len(maze)-1 && maze[i][j] == 1) {
		sol[i][j] = 1
		return true
	}
	
	if(i >= len(maze) || j >= len(maze) || maze[i][j] == 0) {
		return false
	}

	res := next(sol, maze, i+1, j) || next(sol, maze, i, j+1)
	if(res) {
		sol[i][j] = 1
	}

	return res
}

