package main

import "fmt"

var (
	dToC = map[string][]string{"2": {"a", "b", "c"}, "3": {"d", "e", "f"}, "4": {"g", "h", "i"},
		"5": {"j", "k", "l"}, "6": {"m", "n", "o"}, "7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"}, "9": {"w", "x", "y", "z"}}
)

func findCombinations(currentcombs []string, digits string, index int) []string {
	if index >= len(digits) {
		return currentcombs
	}
	digit := string(digits[index])

	if v, ok := dToC[string(digit)]; ok {
		combs := make([]string, 0)
		for _, digitVal := range v {
			newCombs := make([]string, 0)
			for _, comb := range currentcombs {
				newCombs = append(newCombs, comb+digitVal)
			}
			combs = append(combs, findCombinations(newCombs, digits, index+1)...)
		}
		return combs
	}
	return currentcombs
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return make([]string, 0)
	}
	return findCombinations([]string{""}, digits, 0)
}

func main() {
	fmt.Println(letterCombinations("23"))
}
