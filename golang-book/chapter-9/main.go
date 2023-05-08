package main

import "fmt"

/*type circle struct {
	x, y, z int
}

func CA(c circle) float64 {
	return math.Pi * float64(c.z*c.z)
}

func main() {

	//c := new(circle)

	c := circle{100, 1121, 12}

	c.x = 121
	c.y = -5 * -4

	fmt.Println(CA(c))
}
*/

/*func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
*/

/*type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Talk()
}
*/

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(c))
	fmt.Printf("Rectangle area: %f\n", getArea(r))
}
