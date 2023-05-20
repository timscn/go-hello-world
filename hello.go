package main

import (
	"fmt"
	userdata "go-hello-world/entities"
	"go-hello-world/strutil"
	"log"
	"runtime"
	"time"

	"rsc.io/quote"
)

const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 10, 20, 30, 40, 50, 60

func mainOldOne() {
	const TESTME = "test struct"
	const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
	println(TESTME, TUESDAY)
	var p = userdata.Person{FirstName: "Rajeev", LastName: "Singh", Age: 25}
	println(p.FirstName)
	println(p.LastName)
	println(p.Age)
	fmt.Println(quote.Go())
	fmt.Println("before Map")
	id := 1
	mapData := helloMap()
	value, idForCheck := mapData[id]
	if idForCheck {
		fmt.Println("the is: ", id, "value is: ", value)
	}
	delete(mapData, 1)
	fmt.Println("loop throught the entire mapdata")
	for key, value := range mapData {
		fmt.Println(key, value)
	}
	fmt.Println("Finally I'm here :D")
	i1 := 19
	var j1 = i1
	fmt.Println("Printing Reference Addresses ")
	fmt.Printf("Address of i1=%d:\t%p\n", i1, &i1)
	fmt.Printf("Address of j1=%d:\t%p\n", j1, &j1)
	var b bool
	fmt.Printf("%T\n", b)
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
	fmt.Println("Testing pkg strUtils")
	fmt.Println(strutil.Reverse("olleh"))
	fmt.Println("Im working on FizzBuss")
	checkFizzBuss()
	fmt.Println("Testing pkg challengeF")
	f1, f2, f3 := twoum(3, 4)
	fmt.Println(f1, f2, f3)
	fmt.Println(sumInts(5, -2, 0, 9))
	fmt.Println("Testing pkg challengeF - Fib sequence")
	result := 0
	for i := 0; i <= 4; i++ {
		result = fibonacci(i)
		fmt.Println("So Fib number is: ", result)
	}
	fmt.Println("Factorial is: ", factorial(13))
	lFunction := lambdaTest(10)
	fmt.Println("Test Lambda function", lFunction(10))

	fmt.Println("Debugging with Closures")
	log.SetFlags(log.Llongfile)
	var where = log.Print
	a := 2 * 2
	where()
	fmt.Println(a)

	where12 := func() { // debugging function
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	b12 := 2 * 2
	where12()
	fmt.Println(b12)

	fmt.Println("Timing a function")
	start := time.Now()
	Calculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Calculation took this amount of time: %s\n", delta)
}
