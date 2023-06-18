package main

type MyNumber interface {
	int | float64
}

func sumIntsG(m map[string]int) int {
	result := 0
	for _, v := range m {
		result += v
	}
	return result
}

func sumFloatsG(m map[string]float64) float64 {
	result := 0.0
	for _, v := range m {
		result += v
	}
	return result
}

func sumIntsOrFloatsG[K comparable, V int | float64](m map[string]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}

func sumNumbers[K comparable, V MyNumber](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
