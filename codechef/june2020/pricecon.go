package main

import (
	"fmt"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {

	var t, n, k int

	fmt.Scanln(&t)

	output := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)
		var r, a int

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &a)
			if a > k {
				r += a - k
			}
		}

		output[i] = r
	}
	utils.PrintSlice(output)
}
