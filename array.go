package main

import (
	"fmt"
)

func main123345() {
	array1 := [...]int{1, 2, 5, 5}
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
	array2 := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array2)
	fmt.Println(x)
	// array3 := [3]float64{7.0, 8.5, 9.1}
	fmt.Println(printArray(14))
	fmt.Println(fibonacci(5))
	arr := fibs(10)
	for i := 0; i < 10; i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, arr[i])
	}
	fmt.Println(testG(3, 3, 1))
}

func Sum(inputArray *[3]float64) (sum float64) {
	for _, v := range *inputArray {
		sum += v
	}
	return sum
}

func printArray(j int) (returnArray []int) {
	for i := 0; i <= j; i++ {
		returnArray = append(returnArray, i)
	}
	return returnArray
}
