func maxSlidingWindow(nums []int, k int) []int {
	o, q := make([]int, 0), make([]int, 0)

	for i := 0; i < k; i++ {
		q = addElement(nums, i, q, k)
	}

	o = append(o, nums[q[0]])

	for i := k; i < len(nums); i++ {
		q = addElement(nums, i, q, k)
		o = append(o, nums[q[0]])
	}
	return o
}

func addElement(nums []int, index int, q []int, k int) []int {
	if len(q) > 0 && index-q[0] == k {
		q = q[1:len(q)]
	}

	for i := len(q) - 1; i >= 0; i-- {
		if nums[index] >= nums[q[i]] {
			q = q[:len(q)-1]
			continue
		}
		break
	}

	return append(q, index)
}