package main

import (
	"clases/ejercicio2"
	"clases/ejercicioFibonacci"
	"clases/ejercicios"
	"fmt"
)

func main() {
	min, max := ejercicios.MinMax(12, 13, 54)
	fmt.Printf("Min: %d\nMax: %d\n", min, max)
}

func MinMaxMedia() {
	arreglo := []int{5, 6, 9, 10, 16, 24, 90, 3, 246}
	max, min, media := ejercicio2.MinMaxMedia(arreglo)
	fmt.Printf("Min: %d\nMax: %d\nMedia: %d\n", max, min, media)
}

func Fibonacci() {
	var a, b int
	a, b = ejercicioFibonacci.Fibonacci(a, b)
	fmt.Printf("")
}
