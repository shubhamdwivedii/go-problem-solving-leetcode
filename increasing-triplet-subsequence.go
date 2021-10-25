func increasingTripletLIS(nums []int) bool {
	// longest increasing subsequence

	lis := make([]int, len(nums))
	for i := range lis {
		lis[i]++
	}

	left := 0
	right := 1

	for right < len(nums) {
		for left < right {
			if nums[right] > nums[left] {
				if lis[right] <= lis[left] {
					lis[right] = lis[left] + 1
					if lis[right] >= 3 {
						return true
					}
				}
			}

			left++
		}
		left = 0
		right++
	}

	return false
}

func increasingTriplet(nums []int) bool {
	ni := math.MaxInt64
	nj := math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] < ni {
			ni = nums[i]
			continue
		}

		if nums[i] < nj && nums[i] > ni {
			nj = nums[i]
		}

		if nums[i] > nj {
			return true
		}
	}

	return false
}
