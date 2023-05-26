package main

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Id             int
	Name, Location string
}

func (u User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	u      User
	GameId int
}

func GreetingsForPlayer(p Player) string {
	str := p.u.Greetings()
	return str + strconv.Itoa(p.GameId)
}

func main1239999() {
	p := Player{}
	p.u.Id = 1
	p.u.Name = "V"
	p.u.Location = "Chicago"
	p.GameId = 9999
	fmt.Println(p.u.Greetings())
	ids := []int{2, 4, 5, 1, 8, 10, 43, 44}
	for _, v := range ids {
		if v%2 == 0 {
			continue
		}
		fmt.Println(v)
	}
	switch time.Now().Weekday() {
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Today is Weekday :D")
	}
}
