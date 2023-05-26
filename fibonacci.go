package main

func fibonacci(input int) (result int) {
	if input <= 1 {
		return 1
		// the above for the first 2 cases
	} else {
		result = fibonacci(input-1) + fibonacci(input-2)
	}
	return
}

func fibs(n int) []int64 {
	if n < 0 {
		return nil
	}
	fibArray := make([]int64, n)

	fibArray[0] = 1 // base case for 1
	fibArray[1] = 1 // base case for 1
	for i := 2; i < n; i++ {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]
	}
	return fibArray
}

func fibChannel(ch chan int, n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- i
		x, y = y, x+y
	}
	close(ch)
}
