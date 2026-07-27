func findDuplicate(nums []int) int {
	encMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := encMap[n]; ok {
			return n
		} else {
			encMap[n] = true
		}
	}
	return -1
}

// TON SO1 
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[0]

    // Find intersection inside the cycle
    for {
        slow = nums[slow]
        fast = nums[nums[fast]]

        if slow == fast {
            break
        }
    }

    // Find entrance to the cycle
    slow = nums[0]

    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}