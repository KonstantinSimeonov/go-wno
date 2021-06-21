package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lcs_rec(m, n int, a, b string) int {
	switch {
	case m < 0 || n < 0:
		return 0
	case a[m] == b[n]:
		return 1 + lcs_rec(m-1, n-1, a, b)
	default:
		result_a := lcs_rec(m-1, n, a, b)
		result_b := lcs_rec(m, n-1, a, b)
		if result_a > result_b {
			return result_a
		}

		return result_b
	}
}

func max_int(x, y int) int {
	// branchless goes brrrrrr http://web.archive.org/web/20130821015554/http://bob.allegronetwork.com/prog/tricks.html
	x -= y
	x &= (-x) >> 31
	x += y
	return x
	//if x > y {
	//	return x
	//}

	//return y
}

func lcs(a, b string) [][]int {
	table := make([][]int, len(a)+1)
	for i, _ := range table {
		table[i] = make([]int, len(b)+1)
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max_int(table[i-1][j], table[i][j-1])
			}
		}
	}

	return table
}

// rotates 2 rows instead of using a whole table
// uses 2 * len(b) memory, but doesn't allow backtracking
func lcs_rotate(a, b string) int {
	table := make([][]int, 2)

	for i, _ := range table {
		table[i] = make([]int, len(b)+1)
	}

	i := 1
	for ; i <= len(a); i++ {
		ix := i % 2
		prev_row := (i - 1) % 2
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				table[ix][j] = table[prev_row][j-1] + 1
			} else {
				table[ix][j] = max_int(table[prev_row][j], table[ix][j-1])
			}
		}
	}

	return table[(i-1)%2][len(b)]
}

func backtrack_lcs(a, b string, table [][]int) string {
	result := ""
	for i, j := len(a), len(b); i > 0 && j > 0; {
		if a[i-1] == b[j-1] {
			result = string(a[i-1]) + result
			i--
			j--
			continue
		}

		if table[i-1][j] >= table[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return result
}

func debug_table(a, b string, table [][]int) {
	fmt.Print("    ")
	for _, v := range b {
		fmt.Printf(" %c", v)
	}
	fmt.Println()
	for i, r := range table {
		if i > 0 {
			fmt.Printf("%c ", a[i-1])
		} else {
			fmt.Print("  ")
		}
		fmt.Println(r)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)

	l1, _ := r.ReadString('\n')
	l2, _ := r.ReadString('\n')

	x := strings.Trim(l1, "\n")
	y := strings.Trim(l2, "\n")

	// table := lcs(x, y)
	// ans := backtrack_lcs(x, y, table)

	// debug_table(x, y, table)
	// fmt.Println(table[len(x)][len(y)])

	fmt.Println(lcs_rotate(x, y))
}
