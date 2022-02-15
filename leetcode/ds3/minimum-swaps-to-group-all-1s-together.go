func minSwaps(data []int) int {
	maxOnes, left, totalOnes, currentOnes := 0, 0, 0, 0
	for _, e := range data {
		if e == 1 {
			totalOnes++
		}
	}

	for i := 0; i < len(data); i++ {
		if data[i] == 1 {
			currentOnes++
		}
		if i >= left+totalOnes {
			if data[left] == 1 {
				currentOnes--
			}
			left++
		}
		if currentOnes > maxOnes {
			maxOnes = currentOnes
		}
	}

	return totalOnes - maxOnes
}