func missingNumber(nums []int) int {
	max := -1
	sum := 0
	zeroFound := false
	for _, n := range nums {
		if n == 0 {
			zeroFound = true
		}
		if n > max {
			max = n
		}
		sum += n
	}
	// sum of 1....n = (n*(n+1))/2

	targetSum := (max * (max + 1)) / 2

	if targetSum == sum {
		if !zeroFound {
			return 0
		}
		return max + 1
	}
	return targetSum - sum
}