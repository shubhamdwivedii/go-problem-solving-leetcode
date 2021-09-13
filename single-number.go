func singleNumber(nums []int) int {
	encMap := make(map[int]int)

	for _, n := range nums {
		if t, ok := encMap[n]; !ok {
			encMap[n] = 1
		} else {
			if t <= 1 {
				encMap[n]++
			}
		}
	}

	for n, times := range encMap {
		if times < 2 {
			return n
		}
	}

	return nums[0]
}