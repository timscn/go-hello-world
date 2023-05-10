package main

import (
	"errors"
	"fmt"
)

type Celsius float32
type Fahrenheit float32

const (
	January   = iota
	February  = iota
	March     = iota
	April     = iota
	May       = iota
	June      = iota
	July      = iota
	August    = iota
	September = iota
	October   = iota
	November  = iota
	December  = iota
)

func toFahrenheit(celsius Celsius) (Fahrenheit, error) {
	myError := errors.New("I dont like when zero")
	if celsius == 0 {
		return 0, myError
	}
	return Fahrenheit((celsius * 9 / 5) + 32), nil
}

func seasonCh(month int) string {
	switch month {
	case January, February, December:
		return "Winter"
	case March, April, May:
		return "Spring"
	case June, July, August:
		return "Summer"
	case September, October, November:
		return "Autumm"
	default:
		return "please provide value between 1-12"
	}
}

func abs(x int) int {
	fmt.Println("entering the abs and value is: ", x)
	if x < 0 {
		fmt.Println("entering the if and value is: ", -x)
		return -x
	}
	fmt.Println("outside of if and value is: ", x)
	return x
}
