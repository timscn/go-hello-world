package main

import (
	"fmt"
)

func bubbleSort(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func Filter(s []int, fn func(int) bool) []int {
	var result []int // == nil
	for i := 0; i < len(s); i++ {
		if even(s[i]) {
			result = append(result, i)
		}
	}
	return result
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	return ns

}

func insertSlice(slice, insertion []string, index int) (result []string) {
	result = make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func multiDimenstion() {
	slicetToChange := []int{1, 2, 3}
	fmt.Println(cap(slicetToChange))
	newSLice := enlarge(slicetToChange, 5)
	fmt.Println(cap(newSLice))

	array2D := [][]int{{1, 2, 3}, {9, 8, 7}, {1, 4, 2}, {19, 18, 87}, {29, 48, 67}}
	fmt.Printf("type of array is: %T\n", array2D)
	fmt.Println(array2D)
	fmt.Printf("Number of rows in array is: %d\n", len(array2D))
	fmt.Printf("Number of columns in array is: %d\n", len(array2D[2]))
	fmt.Printf("Number of elements in array is: %d\n", len(array2D)*len(array2D[0]))
	fmt.Println("Traversing Array")
	// for _, row := range array2D {
	// 	for _, value := range row {
	// 		fmt.Printf("printing value %d from row number: %d\n", value, row)
	// 	}
	// }
	for i := 0; i < len(array2D); i++ {
		fmt.Println("row number is: ", i+1)
		for j := 0; j < len(array2D[i]); j++ {
			fmt.Printf("printing value %d and row number is: %d\n", array2D[i][j], array2D[i])
		}
	}
}

func mainArray2() {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	bubbleSort(sla)
	fmt.Println("after sort:  ", sla)
	// sliceTest := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(Filter(sliceTest, even))
	// multiDimenstion()
	// array1 := [...]int{1, 2, 5, 5}
	// for i := 0; i < len(array1); i++ {
	// 	fmt.Println(array1[i])
	// }
	// array2 := [3]float64{7.0, 8.5, 9.1}
	// x := Sum(&array2)
	// fmt.Println(x)
	// // array3 := [3]float64{7.0, 8.5, 9.1}
	// fmt.Println(printArray(14))
	// fmt.Println(fibonacci(5))
	// arr := fibs(10)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("The %d-th Fibonacci number is: %d\n", i, arr[i])
	// }
	// fmt.Println(testG(3, 3, 1))
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
