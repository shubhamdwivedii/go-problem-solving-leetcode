func majorityElement(nums []int) int {
	halfLen := len(nums) / 2

	countMap := make(map[int]int)

	for _, n := range nums {
		if c, ok := countMap[n]; ok {
			if c+1 > halfLen {
				return n
			}
			countMap[n]++
		} else {
			if halfLen < 1 {
				return n
			}
			countMap[n] = 1
		}
	}
	return -1
}