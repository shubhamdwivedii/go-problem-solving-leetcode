func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := 0
	ptr2 := 0

	res := make([]int, m+n)

	idx := 0
	for idx < m+n {
		if ptr1 >= m && ptr2 < n {
			res[idx] = nums2[ptr2]
			ptr2++
		} else if ptr2 >= n && ptr1 < m {
			res[idx] = nums1[ptr1]
			ptr1++
		} else if ptr1 < m && ptr2 < n {
			if nums1[ptr1] <= nums2[ptr2] {
				res[idx] = nums1[ptr1]
				ptr1++
			} else {
				res[idx] = nums2[ptr2]
				ptr2++
			}
		}
		idx++
	}

	copy(nums1, res)
}