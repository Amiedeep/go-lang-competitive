package main

import "fmt"

func main() {
	var c int

	fmt.Scanln(&c)

	o := make([]string, c)

	for i := 0; i < c; i++ {
		var f, t, v, n int
		fmt.Scanln(&n)
		o[i] = "YES"
		for j := 0; j < n; j++ {
			fmt.Scan(&v)
			switch v {
			case 5:
				f++
			case 10:
				if f > 0 {
					f--
					t++
				} else {
					o[i] = "NO"
					break
				}
			case 15:
				if t > 0 {
					t--
				} else if f > 1 {
					f = f - 2
				} else {
					o[i] = "NO"
					break
				}
			}
		}
	}
	printSlice(o)
}

func printSlice(input []string) {
	for _, element := range input {
		fmt.Printf("%s  ", element)
		fmt.Println()
	}
}
