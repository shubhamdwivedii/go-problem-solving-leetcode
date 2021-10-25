
func permute(nums []int) [][]int {
	var curr []int
	var output [][]int
	permutate(curr, nums, &output)
	return output
}

func permutate(curr, nums []int, output *[][]int) {
	if len(nums) == 0 {
		return
	}

	if len(nums) == 1 {
		curr = append(curr, nums[0])
		*output = append(*output, curr)
	}

	for i, n := range nums {
		// add ith element to copy of curr
		start := make([]int, len(curr))
		copy(start, curr)
		start = append(start, n)

		// remove ith element from nums
		rest := make([]int, len(nums))
		copy(rest, nums)
		// order is not important
		rest[i] = rest[len(rest)-1]
		rest = rest[:len(rest)-1]

		permutate(start, rest, output)
	}
}
