func searchRange(nums []int, target int) []int {
	ptrS := 0

	for ptrS < len(nums) && nums[ptrS] != target {
		ptrS++
	}

	if ptrS == len(nums) {
		return []int{-1, -1}
	}

	ptrE := ptrS

	for ptrE < len(nums) && nums[ptrE] == target {
		ptrE++
	}
	ptrE--

	return []int{ptrS, ptrE}
}