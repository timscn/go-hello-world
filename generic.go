package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func Hello1() (string, error) {
	fmt.Println("The type is: ", reflect.TypeOf(math.Sqrt(125)))

	fmt.Println("some math is here", math.Sqrt(125))
	var first, second, third int

	var array1 [2]int

	var numbers []int
	fmt.Println("Print slice - origin")

	for i := 0; i < 2; i++ {
		numbers = append(numbers, i+1)
	}
	fmt.Println(numbers)
	fmt.Println("Adding something to slice which has: ", len(numbers))

	fmt.Println("Please enter the third Digit")
	fmt.Scanln(&third)
	numbers = append(numbers, third)
	fmt.Println("Print slice - final", numbers)

	fmt.Println("Please enter the first Digit")
	fmt.Scanln(&first)
	fmt.Println("Please enter the second Digit")
	fmt.Scanln(&second)

	array1[0] = first
	array1[1] = second

	negativeError := errors.New("negativeValue")

	if array1[0] < 0 {
		return "negaive array1[0] - plz change", negativeError
	}
	if array1[1] < 0 {
		return "negaive array1[1] - plz change", negativeError
	}

	if array1[0] > array1[1] {
		return " first is bigger", nil
	} else if array1[0] < array1[1] {
		return "second is bigger", nil
	} else {
		return "same digits - please try again", nil
	}
}
