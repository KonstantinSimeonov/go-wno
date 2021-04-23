package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line1, _ := r.ReadString('\n')
	line2, _ := r.ReadString('\n')

	alice := strings.Split(strings.Trim(line1, "\n"), " ")
	bob := strings.Split(strings.Trim(line2, "\n"), " ")

	var a = 0
	var b = 0
	for i := 0; i < 3; i++ {
		an, _ := strconv.Atoi(alice[i])
		bn, _ := strconv.Atoi(bob[i])
		if an > bn {
			a += 1
		} else if an < bn {
			b += 1
		}
	}

	fmt.Println(a, b)
}
