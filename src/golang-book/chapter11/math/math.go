package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Найти меньшее в срезе чисел
func Min(xs []float64) float64 {
	min := float64(1000 * 1000)
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}
