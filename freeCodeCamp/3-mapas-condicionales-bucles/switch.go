package main

import "fmt"

func main() {
	var fruta string

	fmt.Print("Ingresa un valor alfanumerico: ")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Println("Has ingresado 'manzana'.")
	case "banana":
		fmt.Println("Has ingresado 'banana'.")
	case "pera":
		fmt.Println("Has ingresado 'pera'.")
	default:
		fmt.Println("No reconozco ese valor.")
	}
}
