package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	line, _ := reader.ReadString('\n')
	var sum = int64(0)
	for _, v := range strings.Split(strings.Trim(line, "\n"), " ") {
		n, _ := strconv.Atoi(v)
		sum += int64(n)
	}

	fmt.Println(sum)
}
