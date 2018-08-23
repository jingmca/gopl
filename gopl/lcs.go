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
			if X[i] == Y[j] {
				T[i+1][j+1] = T[i][j] + 1
			} else {
				T[i+1][j+1] = max(T[i][j+1], T[i+1][j])
			}
		}
	}
	lcsWordsNoRc(X, T, M, N)
	return T[M][N]
}

func lcsWords(X []rune, table [][]int, M, N int) {
	if M == 0 || N == 0 {
		return
	}

	if table[M][N] == table[M][N-1] {
		lcsWords(X, table, M, N-1)
	} else if table[M][N] == table[M-1][N] {
		lcsWords(X, table, M-1, N)
	} else {
		fmt.Printf("%s\n", string(X[M-1]))
		lcsWords(X, table, M-1, N-1)
	}
}

func lcsWordsNoRc(X []rune, table [][]int, M, N int) {

	k := table[M][N]
	s := make([]rune, k)

	for k > 0 {
		if table[M][N] == table[M-1][N] {
			M--
		} else if table[M][N] == table[M][N-1] {
			N--
		} else {
			//fmt.Printf("%s", string(X[M-1]))
			s[k-1] = X[M-1]
			M--
			N--
			k--
		}

	}
	fmt.Println(string(s))

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
	fmt.Printf("given words: %q, %q\nLCS:\n", string(X), string(Y))

	fmt.Println(lcs(X, Y))
}
