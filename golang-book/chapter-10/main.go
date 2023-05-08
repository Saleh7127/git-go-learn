package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n, ":", string(rune(i+65)))
		time.Sleep(1 * time.Second)
	}
}

func f1(n int) {
	for i := 0; i < 5; i++ {
		j := i + 65
		c := rune(j)
		fmt.Println(n, ":", c)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go f(0)
	go f1(0)
	go f(0)
	go f1(0)

	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}
*/

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
