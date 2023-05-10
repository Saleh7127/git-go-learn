/*package main

import "fmt"

func digits(num int, dn chan int) {
	for num != 0 {
		digit := num % 10
		dn <- digit
		num /= 10
	}
	close(dn)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func addAndMultiply(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}

func main() {
	a := 3
	b := 3
	fmt.Println(addAndMultiply(a, b))
}
*/

package main

import "fmt"

func main() {

	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)

	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}
