func lengthOfLIS(nums []int) int {
	left := 0
	right := 1

	// right goes from 1 to len(nums)-1
	// each time right increments, left is reset to 0.

	lst := make([]int, len(nums))

	// lst should be initialized to 1. longest substring from 0 to index i
	for i := range lst {
		lst[i] = 1
	}

	for right < len(nums) {
		for left < right {
			if nums[left] < nums[right] {
				if lst[left] >= lst[right] { // initially all are 1
					lst[right] = lst[left] + 1 // increase by 1 since right > left, length of subsequence is increased by one.
				}
			}
			left++ // check every element from 0 upto right
		}
		// left has reached right
		left = 0
		right++ // check max subsequence length upto new index right+1.
	}

	fmt.Println(lst)

	max := 1
	for _, n := range lst {
		if max < n {
			max = n
		}
	}

	return max
}