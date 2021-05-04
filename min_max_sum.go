package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	nums := make([]int, 5)

	for i, strnum := range strings.Split(strings.Trim(line, "\n"), " ") {
		n, _ := strconv.Atoi(strnum)
		nums[i] = n
	}

	sort.Ints(nums[:])

	minsum := 0
	maxsum := 0
	for i := 0; i < 4; i++ {
		minsum += nums[i]
		maxsum += nums[4 - i]
	}

	fmt.Printf("%d %d\n", minsum, maxsum)
}
