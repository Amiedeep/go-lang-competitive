//https://practice.geeksforgeeks.org/problems/maximum-of-all-subarrays-of-size-k/0

package slidingwindow

func printMaxOfSubarrays(a []int, k int) []int {
	var q, o []int

	var i int

	for i = 0; i < k-1; i++ {
		q = pushElement(q, a[i])
	}

	for ; i < len(a); i++ {

		if len(q) >= k {
			q = q[1:]
		}

		q = pushElement(q, a[i])

		o = append(o, q[0])
	}

	return o
}

func pushElement(q []int, a int) []int {
	for len(q) != 0 && q[len(q)-1] < a {
		q = q[:len(q)-1]
	}
	q = append(q, a)

	return q
}
