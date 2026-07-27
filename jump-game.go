func canJump(nums []int) bool {
	if len(nums) == 1 { // already on end
		return true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= len(nums)-1-i { // can reach end from this index
			if i == 0 {
				return true
				break
			}
			nums = nums[:i+1] // slice nums so that i is the end now.
		}
	}
	return false
}

// TOn SO1 

func canJump(nums []int) bool {
    goal := len(nums) - 1

    for i := len(nums) - 2; i >= 0; i-- {
        if i + nums[i] >= goal {
            goal = i
        }
    }

    return goal == 0
}

// Goal starts at index 4

// index 3: nums[3] = 1 > can reach goal > new goal = 3
// index 2: nums[2] = 1 > can reach goal > new goal = 2
// index 1: nums[1] = 3 > can reach goal > new goal = 1
// index 0: nums[0] = 2 > can reach goal > new goal = 0

// goal == 0 > true