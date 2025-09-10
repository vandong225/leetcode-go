package binarysearch

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}

	return left
}
