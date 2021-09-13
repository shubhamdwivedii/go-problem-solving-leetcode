func containsDuplicate(nums []int) bool {
	encMap := make(map[int]bool)

	for _, n := range nums {
		if _, ok := encMap[n]; ok {
			return true
		} else {
			encMap[n] = true
		}
	}
	return false
}