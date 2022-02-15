func maxSubArrayLen(nums []int, k int) int {
	ps := make([]int, len(nums))

	ps[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ps[i] = ps[i-1] + nums[i]
	}

	ml := 0

	hm := make(map[int]int)

	hm[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			if ml <= 0 {
				ml = 1
			}
		}

		if v, ok := hm[ps[i]-k]; ok {
			if i-v > ml {
				ml = i - v
			}
		}
		if _, ok := hm[ps[i]]; !ok && ps[i] != 0 {
			hm[ps[i]] = i
		}
	}

	return ml
}