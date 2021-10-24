func sortColors(nums []int) {
	// Simple Selection Sort
	for i := 0; i < len(nums); i++ {

		j := i - 1
		for j >= 0 {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			j--
		}
	}
}