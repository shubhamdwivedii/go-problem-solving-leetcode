func threeSum(arr []int) [][]int {
	t := 0
	if len(arr) < 3 {
		return [][]int{}
	}

	// make sure array is sorted.
	sort.Ints(arr)
	var triplets [][]int
	lastNum := arr[0]
	for i, num := range arr {
		if i != 0 && num == lastNum {
			continue
		} else {
			lastNum = num
		}
		subarr := arr[i+1:]
		target := t - num

		// Now its simple 2SUM problem
		leftPointer := 0
		rightPointer := len(subarr) - 1

		for leftPointer < rightPointer {
			left := subarr[leftPointer]
			right := subarr[rightPointer]

			if left+right == target {

				triplets = append(triplets, []int{num, left, right})
				// rightPointer--
				for subarr[leftPointer] == left && leftPointer < rightPointer {
					leftPointer++
				}
			} else {
				if left+right < target {
					for subarr[leftPointer] == left && leftPointer < rightPointer {
						leftPointer++
					}
				} else {
					for subarr[rightPointer] == right && leftPointer < rightPointer {
						rightPointer--
					}

				}
			}
		}
	}

	// If this reached means no solution found
	return triplets
}