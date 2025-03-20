package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Alumno struct {
	Nombre string
	Notas  []float64
}

func leerCSV(nombreArchivo string) ([]Alumno, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}

	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}

	var alumnos []Alumno

	for _, registro := range registros {
		nombre := registro[0]
		var notas []float64
		for _, notaStr := range registro[1:] {
			nota, err := strconv.ParseFloat(notaStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil
}

func calcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func calcularMediaAritmetica(alumnos []Alumno) float64 {
	var suma float64
	var cantidadNotas int
	for _, alumno := range alumnos {
		for _, nota := range alumno.Notas {
			suma += nota
			cantidadNotas++
		}
	}
	return suma / float64(cantidadNotas)
}

func escribirCSV(nombreArchivo string, alumnos []Alumno) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escritor := csv.NewWriter(archivo)
	defer escritor.Flush()

	for _, alumno := range alumnos {
		promedio := calcularPromedio(alumno.Notas)
		registro := []string{alumno.Nombre, fmt.Sprintf("%.2f", promedio)}
		if err := escritor.Write(registro); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	alumnos, err := leerCSV("notas.csv")
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	mediaAritmetica := calcularMediaAritmetica(alumnos)
	fmt.Printf("La media aritmética de todas las notas es: %.2f\n", mediaAritmetica)

	err = escribirCSV("promedios.csv", alumnos)
	if err != nil {
		fmt.Println("Error al escribir el archivo CSV:", err)
		return
	}
	fmt.Println("Archivo 'promedios.csv' generado con éxito.")

}
