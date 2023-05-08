package main

import (
	"fmt"
	"time"
)

/*
func pinger(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func pinga(c <-chan string) {
	for {
		a := <-c
		fmt.Println(a)
		time.Sleep(1 * time.Second)
	}
}*/

/*
	func ponger(c chan string) {
		for i := 0; i < 5; i++ {
			c <- "pong"
		}
	}
*/
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
