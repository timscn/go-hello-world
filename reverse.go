package main

import (
	"fmt"
	"math"
)

func mainR() {
	originalStr := "Google"
	reverseStr := make([]byte, len(originalStr))
	iteratorForReverse := 0
	for iteratorForOriginal := len(originalStr) - 1; iteratorForOriginal >= 0; iteratorForOriginal-- {
		reverseStr[iteratorForReverse] = originalStr[iteratorForOriginal]
		iteratorForReverse++
	}
	fmt.Println(string(reverseStr))
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("Length and capacity of b: %d and %d\n ", len(b[1:4]), cap(b[1:4]))
	fmt.Println("len is", string(b[:]))
	fmt.Println("cap is", string(b[:]))
	fmt.Println("len is", string(b[:2]))
	fmt.Println("cap is", string(b[:2]))
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Println("math.MiginInt32 is: ", math.MinInt32)
	fmt.Println("math.MaxInt32 is: ", math.MaxInt32)
}
