package main

import "fmt"

func main() {
	miMapa := map[string]string{
		"Colombia":  "Bogota",
		"Brasil":    "Brasilia",
		"Argentina": "Ciudad Autónoma de Buenos Aires",
		"Chile":     "Santiago",
	}

	fmt.Println("La capital de Brasil es:", miMapa["Brasil"])

	miMapa["Uruguay"] = "Montevideo"
	fmt.Println("El mapa de paises es:", miMapa)

	delete(miMapa, "Colombia")
	fmt.Println("El mapa de paises es:", miMapa)

}
