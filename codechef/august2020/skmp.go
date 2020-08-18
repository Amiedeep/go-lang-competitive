package august2020

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	var s, p string
	fmt.Scanln(&t)

	output := make([]string, t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&s)
		fmt.Scanln(&p)
		output[i] = findSmallestCombination(s, p)
	}
	printSliceString(output)
}

func printSliceString(input []string) {
	for _, element := range input {
		fmt.Printf("%s  ", element)
		fmt.Println()
	}
}

func findSmallestCombination(input string, pattern string) string {
	sHash := convertStringTohashMap(input)
	pHash := convertStringTohashMap(pattern)

	var output string
	for i := 'a'; i <= 'z'; i++ {
		if pattern[0] == byte(i) {

			str1 := output + pattern + strings.Repeat(string(i), sHash[i]-pHash[i])

			str2 := output + strings.Repeat(string(i), sHash[i]-pHash[i]) + pattern

			if strings.Compare(str1, str2) > 0 {
				output = str2
			} else {
				output = str1
			}
			continue
		}

		output += strings.Repeat(string(i), sHash[i]-pHash[i])
	}
	return output
}

func convertStringTohashMap(input string) map[rune]int {
	output := map[rune]int{}

	for _, char := range input {
		output[char]++
	}

	return output
}
