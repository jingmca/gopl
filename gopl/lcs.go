package gopl

import (
	"fmt"
)

func lcs(X []rune, Y []rune) int {
	M, N := len(X), len(Y)
	T := make([][]int, M+1)

	for i := range T {
		T[i] = make([]int, N+1)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			// fmt.Printf("current: %c,%c\n", X[i], Y[j])
			// fmt.Printf("T: (%d,%d) = %d, (%d,%d) = %d\n", i, j, T[i][j], i+1, j+1, T[i+1][j+1])
			if X[i] == Y[j] {
				T[i+1][j+1] = T[i][j] + 1
			} else {
				T[i+1][j+1] = max(T[i][j+1], T[i+1][j])
			}
		}
	}

	return T[M][N]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//LCS wikl show
func LCS(X, Y []rune) {
	fmt.Println(lcs(X, Y))
}
