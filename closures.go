package main

import "fmt"

func main090909() {
	sum := closure()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}

func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
