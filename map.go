package main

import "fmt"

func helloMap() map[int]int {

	mapdata := make(map[int]int)
	mapdata[1] = 1
	mapdata[2] = 4564
	return mapdata
}

func main12dfds3() {
	fmt.Println("4 is: ", findDay(4))
}

func findDay(id int) string {
	map1 := make(map[int]string, 7)
	map1[1] = "Monday"
	map1[2] = "Tuesday"
	map1[3] = "Wednesday"
	map1[7] = "Sunday"
	map1[5] = "Friday"
	map1[6] = "Saturday"
	map1[4] = "Thursday"
	fmt.Println(map1)
	val1, isPresent := map1[id]
	if isPresent {
		return val1
	} else {
		return "Wrong Key"
	}
}
