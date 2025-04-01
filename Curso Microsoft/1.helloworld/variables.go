//formas de declarar variables
var firstname string

var firstname, lastname string
var age int

var (
	firstname, lastname string
	age int
)

//Inicializar variables:
var (
	firstname string = "Carlitos"
	lastname string = "Perez"
	age int = 32
)

var (
	firstname = "Roberto"
	lastname = "Gomez"
	age = 32
)

var (
	firstname, lastname, age = "Carlos", "Lopez", 32
)

package main

import "fmt"

func main() {
	firstname, lastname := "John", "Doe"
	age := 32
	fmt.Println(firstname, lastname, age)
}