package main

// problem link https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	cP := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {

		if v, ok := cP[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
