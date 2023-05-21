package main

import (
	"fmt"
	userdata "go-hello-world/entities"
	"reflect"
)

func mainInterfaceFuture() {
	std1 := userdata.User{Name: "Vani"}
	fmt.Println(std1)
	std1.FillDefaultUser()
	fmt.Println(std1)
	fmt.Printf("Vani is %d\n", std1.Age)
	std2 := new(userdata.User)
	fmt.Println(std2)
	std3 := userdata.User{}
	fmt.Println(std3)
	fmt.Println(reflect.TypeOf(std3))         // userdata.User
	fmt.Println(reflect.ValueOf(std3).Kind()) // struct
}
