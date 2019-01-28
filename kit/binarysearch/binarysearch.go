package binarysearch

// Return larger index if not found
func BinarySearchLarge(arr []int, l int, r int, x int) int {
	if l > r || len(arr) == 0 {
		return -1
	}

	for l < r {
		m := l + (r-l)/2
		if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func BinarySearchSmall(arr []int, l int, r int, x int) int {
	if l > r || len(arr) == 0 {
		return -1
	}

	for l < r {
		m := l + (r-l+1)/2
		if arr[m] > x {
			r = m - 1
		} else {
			l = m
		}
	}

	return l
}
