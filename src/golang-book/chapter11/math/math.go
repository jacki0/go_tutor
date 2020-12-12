package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Найти меньшее в срезе чисел
func Min(xs []float64) float64 {
	min := float64(1000 * 1000)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

//Найти найбольшее в срезе чисел
func Max(xs []float64) float64 {
	max := float64(-1000 * 4110)
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}
