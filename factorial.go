package main

func factorial(input uint64) (fac uint64) {
	if input <= 1 {
		return 1
	} else {
		return factorial(input-1) * input
	}
}
