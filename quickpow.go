package main

import (
	"fmt"
	"math"
)

func QuickPow(n, p int) int {
	switch {
	case p <= 1:
		return n
	case p % 2 == 0:
		n2 := QuickPow(n, p / 2)
		return n2 * n2
	default:
		return n * QuickPow(n, p -1)
	}
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Println(i, QuickPow(i, 4), math.Pow(float64(i), 4))
	}
}
