package main

import (
	"errors"
	"fmt"
	"sort"
)

func testG(a, b, c int) (int, error) {
	negativeError := errors.New("I dont like when negative")
	if a < 0 || b < 0 || c < 0 {
		return -1, negativeError
	}
	overError := errors.New("I dont like when value is more than 9")
	if a > 9 || b > 9 || c > 9 {
		return -1, overError
	}
	if (a > b) && (a > c) {
		fmt.Println("return a: ", a)
		return a, nil
	} else if b > c {
		fmt.Println("return b: ", b)
		return b, nil
	} else {
		fmt.Println("return c: ", c)
		return c, nil
	}
	// edge case is not covered for instance when A equals B
}

func mai456456n() {
	maxValue, err := testG(9, 1, 1)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("return of toFahrenheit is: ", maxValue)
	}
}

func testC(a, b, c, d int) ([]int, error) {
	negativeError := errors.New("I dont like when negative")
	if a < 0 || b < 0 || c < 0 || d < 0 {
		return nil, negativeError
	}
	overError := errors.New("I dont like when value is more than 9")
	if a > 9 || b > 9 || c > 9 || d > 9 {
		return nil, overError
	}
	arrayA := []int{a, b, c, d}
	sort.Ints(arrayA[:])
	return arrayA, nil
	// edge case is not covered for instance when A equals B
}

func main123() {
	arrayTestMe, err := testC(7, 7, 1, 2)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("return array in increasing order: ", arrayTestMe)
	}
}

func testE(a, b, c, d, e int) [5]int {
	arrayE := [5]int{a, b, c, d, e}
	j := len(arrayE)
	for i := 0; i < j-1; i++ {
		if arrayE[i] > arrayE[i+1] {
			tempValue := arrayE[i]
			arrayE[i] = arrayE[i+1]
			arrayE[i+1] = tempValue
			i = -1
		}
	}
	return arrayE
}

func mainOnlyTest() {
	arrayTestMeE := testE(8, 4, 9, 2, 8)
	fmt.Println("return array in order: ", arrayTestMeE)
}
