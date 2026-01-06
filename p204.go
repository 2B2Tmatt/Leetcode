package main

import "math"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	isPrime := make([]bool, n-2)
	for i := range isPrime {
		isPrime[i] = true
	}
	count := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for j := i; j*i < n; j++ {
				isPrime[j*i-2] = false
			}
		}
	}
	for i := range isPrime {
		if isPrime[i] {
			count++
		}
	}
	return count
}
