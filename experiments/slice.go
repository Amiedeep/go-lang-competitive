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

	// var d [100]int
	// fmt.Println(unsafe.Sizeof(d)) //prints 800

	// e := make([]int, 20)

	// fmt.Println(unsafe.Sizeof(e)) //prints 24

	// var f = [100]int{}

	// fmt.Println(unsafe.Sizeof(f)) //prints 800

	// var f [100]struct{}

	// fmt.Println(unsafe.Sizeof(f)) //prints 0

	c = append(c, 5)
	utils.PrintSlice(c) // 5

}
