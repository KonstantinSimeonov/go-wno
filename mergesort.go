package main

import (
	"fmt"
)

func insertionSort(xs []int, l, r int) {
	for i := 1; i < len(xs); i++ {
		x := xs[i]
		j := 0
		for xs[j] < x && j < i {
			j++
		}

		k := i
		for j < k {
			xs[k] = xs[k - 1]
			k--
		}

		xs[j] = x
	}
}

func merge(xs []int, l, r int) {
	mid := (l + r) / 2
	left := make([]int, mid - l)
	right := make([]int, r - mid)
	copy(left, xs[l:mid])
	copy(right, xs[mid:r])

	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			xs[l] = left[i]
			i++
		} else {
			xs[l] = right[j]
			j++
		}
		l++
	}

	for i < len(left) {
		xs[l] = left[i]
		l++
		i++
	}

	for j < len(right) {
		xs[l] = right[j]
		l++
		j++
	}
}

// in place
func mergesort(xs []int, l, r int) {
	if r - l < 10 {
		insertionSort(xs, l, r)
		return
	}

	mid := (l + r) / 2
	mergesort(xs, l, mid)
	mergesort(xs, mid, r)
	merge(xs, l, r)
}

// not in place
func MergesortSlices(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}

	mid := len(xs) / 2
	left := MergesortSlices(xs[:mid])
	right := MergesortSlices(xs[mid:])

	result := make([]int, len(left) + len(right))
	for i, l, r := 0, 0, 0; i < len(result); i++ {
		switch {
		case l < len(left) && r >= len(right):
			result[i] = left[l]
			l++
		case r < len(right) && l >= len(left):
			result[i] = right[r]
			r++
		case left[l] < right[r]:
			result[i] = left[l]
			l++
		default:
			result[i] = right[r]
			r++
		}
	}

	return result
}

func main() {
	xs := []int{10,9,8,7,11,12,4,3,3, 99, 102, 88, 33, 69, 42, 101, 304, 404, 505, 500}
	mergesort(xs, 0, len(xs))
	fmt.Println(xs)

	xs2 := []int{10,9,8,7,11,12,4,3,3}
	fmt.Println(xs2)
	fmt.Println(MergesortSlices(xs2))
}
