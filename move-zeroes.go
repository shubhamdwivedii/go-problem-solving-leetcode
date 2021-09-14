// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Eg. Input: nums = [0,1,0,3,12] Output: [1,3,12,0,0]
func moveZeroes(nums []int) []int {
	ptrL := 0
	ptrR := 1

	if len(nums) < 2 {
		return nums
	}

	for ptrL < ptrR && ptrR < len(nums) {
		if nums[ptrL] == 0 {
			if nums[ptrR] == 0 {
				ptrR++
				continue
			} else {
				// swap
				nums[ptrL], nums[ptrR] = nums[ptrR], nums[ptrL]
				ptrL++
				ptrR++
			}
		} else {
			ptrL++
			ptrR++
		}
	}
	return nums
}