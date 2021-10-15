
func merge(intervals [][]int) [][]int {
	var merged [][]int

	sort.Slice(intervals, func(i, j int) bool {
		first := intervals[i]
		second := intervals[j]
		return first[0] < second[0]
	})

	current := intervals[0]

	if len(intervals) == 1 {
		return intervals
	}

	for i := 1; i < len(intervals); i++ {
		if current[1] >= intervals[i][0] {
			if current[1] < intervals[i][1] { // current may include interval entirely eg. [1,4], [2,3]
				current[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, current)
			current = intervals[i]

		}

		// at this point either way, current will include the last element.
		if i == len(intervals)-1 {
			merged = append(merged, current)
		}
	}

	return merged
}