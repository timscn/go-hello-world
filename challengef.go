package main

import "fmt"

func twoum(f1 int, f2 int) (sum int, product int, diff int) {
	sum = f1 + f2
	product = f1 * f2
	diff = f1 - f2
	fmt.Println(sum, product, diff)
	return sum, product, diff
}

func sumInts(list ...int) (sum int) {

	for i := 0; i < len(list); i++ {
		sum = sum + list[i]
	}
	defer twoum(3, 4)
	return sum
}

func lambdaTest(input int) func(int) int {
	fmt.Println("input is: ", input)
	input++
	fmt.Println("input++ is: ", input)
	return func(i int) int {
		fmt.Println("inside of func and sum is ", (input + i))
		return (input + i)
	}
}

func Calculation() {
	for i := 0; i < 10000; i++ {
		//do something
	}
}
