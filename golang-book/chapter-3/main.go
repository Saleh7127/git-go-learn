package main

import "fmt"

/*func main() {
	var x []int
	x = make([]int, 10)
	x = make([]int, 5, 10)
	x = []int{1, 2, 3, 4, 5, 6, 7, 8}
	x = x[0:5]
	s := []int{1, 23, 3, 4, 5, 6, 7, 77}
	s2 := append(s)
	fmt.Println(s, s2)
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
*/

func main() {
	/*x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
	x[2] = 1000
	delete(x, 2)
	fmt.Println(x[2])*/

	//elements := make(map[string]string)
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	name, ok := elements["Un"]
	fmt.Println(name, ok)
	if name, ok := elements["H"]; ok {
		fmt.Println(name, ok)
	}
}
