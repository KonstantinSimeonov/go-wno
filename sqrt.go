package main

import (
	"fmt"
	"math"
)

func SqrtNewton(x float64) float64 {
	for guess := 1.0; ; {
		new_guess := guess - (guess*guess-x)/(2*guess)

		if math.Abs(new_guess-guess) < 0.00001 {
			return new_guess
		} else {
			guess = new_guess
		}
	}
}

func SqrtBinarySearch(x float64) float64 {
	left := 0.0
	right := x

	for right-left > 0.00000001 {
		mid := right - (right-left)/2
		if mid*mid > x {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v: Newton(%v), BinarySearch(%v)\n", i, SqrtNewton(float64(i)), SqrtBinarySearch(float64(i)))
	}
}
