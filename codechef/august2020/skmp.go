package august2020

import "strings"

func findSmallestCombination(input string, pattern string) string {

	sHash := convertStringTohashMap(input)
	pHash := convertStringTohashMap(pattern)

	output := ""
	found := false
	for i := 'a'; i < 'z'; i++ {
		if !found && strings.ContainsRune(pattern, i) {
			for sHash[i] > pHash[i] {
				output += string(i)
				sHash[i]--
			}

			for key, value := range pHash {
				sHash[key] -= value
			}
			output += pattern
			found = true
		}

		output += strings.Repeat(string(i), sHash[i])
		sHash[i] = 0
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
