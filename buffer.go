package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main9() {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("ABC")
	fmt.Fprintf(buffer, "testme")
	buffer.WriteString("[DONE]")
	fmt.Println(buffer.String())

	ids := []int{87, 2, 99, 4, 9}
	sort.Ints(ids)
	for _, id := range ids {
		fmt.Println(id)
	}
	fmt.Println("second loop")
	for i, id := range ids {
		fmt.Printf("%d - ID has value: %d\n", i, id)
	}

}
