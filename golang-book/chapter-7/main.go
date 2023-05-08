/*package main

import (
	"fmt"
)

/*func solve(u []int) int {
	mx := u[0]
	for i := 1; i < len(u); i++ {
		if u[i] > mx {
			mx = u[i]
		}
	}
	return mx
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {

	var u uint
	u = 5

	fmt.Println(factorial(u))
}
*/
/*
package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func th() {
	fmt.Println("3rd")
}

func fr() {
	fmt.Println("4th")
}
func main() {
	defer th()
	defer fr()
	first()
	defer second()
}
*/

package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
