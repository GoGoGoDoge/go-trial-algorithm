package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := len(nums)
	last := nums[0]
	count := 1

	for _, n := range nums[1:] {
		if n == last {
			if count >= 2 {
				l--
			}
			count++
		} else {
			count = 1
		}
		last = n
	}

	last = nums[0]
	idx := 1
	count = 1
	for _, n := range nums[1:] {
		if n == last {
			if count < 2 {
				nums[idx] = n
				idx++
			}
			count++
		} else {
			nums[idx] = n
			count = 1
			idx++
		}
		last = n
	}

	return l
}
