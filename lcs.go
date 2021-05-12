package main

import (
	"fmt"
)

func lcs_rec(m, n int, a, b string) int {
	switch {
	case m < 0 || n < 0:
		return 0
	case a[m] == b[n]:
		return 1 + lcs_rec(m - 1, n - 1, a, b)
	default:
		result_a := lcs_rec(m - 1, n, a, b)
		result_b := lcs_rec(m, n - 1, a, b)
		if result_a > result_b {
			return result_a
		}

		return result_b
	}
}

func main() {
	x := "aggtab"
	y := "gxtxayb"

	fmt.Println(x, y, lcs_rec(len(x) - 1, len(y) - 1, x, y))
}
