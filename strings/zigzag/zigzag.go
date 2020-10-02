package main

import (
	"strings"
)

//https://leetcode.com/problems/zigzag-conversion/ problem link

func main() {

	convert("PAYPALISHIRING", 3)
	convert("PAYPALISHIRING", 4)
	convert("AB", 1)
}

func convert(s string, numRows int) string {
	output := make([]string, numRows)

	if numRows == 1 {
		return s
	}

	rowIndex := 0
	f := 1
	for i := 0; i < len(s); i++ {
		output[rowIndex] += string(s[i])
		if rowIndex == numRows-1 {
			f = -1
		}
		if rowIndex == 0 {
			f = 1
		}
		rowIndex += f
	}

	return strings.Join(output, "")
}
