package main

func simpleSwitch(t int) string {
	switch t {
	case 0:
		return "it's zero"
	case 1:
		return "it's one"
	default:
		return "default value"
	}
}

func complicatedSwitch(t int) string {
	switch {
	case t < 0:
		return "negative"
	case t > 0:
		return "positive"
	default:
		return "default value is zero"
	}
}
