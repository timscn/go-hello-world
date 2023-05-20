package main

import "fmt"

func main6666() {
	// 1480. Running Sum of 1D Array
	ids := []int{3, 1, 2, 10, 1}
	for i := 1; i < len(ids); i++ {
		ids[i] += ids[i-1]
	}
	fmt.Printf("the total value isnide if Array is: %d\n", ids[len(ids)-1])
}
