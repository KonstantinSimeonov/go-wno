package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func between(a, b, c int) bool {
	return a <= b && b <= c
}

// https://www.hackerrank.com/challenges/almost-sorted/problem
func main() {
	reader := bufio.NewReader(os.Stdin)

	// first line is redundant
	reader.ReadString('\n')
	line, _ := reader.ReadString('\n')
	// go sucks for competitve programming
	str_nums := strings.Split(strings.Trim(line, "\n"), " ")
	nums := make([]int, len(str_nums)+2)

	// i hate edge cases so here we are
	nums[0] = -1
	nums[len(nums)-1] = 1000001

	for i, x := range str_nums {
		n, _ := strconv.Atoi(x)
		nums[i+1] = n
	}

	// indices of the starts of monotonous sequences
	xtrs := make([]int, 0)
	// cheeky
	decreasing := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] != decreasing {
			decreasing = !decreasing
			xtrs = append(xtrs, i)
		}

		// questionable optimization
		if len(xtrs) > 4 {
			fmt.Println("no")
			return
		}
	}

	// there are two points in the array which are extremas
	// -1 10 2 3 4 1 20 1000001 - 10 and 20
	if len(xtrs) == 4 && xtrs[1]-xtrs[0] == 1 && xtrs[3]-xtrs[2] == 1 {
		fmt.Printf("yes\nswap %d %d\n", xtrs[0], xtrs[2]+1)
		return
	}

	// a decresing monotonous sequence exists
	// -1 1 2 3 10 9 8 7 20 21 1000001 - 10 and 20
	// -1 2 1 1000001 - 2 to 1 - use swap
	if len(xtrs) == 2 && between(nums[xtrs[0]-1], nums[xtrs[1]], nums[xtrs[0]+1]) && between(nums[xtrs[1]-1], nums[xtrs[0]], nums[xtrs[1]+1]) {
		// no ternary op makes me sad
		if xtrs[1]-xtrs[0] == 1 {
			fmt.Printf("yes\nswap %d %d\n", xtrs[0], xtrs[1])
		} else {
			fmt.Printf("yes\nreverse %d %d\n", xtrs[0], xtrs[1])
		}
		return
	}

	fmt.Println("no")
}
