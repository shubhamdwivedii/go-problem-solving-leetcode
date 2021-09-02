package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	smap := make(map[string]bool)
	for len(nums) > 0 {
		n := nums[0]
		nims := nums[1:]
		found := true

		for found {
			if len(nims) < 2 {
				break
			}
			fmt.Println("n", n, "nims", nims)
			ss, nn, ff := findSum(n, nims)
			found = ff
			if ff {
				fmt.Println("founddd", ss, nn)
				str := fmt.Sprintf("%v%v%v", ss[0], ss[1], ss[2])

				if _, ok := smap[str]; !ok {
					res = append(res, ss)
					str2 := fmt.Sprintf("%v%v%v", ss[0], ss[2], ss[1])
					str3 := fmt.Sprintf("%v%v%v", ss[1], ss[0], ss[2])
					str4 := fmt.Sprintf("%v%v%v", ss[1], ss[2], ss[0])
					str5 := fmt.Sprintf("%v%v%v", ss[2], ss[1], ss[0])
					str6 := fmt.Sprintf("%v%v%v", ss[2], ss[0], ss[1])
					smap[str] = true
					smap[str2] = true
					smap[str3] = true
					smap[str4] = true
					smap[str5] = true
					smap[str6] = true
				}
			}
			nims = nn
		}
		nums = nums[1:]
	}
	return res
}

func findSum(s int, nums []int) ([]int, []int, bool) {
	for j := 0; j < len(nums)-1; j++ {
		for k := j + 1; k < len(nums); k++ {
			if nums[j]+nums[k] == -s {
				cpy := make([]int, len(nums))
				copy(cpy, nums)
				var nim []int
				fmt.Println("found", nums, j, k)
				nim = append(nim, cpy[:j]...)
				nim = append(nim, cpy[j+1:k]...)
				nim = append(nim, cpy[k+1:]...)
				fmt.Println("nim", nim)
				return []int{s, nums[j], nums[k]}, nim, true
			}
		}
	}
	return nil, nil, false
}

func main() {
	fmt.Println("RES:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("RES:", threeSum([]int{0, 0, 0, 0}))
	fmt.Println("RES:", threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println("RES:", threeSum([]int{3, 0, -2, -1, 1, 2}))
	fmt.Println("RES:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("RES:", threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
