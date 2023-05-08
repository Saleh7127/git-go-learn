package main

import "fmt"
import "golang-book/chapter-11/math"

func main() {
	xs := []float64{1, 3}
	avg := math.Average(xs)
	fmt.Println(avg)
}
