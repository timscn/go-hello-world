package main

import (
	"fmt"
	userdata "go-hello-world/entities"
	"time"
)

const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 10, 20, 30, 40, 50, 60

func main() {
	// value, err := Hello1()
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(value)
	// }
	const TESTME = "test struct"
	const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
	println(TESTME, TUESDAY)
	var p = userdata.Person{FirstName: "Rajeev", LastName: "Singh", Age: 25}
	println(p.FirstName)
	println(p.LastName)
	println(p.Age)
	// viktor := User{firstName: "etr", lastNmae: "ewrwe"}
	// fmt.Println(viktor)

	// fmt.Println(quote.Go())
	// fmt.Println("before Map")
	// id := 1
	// mapData := helloMap()s
	// value, idForCheck := mapData[id]
	// if idForCheck == true {
	// 	fmt.Println("the is: ", id, "value is: ", value)
	// }
	// delete(mapData, 1)
	// fmt.Println("loop throught the entire mapdata")
	// for key, value := range mapData {

	// 	fmt.Println(key, value)
	// }
	fmt.Println("Finally I'm here :D")
	i1 := 19
	var j1 = i1
	fmt.Println("Printing Reference Addresses ")
	fmt.Printf("Address of i1=%d:\t%p\n", i1, &i1)
	fmt.Printf("Address of j1=%d:\t%p\n", j1, &j1)
	var b bool
	fmt.Println(b)
	// fmt.Println(helloMap())

	fahrenheitValue, err := toFahrenheit(112)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("return of toFahrenheit is: ", fahrenheitValue)
	}
	fmt.Println("Testing Strings")
	myString := "Prefix"
	fmt.Println("Does the string predefined Prefix: ", hasPrefix(myString))
	t := time.Now().UTC()
	fmt.Println(t)
	fmt.Println("Does the string predefined Sufix: ", hasSufix(myString))
	fmt.Println("Pointers one more time")
	address := 10
	fmt.Println(&address)
	addressNew := getAddressOfVar(address)
	fmt.Println(getAddressOfVar(address))
	fmt.Println(addressNew)
	fmt.Println("Testing simpleSwitch")
	fmt.Println(simpleSwitch(2))
	fmt.Println("Testing complicatedSwitch")
	fmt.Println(complicatedSwitch(2))
	fmt.Println("Testing Challenge 2")
	s := seasonCh(12)
	fmt.Println(s)
	fmt.Println("before LoopMe")
	fmt.Println(loopMe(15))
	fmt.Println("after LoopMe")
	fmt.Println("before loopConditionbased")
	fmt.Println("value of sum is: ", sum)
	fmt.Println(loopConditionbased(18))
	fmt.Println("after loopConditionbased")
	fmt.Println("abs")
	fmt.Println(abs(-10))
}
