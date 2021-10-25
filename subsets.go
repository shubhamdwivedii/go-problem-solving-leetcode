func subsets(nums []int) [][]int {
	// at each step, either add n ([...,n]) or not add n [...]
	// sort.Ints(nums)
	var subsets [][]int
	var curr []int
	add(curr, nums, 0, &subsets)
	return subsets
}

func add(curr []int, nums []int, idx int, subsets *[][]int) {
	if idx >= len(nums) {
		*subsets = append(*subsets, curr)
		return
	}

	add(curr, nums, idx+1, subsets)
	add(append(curr, nums[idx]), nums, idx+1, subsets)
}
