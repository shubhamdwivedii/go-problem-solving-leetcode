func findDuplicate(nums []int) int {
	encMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := encMap[n]; ok {
			return n
		} else {
			encMap[n] = true
		}
	}
	return -1
}