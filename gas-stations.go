func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		start := i + 1
		fuel := gas[start-1] - cost[start-1]
		for start != i {
			if start > len(gas)-1 {
				start = 0
				continue
			}

			// check fuel before driving
			if fuel < 0 {
				break // start from next city i++
			}

			fuel += gas[start] - cost[start]
			// driving towards next city
			start++
		}

		// looped through
		if fuel < 0 {
			continue
		} else {
			return i
		}
	}
	return -1
}