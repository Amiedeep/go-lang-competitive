package main

import (
	"fmt"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {
	a := []int{1, 2}

	// last element is exclusive
	a = a[0 : len(a)-1]

	utils.PrintSlice(a) // 1

	var b []int

	b = append(b, 3)

	utils.PrintSlice(b) // 3

	var c []int

	if c == nil {
		fmt.Println("nil slice") // print
	}

	c = append(c, 5)
	utils.PrintSlice(c) // 5

}
