package gopl

import (
	"fmt"
)

//MaxSum will show
func MaxSum(x []int) int {
	maxsum := -1000000
	sum := 0
	for _, i := range x {
		if sum+i >= 0 {
			sum += i
			if sum >= maxsum {
				maxsum = sum
			}
		} else {
			sum = 0
			if i >= maxsum {
				maxsum = i
			}
		}
	}
	defer fmt.Println(maxsum)
	return maxsum
}
