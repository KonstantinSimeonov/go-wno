package main

import (
	"fmt"
)

func knapsack(capacity int, costs, weights []int) int {
	if (len(costs) <= 0 || capacity <= 0) {
		return 0
	}

	without_item := knapsack(capacity, costs[1:], weights[1:])

	if weights[0] > capacity {
		fmt.Printf("can't take (%dc, %dw)\n", costs[0], weights[0])
		return without_item
	}

	with_item := costs[0] + knapsack(capacity - weights[0], costs[1:], weights[1:])

	if (with_item > without_item) {
		fmt.Printf("taking (%dc, %dw)\n", costs[0], weights[0])
		return with_item
	}

	fmt.Printf("skipping (%dc, %dw)\n", costs[0], weights[0])
	return without_item
}

func main() {
	cs := []int{5, 3, 10, 2}
	ws := []int{10, 8, 20, 4}

	result := knapsack(23, cs, ws)

	fmt.Println(result)
}
