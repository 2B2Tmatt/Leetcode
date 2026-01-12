package main

func twoSum(nums []int, target int) []int {
	answers := make([]int, 2)
	seen := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if idx, exists := seen[need]; exists {
			answers[0] = i
			answers[1] = idx
			break
		} else {
			seen[num] = i
		}
	}
	return answers
}
