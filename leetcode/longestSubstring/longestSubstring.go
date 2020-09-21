func lengthOfLongestSubstring(s string) int {

	longestSoFar, longest, start := 0, 0, 0
	h := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		
		
		if v, ok := h[s[i]]; ok && v >= start {
			start = v+1
		}

		h[s[i]] = i

		longest = i - start + 1
		if longest > longestSoFar {
			longestSoFar = longest
		}
	}
	return longestSoFar
}
