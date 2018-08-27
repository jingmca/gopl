package gopl

import (
	"fmt"
)

// LSS will show
func LSS(X []int, K int) int {
	minlen := -1

	for i, v := range X {
		if v >= K {
			minlen = 1
			break
		} else {
			sum := 0
			for j, len := i, 0; j >= 0; j-- {
				len++
				sum += X[j]
				if sum >= K {
					if (minlen == -1) || (len <= minlen) {
						minlen = len
					}
				}
			}
		}
	}

	fmt.Println(minlen)
	return minlen
}
