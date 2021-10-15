func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	// First Find the Pivot element (index) using Binary Search.

	low := 0
	high := len(nums) - 1
	pivot := -1

	for low <= high { // if low == high ie: no pivot
		mid := (low + high) / 2

		if nums[mid] == target { // just in case
			return mid
		}
		if mid == 0 {
			// either of first two elements is pivot
			if nums[mid] > nums[mid+1] {
				pivot = mid
			} else {
				pivot = 1
			}
			break
		}

		if low == high {
			// (2*low)/2 == low (len(arr))
			if low > 0 && low < len(nums)-1 {
				if nums[low] > nums[low-1] && nums[low] > nums[low+1] {
					pivot = low
				}
			}
			break // no pivot
		}

		// Pivot element i should be greater than both i-1 and i+1 elements
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			pivot = mid
			break
		}
		// mid is not pivot

		if nums[mid] > nums[0] {
			// nums is strictly increasing upto mid, pivot is after mid.
			low = mid + 1
			// first half can be discarded
		} else {
			// pivot is before mid.
			high = mid - 1
			// second half can be discarded
		}
	}

	if pivot == -1 {
		// No Pivot
		low = 0
		high = len(nums) - 1
	} else if nums[pivot] == target { // just in case
		return pivot
	} else {
		// check if target before pivot, its possible that target is at num[0], hence ">="
		if target >= nums[0] && target < nums[pivot] {
			// discard second half after pivot
			low = 0
			high = pivot - 1
		} else {
			// target is after pivot
			low = pivot + 1
			high = len(nums) - 1
			// discard first half
		}
	}

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// target before mid, discard first half
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low >= len(nums) || nums[low] != target {
		return -1
	}

	return low
}