package main

import (
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	boats := 0
	slices.Sort(people)
	p1, p2 := 0, len(people)-1			
	for p1 <= p2{
		if people[p1] + people[p2] <= limit{
			boats++	
			p1++
			p2--
		}else{
			p2--
			boats++
		}
	}
	return boats 
}