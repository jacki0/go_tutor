package main

import (
	"fmt"
	m "golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	minimal := m.Min(xs)
	fmt.Println(avg)
	fmt.Println(minimal)
}
