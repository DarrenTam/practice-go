package binary_search

func binarySearch(nums []int, value int) int {
	midpoint := len(nums) / 2
	left := nums[0:midpoint]
	right := nums[midpoint:]

	if len(nums) == 1 && nums[midpoint] != value {
		return -1
	}

	if nums[midpoint] == value {
		return nums[midpoint]
	} else if nums[midpoint] > value {
		return binarySearch(left, value)
	} else {
		return binarySearch(right, value)
	}

}
