package main

import (
	"fmt"

	"rsc.io/quote"
)

func Hello() string {
	return "Hello, world :D"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(quote.Go())
}
