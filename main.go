package main

import (
	solutions "dsa-go/kata"
	"fmt"
)

func main() {
	var stack solutions.Stack

	// fmt.Println(stack.HammingWeight(50))

	// fmt.Println(stack.NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	// fmt.Println(stack.NextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))

	fmt.Println(stack.NextGreaterElements([]int{5, 4, 3, 2, 1}))
	fmt.Println(stack.NextGreaterElements([]int{1, 2, 1}))
	fmt.Println(stack.NextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(stack.NextGreaterElements([]int{1, 2, 3, 2, 1}))
}
