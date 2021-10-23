func productExceptSelf(nums []int) []int {
	product := 1
	zeroEnc := false
	zeroIdx := -1
	for i, n := range nums {
		if n == 0 {
			if !zeroEnc {
				zeroEnc = true
				zeroIdx = i
				continue
			} else {
				product = 0
				break
			}
		}
		product *= n
	}

	res := make([]int, len(nums))
	if product == 0 {
		return res
	}

	if zeroEnc {
		res[zeroIdx] = product
		return res
	}

	for i := 0; i < len(res); i++ {
		res[i] = int(product / nums[i])
	}
	return res
}