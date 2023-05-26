package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	go say("hello")
	say("world")
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sumChan(a[:len(a)/2], c)
	go sumChan(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	// wg := new(sync.WaitGroup)
	channelBUffered := make(chan int)
	// channelBUffered <- 1
	// go func() {
	// 	wg.Wait()
	// 	channelBUffered <- 1
	// 	channelBUffered <- 2
	// 	channelBUffered <- 3
	// 	close(channelBUffered)
	// }()
	// msg := <-channelBUffered
	// Read from the messages channel
	// for msg1 := range msg {
	// 	fmt.Println(msg1)
	// }
	channelBUffered = make(chan int)
	go chanBUff(channelBUffered)
	for value := range channelBUffered {
		fmt.Println(value)
	}
	channel1 := make(chan int, 10)
	go fibChannel(channel1, len(channel1))
	for i := range channel1 {
		fmt.Println(i)
	}
	randNumber := rand.Intn(5)
	for true {
		fmt.Println("Guess the random number between 1 & 5")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		clearInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Enter a valid number, %s is not a number\n", input)
			continue
		}
		if randNumber > clearInput {
			fmt.Println("should be bigger than: ", clearInput)
		} else if randNumber < clearInput {
			fmt.Println("should be smaller than: ", clearInput)
		} else if randNumber == clearInput {
			fmt.Println("yup, exactly - good jobs, its: ", clearInput)
			time.Sleep(1 * time.Second)
			fmt.Println("Thank you...bye")
			time.Sleep(1 * time.Second)
			break
		}
	}
	arr := [5]int{1, 2, 4, 5, 6}
	fmt.Println("arr is: ", arr)
	sl1 := make([]int, len(arr))
	fmt.Println("sl1 is: ", sl1)
	sl1 = arr[3:]
	fmt.Println("sl1 after is: ", sl1)
	sl1 = append(sl1, 34432)
	fmt.Println("sl1 after append is: ", sl1)
	val := 10
	fmt.Println("val is: ", val)
	changeMeValue(val)
	fmt.Println("val after is: ", val)
	changeMePtr(&val)
	fmt.Println("val changeMePtr is: ", val)
	fmt.Println("val address is: ", &val)
	fmt.Println("val changeMePtr is: ", val)
}

func changeMePtr(val *int) {
	*val += 10
}

func changeMeValue(val int) int {
	return val + 10
}

func chanBUff(chanB chan int) {
	for i := 0; i < 5; i++ {
		chanB <- i
	}
	close(chanB)
}

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(str)
	}
}

func sumChan(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
