func shortestDistanceColor(colors []int, queries [][]int) []int {
	answers := make([]int, len(queries))
	for i := 0; i < len(answers); i++ {
		answers[i] = -1
	}

	for index, query := range queries {
		k, c := query[0], query[1]
		for i := 0; ; i++ {
			if k-i >= 0 && colors[k-i] == c {
				answers[index] = i
				break
			}
			if i+k < len(colors) && colors[k+i] == c {
				answers[index] = i
				break
			}
		}
	}
	return answers

}