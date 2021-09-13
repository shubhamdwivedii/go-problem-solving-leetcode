func maxSubArray(arr []int) int {
	maxCurr := []int{arr[0]}
	maxGlobal := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		num := arr[i]
		maxCurr = maxArr([]int{num}, append(maxCurr, num))
		// max of [current] and [prevMax] + [current]
		if sumArr(maxCurr) > sumArr(maxGlobal) {
			maxGlobal = maxCurr
		}
	}
	return sumArr(maxGlobal)
}

func maxArr(a []int, b []int) []int {
	sumA := sumArr(a)
	sumB := sumArr(b)
	if sumA >= sumB {
		return a
	}
	return b
}

func sumArr(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}