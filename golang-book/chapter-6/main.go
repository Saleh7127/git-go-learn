package main

import (
	"fmt"
)

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
	/*elements := map[string]string{
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
	}*/

	/*x := map[int]map[int]map[int]int{

		1: map[int]map[int]int{
			2: map[int]int{
				3:  4,
				10: 44,
				12: 545,
			},
			55: map[int]int{
				35:  4,
				13:  41,
				110: 414,
				112: 5145,
			},
		},
		5: map[int]map[int]int{
			6: map[int]int{
				7: 8,
			},
		},
	}*/

	//fmt.Println(x[1][2][3])

	/*if el, ok := x[1]; ok {
		for _, j := range el {
			fmt.Println(j)
		}
	}*/

	/*for key1, val1 := range x {
		if key1 == 1 {
			fmt.Printf("Key 1: %d\n", key1)
			for key2, val2 := range val1 {
				if key2 == 2 {
					fmt.Printf("\tKey 2: %d\n", key2)
					for key3, val3 := range val2 {
						if key3 == 10 {
							fmt.Printf("\t\tKey 3: %d, Value: %d\n", key3, val3)
						}
					}
				}
			}
		}
	}*/

	u := []int{1, 2, 4, 1, 5, 4, 100, 1, 11, 1, 1, 1, 1, 1, 521, 521, 521, 21, 21, 3, 145, 48, 7, 2, 3, 3, 2}
	mx := u[0]
	for i := 1; i < len(u); i++ {
		if u[i] > mx {
			mx = u[i]
		}
	}
	fmt.Println(mx)
}
