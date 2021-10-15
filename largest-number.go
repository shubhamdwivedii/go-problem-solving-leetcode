func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		is := fmt.Sprintf("%v", nums[i])
		js := fmt.Sprintf("%v", nums[j])
		return is+js > js+is
	})

	res := ""

	sum := 0
	for _, n := range nums {
		res += fmt.Sprintf("%v", n)
		sum += n
	}

	if sum == 0 {
		return "0"
	}

	return res
}

