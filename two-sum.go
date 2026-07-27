// TOn2 TO1

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 1}

}


// More Optimal TOn SOn

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums {
        diff := target - num 
        if j, ok := seen[diff]; ok {
            return []int{i, j}
        }
        seen[num] = i 
    }

    return []int{0,0}
}