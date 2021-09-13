func generate(numRows int) [][]int {
	var res [][]int
	res = append(res, []int{1})

	curr := []int{1, 1}
	n := 1

	for n < numRows {
		res = append(res, curr)
		newCurr := make([]int, len(curr)+1)
		newCurr[0] = 1
		newCurr[len(newCurr)-1] = 1
		midPoint := len(newCurr)/2 + len(newCurr)%2
		for i := 1; i < midPoint; i++ {
			newCurr[i] = curr[i] + curr[i-1]
			newCurr[len(newCurr)-1-i] = newCurr[i]
		}
		curr = newCurr
		n++
	}

	return res
}