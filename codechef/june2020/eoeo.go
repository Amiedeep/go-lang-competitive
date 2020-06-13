package main

import (
	"fmt"
)

func main() {
	var t int
	var ts uint64

	fmt.Scanln(&t)

	o := make([]uint64, t)

	for i := 0; i < t; i++ {

		fmt.Scanln(&ts)

		o[i] = findWins(ts)
	}
	printSlice(o)
}

func printSlice(input []uint64) {
	for _, element := range input {
		fmt.Printf("%d  ", element)
		fmt.Println()
	}
}

func findWins(ts uint64) uint64 {
	for ts%2 == 0 {
		ts = ts / 2
	}

	return ts / 2
}
