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

// TOn SO1

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}