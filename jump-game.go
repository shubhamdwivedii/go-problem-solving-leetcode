func canJump(nums []int) bool {
	if len(nums) == 1 { // already on end
		return true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= len(nums)-1-i { // can reach end from this index
			if i == 0 {
				return true
				break
			}
			nums = nums[:i+1] // slice nums so that i is the end now.
		}
	}
	return false
}