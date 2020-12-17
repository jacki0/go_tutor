package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	// Парсинг
	flag.Parse()
	//  Генерация числа от 0 до max
	fmt.Println(rand.Intn(*maxp))
}