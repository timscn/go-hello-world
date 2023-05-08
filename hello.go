package main

import (
	"fmt"
	"math"
	"reflect"

	"rsc.io/quote"
)

func Hello1() string {
	fmt.Println("The type is: ", reflect.TypeOf(math.Sqrt(125)))

	fmt.Println("some math is here", math.Sqrt(125))
	var first, second int

	var array1 [2]int

	var numbers []int
	fmt.Println("Print slice - origin")

	for i := 0; i < 5; i++ {
		numbers = append(numbers, i+1)
	}
	fmt.Println(numbers)
	fmt.Println("Adding somethong to slice which has: ", len(numbers))
	var third int

	fmt.Println("Please enter the third Digit")
	fmt.Scanln(&third)
	numbers = append(numbers, third)
	fmt.Println("Print slice - final", numbers)

	fmt.Println("Please enter the first Digit")
	fmt.Scanln(first)
	fmt.Println("Please enter the second Digit")
	fmt.Scanln(second)

	array1[0] = first
	array1[1] = second

	if array1[0] > array1[1] {
		return " first is bigger"
	} else if array1[0] < array1[1] {
		return "second is bigger"
	} else {
		return "same digits - please try again"
	}
}

func main() {
	fmt.Println(Hello1())
	fmt.Println(quote.Go())
}
