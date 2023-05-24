package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main12354() {
	GetResult(users)
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func coinsPerUser(userName string) int {
	total := 0
	for i := 0; i < len(userName); i++ {
		switch strings.ToLower(string(userName[i])) {
		case "a":
			total++
		case "e":
			total++
		case "i":
			total = total + 2
		case "o":
			total = total + 3
		case "u":
			total = total + 4
		}
	}
	return total
}

func GetResult([]string) string {
	for _, value := range users {
		finalCount := coinsPerUser(value)
		if finalCount > 10 {
			finalCount = 10
		}
		distribution[value] = finalCount
		coins = coins - finalCount
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
	return strconv.Itoa(distribution["Matthew"]) + " " + strconv.Itoa(distribution["Sarah"]) + " " + strconv.Itoa(distribution["Augustus"]) + " " + strconv.Itoa(distribution["Heidi"]) + " " + strconv.Itoa(distribution["Emilie"]) + " " + strconv.Itoa(distribution["Peter"]) + " " + strconv.Itoa(distribution["Giana"]) + " " + strconv.Itoa(distribution["Adriano"]) + " " + strconv.Itoa(distribution["Aaron"]) + " " + strconv.Itoa(distribution["Elizabeth"])
}
