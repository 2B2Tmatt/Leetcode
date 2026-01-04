package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 1
	hasInc := false
	for i < len(arr) && arr[i] > arr[i-1] {
		i++
		hasInc = true
	}
	if i == len(arr) {
		return false
	}
	if arr[i] == arr[i-1] || !hasInc {
		return false
	}
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			return false
		}
		i++
	}
	return true
}