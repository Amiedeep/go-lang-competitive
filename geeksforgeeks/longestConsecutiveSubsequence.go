//https://practice.geeksforgeeks.org/problems/longest-consecutive-subsequence/0

package geeksforgeeks

import "github.com/amiedeep/go-lang-competitive/utils"

func longestConsecutiveSubsequence(a []int) int {
	var m int

	set := utils.NewSet()

	for _, value := range a {
		set.Add(value)
	}

	for k := range set.KeySet {
		c := 1
		t := k - 1

		for {
			if !set.Contains(t) {
				break
			}
			c++
			set.Delete(t)
			t--
		}

		t = k + 1
		for {
			if !set.Contains(t) {
				break
			}
			c++
			set.Delete(t)
			t++
		}

		m = utils.Max(c, m)

		set.Delete(k)
	}

	return m
}
