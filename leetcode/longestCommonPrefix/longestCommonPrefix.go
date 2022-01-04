package main

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		charAtI := strs[0][i]

		for _, v := range strs {
			if len(v) <= i || v[i] != charAtI {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

// func main() {
// 	input := []string{"flower", "flow", "flight"}

// 	fmt.Println(longestCommonPrefix(input))
// }
