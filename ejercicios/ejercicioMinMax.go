package ejercicios

func MinMax(a, b, c int) (int, int) {
	var min, max int
	if a > b || a > c {
		max = a
	} else if b > c {
		max = b
	} else {
		max = c
		if b > a {
			min = a
		} else {
			min = b
		}
	}
	return min, max
}
