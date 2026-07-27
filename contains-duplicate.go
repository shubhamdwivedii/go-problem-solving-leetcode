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

// T On, S On 
// If they ask "Can you do it with O(1) extra space?", you can sort the array and check adjacent elements, which is O(n log n) time and potentially O(1) auxiliary space depending on the sorting implementation.