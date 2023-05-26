package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// interface
type Shape interface {
	area() float32
}

type CustomError struct {
	Code int
	Err  error
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("status %d: err %v", err.Code, err.Err)
}

// first struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// another struct to implement interface
type Triangle struct {
	base, height float32
}

// Triangle provides implementation for area()
func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// access method of the interface
func calculate(s Shape) {
	fmt.Println("Area:", s.area())
}

// main function
func mainIF() {

	// assigns value to struct members
	// rectangle := Rectangle{7, 4}
	// triangle := Triangle{2, 3}

	// // call calculate() with struct variable rectangle
	// calculate(rectangle)
	// // call calculate() with struct variable triangle
	// calculate(triangle)
	str, err := sqrt(-2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(str)
	}
}

func sqrt(x float64) (string, error) {
	if x < 0 {
		return "0", &CustomError{503, errors.New("unavailable")}
	}
	return fmt.Sprint(math.Sqrt(x)), nil
}
