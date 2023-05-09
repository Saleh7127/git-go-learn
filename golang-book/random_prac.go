/*package main

import "fmt"
*/
/*func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		u := i
		out = append(out, &u)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}
*/

/*func main() {
	var values []*int
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println(val)
		}(val)
	}
}
*/
/*
func DeduplicateStrings(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}*/

/*func main() {

	in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1} // any item can be sorted
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}
	result := in[:1+j]
	fmt.Println(result)

}

func moveToFront(needle string, haystack []string) []string {

	if len(haystack) != 0 && haystack[0] == needle {
		return haystack
	}
	prev := needle
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}

func main() {

	haystack := []string{"a", "b", "c", "d", "e"} // [a b c d e]
	haystack = moveToFront("c", haystack)         // [c a b d e]
	haystack = moveToFront("f", haystack)         // [f c a b d e]
	fmt.Println(haystack)

}

*/

/*package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
}
*/

/*
package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

func (e Employee) changeName(newName string) {
	e.name = newName
}
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

*/

/*package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
	var cc int
	fmt.Scanf("%d", &cc)
}
*/

/*

package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%c ", i)
	}
}
func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}
*/

package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}
