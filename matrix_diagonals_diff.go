package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNums(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	elems := strings.Split(strings.Trim(line, "\n"), " ")
	var nums = make([]int, len(elems))
	for i := 0; i < len(elems); i++ {
		ni, _ := strconv.Atoi(elems[i])
		nums[i] = ni
	}

	return nums
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := ReadNums(reader)[0]

	var diff = 0
	for i := 0; i < n; i++ {
		r := ReadNums(reader)
		diff += r[i] - r[n-i-1]
	}

	if diff < 0 {
		diff = -diff
	}

	fmt.Println(diff)
}
