package main

import "fmt"

func main() {
	n := 8

	pattern := make([][]byte, n)

	for i := range pattern {
		pattern[i] = make([]byte, n)
	}

	for i, left, right := 0, n/2, n/2; i < n && left >= 0 && right < n; i, left, right = i+1, left-1, right+1 {
		pattern[i][left], pattern[i][right] = '*', '*'

		if i == n/4 {
			for j := left + 1; j < right; j++ {
				pattern[i][j] = '*'
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pattern[i][j] == '*' {
				fmt.Printf("%c ", pattern[i][j])
			} else {
				fmt.Printf("%c ", '.')
			}

		}
		fmt.Println()
	}
}
