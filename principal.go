package main

import (
	"clases/arreglos"
	"clases/ejercicio2"
	"clases/ejercicioFibonacci"
	"clases/ejercicios"
	"fmt"
)

func main() {

	Fibonacci()

}

func MinMax() {
	min, max := ejercicios.MinMax(12, 13, 54)
	fmt.Printf("Min: %d\nMax: %d\n", min, max)
}

func MinMaxMedia() {
	arreglo := []int{5, 6, 9, 10, 16, 24, 90, 3, 246}
	max, min, media := ejercicio2.MinMaxMedia(arreglo)
	fmt.Printf("Min: %d\nMax: %d\nMedia: %d\n", max, min, media)
}

func Fibonacci() {
	ejercicioFibonacci.Fibonacci(100)
}

func ArregloParImpar() {
	var conteoPares, conteoImpares int
	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arregloPar, arregloImpar := arreglos.ArregloParImpar(arreglo)
	for i := 0; i < conteoPares; i++ {
		fmt.Print(arregloPar[i], " - ")
	}
	fmt.Printf("Arreglo pares: %d\n", arregloPar[conteoPares])
	fmt.Printf("")
	for i := 0; i < conteoImpares; i++ {
		fmt.Print(arregloImpar[i], " - ")
	}
	fmt.Printf("Arreglo impares: %d\n", arregloImpar[conteoImpares])

}

func MaxNumDeArray() {
	arreglo := []int{5, 6, 9, 10, 16, 24, 90, 3, 246}
	numMayor := arreglos.MaxNumDeArray(arreglo)
	fmt.Printf("El número mayor del array es: %d\n", numMayor)
}
