func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := merge(nums1, nums2)
	l := len(merged)
	merge(nums1, nums2)

	fmt.Println(merged)

	if l%2 == 0 {
		m1 := float64(merged[(l/2)-1])
		m2 := float64(merged[l/2])
		if m1 == 0 && m2 == 0 {
			return 0
		}
		return (m1 + m2) / 2
	}
	return float64(merged[l/2])
}

func merge(nums1 []int, nums2 []int) []int {
	var merged []int

	ptr1 := 0
	ptr2 := 0

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] < nums2[ptr2] {
			merged = append(merged, nums1[ptr1])
			ptr1++
		} else {
			merged = append(merged, nums2[ptr2])
			ptr2++
		}
	}

	for ptr1 < len(nums1) {
		merged = append(merged, nums1[ptr1])
		ptr1++
	}

	for ptr2 < len(nums2) {
		merged = append(merged, nums2[ptr2])
		ptr2++
	}

	return merged
}
