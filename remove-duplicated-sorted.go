func removeDuplicates(nums []int) int {
	added := make(map[int]bool)

	var res []int

	for _, n := range nums {
		if _, ok := added[n]; !ok {
			res = append(res, n)
			added[n] = true
		}
	}
	nums = nums[:len(res)]
	copy(nums, res)
	return len(res)
}