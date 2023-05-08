package main

import "fmt"

/*
	func zero(xPtr *int) {
		*xPtr = 0
	}

	func main() {
		x := 5
		zero(&x)
		fmt.Println(x) // x is 0
	}
*/
/*func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
*/

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}
