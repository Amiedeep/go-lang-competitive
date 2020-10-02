package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1, str2, result string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	i, j := len(str1)-1, len(str2)-1

	valI, valJ, carry, value := 0, 0, 0, 0
	for ; ; i, j = i-1, j-1 {
		if i < 0 && j < 0 {
			break
		}

		if i < 0 {
			valI = 0
			valJ, _ = strconv.Atoi(string(str2[j]))

		} else if j < 0 {
			valI, _ = strconv.Atoi(string(str1[i]))
			valJ = 0
		} else {
			valI, _ = strconv.Atoi(string(str1[i]))
			valJ, _ = strconv.Atoi(string(str2[j]))
		}
		value = (valI + valJ + carry) % 10
		carry = (valI + valJ + carry) / 10

		result = strconv.Itoa(value) + result
	}
	if carry == 1 {
		result = strconv.Itoa(carry) + result
	}

	fmt.Println(result)
}
