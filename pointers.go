package main

import "fmt"

func getAddressOfVar(digit int) *int {
	return &digit
}

func main77777() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n ", a)
	fmt.Printf("%T\n ", b)
	fmt.Println("Value of B Address is: ", *b)
	// Change value with Pointers
	*b += a
	fmt.Println("The pre-final Value of B is:  ", *b)
	a = 20
	fmt.Println("The final Value of B is:  ", *b)
}
