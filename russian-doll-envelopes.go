func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		first := envelopes[i]
		second := envelopes[j]

		if first[0] < second[0] { // sort according to increasing widths.
			// if first[1] <= second[1] {
			return true
			// }
		}

		if first[0] == second[0] {
			if first[1] > second[1] { // if widths equal, sort in decreasing order of heights
				return true
				// eg. [4,8], [4,5] // correct order (only this way LIS will work)
			}
		}

		return false
	})

	// now apply longest-increasing-subsequence to heights.

	lis := make([]int, len(envelopes))

	for i := range lis {
		lis[i] = 1
	}

	left := 0
	right := 1

	for right < len(envelopes) {
		for left < right {
			if envelopes[left][1] < envelopes[right][1] { // comparing heights
				if lis[right] <= lis[left] {
					lis[right] = lis[left] + 1
				}
			}
			left++
		}
		// left has reached right
		left = 0
		right++
	}

	max := 1
	for _, n := range lis {
		if max < n {
			max = n
		}
	}

	return max
}