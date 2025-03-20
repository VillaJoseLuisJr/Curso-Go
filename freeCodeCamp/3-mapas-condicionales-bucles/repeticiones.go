package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		var suma int = 0
		for i := 0; i <= 100; i++ {
			if i%2 != 0 {
				suma = suma + i
			}
		}
		fmt.Println("La suma de los primeros 100 numeros es: ", suma)
	*/

	/*
		miMapa := map[string]string{
			"Colombia":  "Bogota",
			"Brasil":    "Brasilia",
			"Argentina": "Ciudad Autónoma de Buenos Aires",
			"Chile":     "Santiago",
		}

		for k, v := range miMapa {
			fmt.Println("La capital de " + k + " es: " + v)
		}*/

	var fruta string = "pera"

	for {
		fmt.Println("Indique su fruta de preferencia: ")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)

		if fruta == "naranja" {
			fmt.Println("El usuario finalizó escogiendo naranja")
			break
		}
	}

}
