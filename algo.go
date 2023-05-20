package main

import (
	"fmt"
	"math"
)

var initial string = "hackerhappy"
var desired string = "hackerrank"
var operations int = 9

func main01131() {
	// fmt.Println(saveThePrisoner())
	fmt.Println(appendAndDelete(initial, desired, operations))
}

func saveThePrisoner() (warmPrisoner int) {
	prisoners := 7
	sweets := 5
	startWith := 2

	warmPrisoner = (startWith + sweets - 1) % prisoners
	if warmPrisoner == 0 {
		return prisoners
	} else {
		return warmPrisoner
	}
}

func appendAndDelete(s, t string, k int) (result string) {
	if len(s)+len(t) <= int(k) {
		return "Yes"
	}
	min := math.Min(float64(len(s)), float64(len(t)))
	length := 0
	for i := 0; i < int(min); i++ {
		if s[i] == t[i] {
			length++
		} else {
			break
		}
	}
	total := (len(s) - length) + (len(t) - length)
	if total <= int(k) && (total-int(k))%2 == 0 {
		return "Yes"
	}
	return "No"
}
