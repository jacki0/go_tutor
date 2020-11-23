package main

import "fmt"

func main() {
    var feets float64
    fmt.Scanf("%f", &feets)
    meters := feets * 0.3048
    fmt.Println(meters)
}
