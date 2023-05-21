package main

import (
	"fmt"
	userdata "go-hello-world/entities"
)

func main() {
	std1 := userdata.User{Name: "Vani"}
	fmt.Println(std1)
	std1.FillDefaultUser()
	fmt.Println(std1)
}
