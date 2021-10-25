func rob(nums []int) int {
	maxRob := make([]int, len(nums))
	max := 0

	for i, h := range nums {
		if i < 2 {
			if i < 1 || maxRob[i-1] < h {
				maxRob[i] = h
			} else {
				maxRob[i] = maxRob[i-1]
			}
		} else {
			maxA := maxRob[i-2] + h
			maxB := maxRob[i-1]

			if maxA > maxB {
				maxRob[i] = maxA
			} else {
				maxRob[i] = maxB
			}
		}
		if maxRob[i] > max {
			max = maxRob[i]
		}
	}

	return max

}
