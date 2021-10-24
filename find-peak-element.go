func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	idx := 1

	if nums[0] > nums[idx] {
		return 0
	}

	for idx < len(nums)-1 {
		if nums[idx-1] < nums[idx] && nums[idx] > nums[idx+1] {
			return idx
		}
		idx++
	}

	return len(nums) - 1

}