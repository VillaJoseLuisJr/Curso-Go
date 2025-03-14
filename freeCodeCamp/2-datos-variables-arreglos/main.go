package main

import "fmt"

func main() {
	var nombrePersona string = "Leonardo"
	var apellidoPersona = "Tortuga"
	segundoNombre := "Pepe"

	fmt.Println("Hola mundo, soy " + nombrePersona)
	fmt.Println("Mi apellido es " + apellidoPersona)
	fmt.Println("Mi segundo nombre es " + segundoNombre)

	//Parte numérica
	var añoActual int16 = 2025
	var añoReducido int8 = 25
	edad := 27
	fmt.Println("El año actual es", añoActual)
	fmt.Println("El año abreviado es", añoReducido)
	fmt.Println("La edad es", edad)

	//Arreglos
	var listaFrutas = [4]string{"Pera", "Naranja"}
	fmt.Println(listaFrutas[1])

	//listaPaises := [3]string{"Argentina", "Brasil", "Peru"}
	listaPaises := []string{"Argentina", "Brasil", "Peru"}
	fmt.Println(listaPaises)
	listaPaises = append(listaPaises, "Chile")
	fmt.Println(listaPaises)

	listaPaises2 := listaPaises[1:3]
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[1:]
	fmt.Println(listaPaises3)

	listaPaises4 := listaPaises[:3]
	fmt.Println(listaPaises4)
}
