package main

import "fmt"

func main() {
	var edad int = 21
	var permiso bool = false

	if edad < 18 && permiso {
		fmt.Println("El menor de edad puede viajar sólo porque tiene permiso")
	} else if edad < 18 && !permiso {
		fmt.Println("El menor de edad no puede viajar solo")
	} else {
		fmt.Println("La persona puede viajar sólo porque es mayor de edad")
	}
	fmt.Println("Fin de programa")
}
