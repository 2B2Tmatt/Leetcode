package main

func isBadVersion(version int) bool {
	return false
}
func firstBadVersion(n int) int {
	p1, p2 := 1, n
	for p1 <= p2 {
		mid := (p1 + p2) / 2
		if isBadVersion(mid) {
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			}
			p2 = mid - 1
		} else {
			p1 = mid + 1
		}
	}
	return p1
}
