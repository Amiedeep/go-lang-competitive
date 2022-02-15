func characterReplacement(s string, k int) int {
	start, end, maxChar, maxSubString := 0, k, 0, k+1
	hm := make(map[byte]int)

	for i := 0; i <= k; i++ {
		hm[s[i]]++
	}

	maxChar = findMax(hm)

	for end = k + 1; end < len(s); end++ {
		hm[s[end]]++
		maxChar = findMax(hm)

		if end-start-maxChar >= k {
			hm[s[start]]--
			start++
		}

		if end-start+1 > maxSubString {
			maxSubString = end - start + 1
		}
	}

	return maxSubString
}

func findMax(input map[byte]int) int {
	max := 1
	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}
