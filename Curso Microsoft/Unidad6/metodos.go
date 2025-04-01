package main

import (
	"fmt"
	"geometry"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s geometry.Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())

	var s geometry.Shape = geometry.NewSquare(3)
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	printInformation(s)

	c := Circle{6}
	printInformation(c)
}

/* type coloredTriangle struct {
	triangle
	color string
}

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter (colored)", t.perimeter())
	fmt.Println("Perimeter (normal)", t.triangle.perimeter())
}*/

/*type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}

type square struct {
	size int
}



func (s square) perimeter() int {
	return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())
}
*/
