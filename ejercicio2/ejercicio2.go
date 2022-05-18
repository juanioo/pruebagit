//Realizar otra función en donde se pase un array de números y devuelva el mínimo,
//máximo// y la media.//

package ejercicio2

func MinMaxMedia(arreglo []int) (max int, min int, media int) {
	totalNumeros := len(arreglo)
	min = arreglo[0]
	max = arreglo[0]
	var sumaNumeros int
	for i := 0; i < totalNumeros; i++ {
		if max < arreglo[i] {
			max = arreglo[i]
		}
		if min > arreglo[i] {
			min = arreglo[i]
		}
	}
	for i := 0; i < totalNumeros; i++ {
		sumaNumeros = (sumaNumeros + arreglo[i])
		media = int(sumaNumeros / len(arreglo))
	}
	return max, min, media
}
