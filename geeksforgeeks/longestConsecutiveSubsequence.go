//https://practice.geeksforgeeks.org/problems/longest-consecutive-subsequence/0

package geeksforgeeks

import "github.com/amiedeep/go-lang-competitive/utils"

func longestConsecutiveSubsequence(a []int) int {
	var m int

	h := map[int]bool{}

	for _, value := range a {
		h[value] = true
	}

	for k := range h {
		c := 1
		t := k - 1

		for {
			_, ok := h[t]
			if !ok {
				break
			}
			c++
			delete(h, t)
			t--
		}

		t = k + 1
		for {
			_, ok := h[t]
			if !ok {
				break
			}
			c++
			delete(h, t)
			t++
		}

		m = utils.Max(c, m)

		delete(h, k)
	}

	return m
}
