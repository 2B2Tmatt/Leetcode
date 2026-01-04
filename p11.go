package main

func maxArea(height []int) int {
	var max, current int
	i1, i2 := 0, len(height)-1
	for i1 < i2 {
		current = calculateArea(height, i1, i2)
		if current > max {
			max = current
		}

		if height[i1] > height[i2] {
			i2--

		} else {
			i1++
		}
	}
	return max
}

func calculateArea(height []int, i1, i2 int) int {
	var val int
	if height[i2] < height[i1] {
		val = height[i2]
	} else {
		val = height[i1]
	}
	return val * (i2 - i1)
}
