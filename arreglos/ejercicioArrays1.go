//Dado un arreglo del tipo de datos int, analizar cuales son números pares y cuales son los impares
// y almacenarlos en dos arreglos distintos. Luego imprimir los tres arreglos.

package arreglos

func ArregloParImpar(arreglo []int) ([]int, []int) {
	var conteoPares, conteoImpares int
	conteoPares = 0
	conteoImpares = 0
	for i := 0; i < 9; i++ {
		if arreglo[i]%2 == 0 {
			conteoPares++
		} else {
			conteoImpares++
		}
	}
	arregloPar := make([]int, conteoPares)
	arregloImpar := make([]int, conteoImpares)
	conteoPares = 0
	conteoImpares = 0
	for i := 0; i < 9; i++ {
		if arreglo[i]%2 == 0 {
			arregloPar[conteoPares] = arreglo[i]
			conteoPares++
		} else {
			arregloImpar[conteoImpares] = arreglo[i]
			conteoImpares++
		}
	}
	return arregloPar, arregloImpar
}
