package main

func fibonacci(input int) (result int) {
	if input <= 1 {
		return 1
	} else {
		result = fibonacci(input-1) + fibonacci(input-2)
	}
	return
}
