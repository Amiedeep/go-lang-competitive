package main

import "fmt"

func main() {
	// m := make(map[int]bool)
	m := map[int]bool{}

	m[1] = true
	m[2] = true
	m[3] = true

	//prints 1,3
	for key := range m {
		fmt.Println(key)
		delete(m, 2)
	}

	var a map[int]int

	if a == nil {
		fmt.Println("nil map") //prints
	}
	fmt.Println(a[5]) //prints 0
	// if a[5] == nil { // doesnt work since map value is int (cannot compare)
	// 	fmt.Println(a[5])
	// }
	// a[5] = 5  This gives nil map exception

	a = make(map[int]int)

	if a == nil {
		fmt.Println("nil map") // doesnt print
	}

	// fmt.Println(a[5]) //prints 0
	a[5] = 5
	fmt.Println(a[5]) // prints 5

	z := map[int]struct{}{

		1: {},
	}

	for key, value := range z {
		fmt.Println(key, value)
	}

}
