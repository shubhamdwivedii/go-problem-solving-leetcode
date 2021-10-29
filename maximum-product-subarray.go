
func max(a, b, c int) int {
	if a > b {
		if c > a {
			return c
		}
		return a
	}
	if c > b {
		return c
	}
	return b
}

func min(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		}
		return a
	}
	if c < b {
		return c
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// For +ve only numbers the product will be ever increasing.
	// we need to keep track of -ve with a Min product in addition to Max

	// -ve * -ve is +ve, if number is -ve, then -ve * Min = New Max, and -ve * Max = New Min

	res := 0

	currMin := 1
	currMax := 1

	for _, n := range nums {
		if n == 0 { // Zero will make product = 0
			currMax = 1
			currMin = 1
			continue
		}

		tempMax := currMax
		currMax = max(currMax*n, currMin*n, n)
		currMin = min(tempMax*n, currMin*n, n)
		res = max(currMax, currMin, res)
	}
	return res
}
