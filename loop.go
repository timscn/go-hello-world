package main

import "fmt"

var sum int

func loopMe(iteration int) int {
	fmt.Println("before the loop sum is: ", sum)
	for i := 0; i < iteration; i++ {
		// fmt.Println("before sum+ is: ", sum)
		sum = i
	}
	return sum
}

func loopConditionbased(iteration int) int {

	var i, sum int
	for i < iteration {
		i++
		fmt.Println("First sum is: ", sum)
		if sum == 5 {
			continue
		}
		if sum >= 5 {
			break
		}
		sum++
		fmt.Println("At the end sum is: ", sum)
	}
	return sum
}
