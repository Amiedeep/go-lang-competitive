package main

// problem url https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {

	digitToLetter := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	if len(digits) == 0 {
		return make([]string, 0)
	}
	return findCombinations(digits, "", digitToLetter)
}

func findCombinations(digits string, currentCombination string, digitToLetter map[byte][]byte) []string {

	if len(currentCombination) == len(digits) {
		return []string{currentCombination}
	}

	comb := digitToLetter[digits[len(currentCombination)]]

	output := make([]string, 0)
	for i := 0; i < len(comb); i++ {
		output = append(output, findCombinations(digits, currentCombination+string(comb[i]), digitToLetter)...)
	}

	return output
}
