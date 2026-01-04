package main

func searchRange(nums []int, target int) []int {
	targets := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				targets[0] = mid
				break
			}
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				targets[1] = mid
				break
			}
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return targets
}
