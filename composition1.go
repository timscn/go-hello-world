package main

import "fmt"

type Contact struct {
	name  string
	phone string
}

type BusinessContact struct {
	c          Contact
	workAdress string
}

func (b BusinessContact) info() {
	fmt.Println("Contact is at: ", b.c.name, b.c.phone, b.workAdress)
}
