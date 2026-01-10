package main

import (
	"slices"
)

type numVal struct {
	val int
	idx int
}

func twoSum(nums []int, target int) []int {
	newNums := make([]numVal, 0)
	for i, num := range nums {
		newNums = append(newNums, numVal{val: num, idx: i})
	}
	slices.SortFunc(newNums, func(a, b numVal) int {
		return a.val - b.val
	})
	answers := make([]int, 2)
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 {
		if newNums[p1].val+newNums[p2].val == target {
			answers[0] = newNums[p1].idx
			answers[1] = newNums[p2].idx
			return answers
		} else if newNums[p1].val+newNums[p2].val > target {
			p2--
		} else {
			p1++
		}
	}
	return answers
}
