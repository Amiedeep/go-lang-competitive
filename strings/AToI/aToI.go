package main

import (
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Scan(&input)

	output := 0

	for i := 0; i < len(input); i++ {
		if input[i] < '0' || input[i] > '9' {
			log.Fatal("Not a number")
		}

		output = output*10 + int(input[i]-'0')
	}

	fmt.Println(output)
}
