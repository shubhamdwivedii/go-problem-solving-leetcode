func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	if len(nums) < 1 {
		return 0
	}

	for _, n := range nums {
		if _, ok := numMap[n]; !ok {
			numMap[n] = true
		}
	}

	max := -1
	for n := range numMap {
		if _, ok := numMap[n-1]; !ok { // n might be start of a sequence
			length := 0
			_, ok := numMap[n+length]
			for ok {
				length++
				_, ok = numMap[n+length]
			}
			if length > max {
				max = length
			}
		}
	}
	return max
}

// TOn SOn


func longestConsecutive(nums []int) int {
    key := make(map[int]bool)

    for _, num := range nums {
        key[num] = true 
    }

    max := 0 
    for k := range key {
        if !key[k-1] {
            count := 1 
            j := k+1 
            for key[j] {
                count++ 
                j++ 
            }
            if count > max {
                max = count
            }
        }
    }
    return max
}