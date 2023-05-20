package main

import "fmt"

func main888() {
	currentArray := []int{1, 2, 3, 4, 5, 6, 7, 9}
	// currentArray := []int{6, 3, 7}
	// currentArray := [4]int{1, 1, 1, 1}
	// currentArray := []
	newArray := [3]int{}
	if len(currentArray) >= 3 {
		newArray[0] = currentArray[0]
		newArray[1] = currentArray[2]
		newArray[2] = currentArray[len(currentArray)-2]
		fmt.Println(newArray)
	} else {
		fmt.Println("array has less than 3 elements")
	}
}
