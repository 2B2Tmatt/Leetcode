package main

func missingNumber(nums []int) int {
	sum, expected := 0, len(nums)*(len(nums)+1)/2
	for _, num := range nums {
		sum += num
	}
	return expected - sum
}
