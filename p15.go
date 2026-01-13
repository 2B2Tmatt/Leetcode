package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)

	seenTrip := make(map[[3]int]struct{})

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		pairs := twoSumHelper(nums[i+1:], -nums[i])

		for _, p := range pairs {
			trip := [3]int{nums[i], p[0], p[1]}
			if _, ok := seenTrip[trip]; ok {
				continue
			}
			seenTrip[trip] = struct{}{}
			res = append(res, []int{trip[0], trip[1], trip[2]})
		}
	}

	return res
}

func twoSumHelper(nums []int, target int) [][2]int {
	l, r := 0, len(nums)-1
	res := make([][2]int, 0)

	for l < r {
		sum := nums[l] + nums[r]

		if sum == target {
			res = append(res, [2]int{nums[l], nums[r]})

			l++
			r--

			for l < r && nums[l] == nums[l-1] {
				l++
			}
			for l < r && nums[r] == nums[r+1] {
				r--
			}

		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return res
}
