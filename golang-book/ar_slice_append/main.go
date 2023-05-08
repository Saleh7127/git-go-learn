package main

import "fmt"

/*


func SubtractOneFromLength(s []int) []int {
	s = s[0 : len(s)-1]
	return s
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 121, 12, 12, 121, 21, 456, 45, 45, 45, 4, 87, 8, 78, 6, 23, 1}
	s := a[1:3]
	s1 := s[0:1]

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(s1)

	a = a[1:10]

	fmt.Println(a)

	fmt.Println("Before: len(slice) =", len(a))
	newSlice := SubtractOneFromLength(a)
	fmt.Println("After:  len(slice) =", len(a))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}
*/
/*
type path []byte

func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' > b && b > 'z' {
			(*p)[i] = b + 32
		}
	}
}

func main() {
	pathName := path("/usr/Bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}
*/

/*func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}*/

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

func main() {
	/*slice := make([]int, 10)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	copy(newSlice, slice)
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))*/

	/*slice := make([]int, 10, 20) // Note capacity > length: room to add element.
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)
	slice = Insert(slice, 5, 99)
	fmt.Println(slice)*/

	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	slice1 := []int{4545, 745, 45, 4}
	slice = Append(slice, slice1...)
	fmt.Println(slice)

}
