package main

import (
	"fmt"

	"github.com/amiedeep/go-lang-competitive/utils"
)

func main() {
	var t int
	var s string
	fmt.Scanln(&t)

	o := make([]int, t)

	for i := 0; i < t; i++ {

		fmt.Scanln(&s)
		var p int
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] {
				p++
				i = i + 1
			}
		}
		o[i] = p
	}

	utils.PrintSlice(o)
}
