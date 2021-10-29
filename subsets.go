func subsets(nums []int) [][]int {
	var output [][]int
	var curr []int
	addNum(0, nums, curr, &output)
	return output
}

func addNum(idx int, nums, curr []int, output *[][]int) {
	if idx > len(nums)-1 {
		*output = append(*output, curr)
		return
	}

	addNum(idx+1, nums, curr, output)
	newCurr := make([]int, len(curr)+1)
	copy(newCurr, curr)
	// newCurr = append(newCurr, nums[idx])
	newCurr[len(newCurr)-1] = nums[idx]
	addNum(idx+1, nums, newCurr, output)
}
