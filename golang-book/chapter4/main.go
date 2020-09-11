package main

import "fmt"

func main() {
    var F float64
    fmt.Scanf("%f", &F)
    C := (F - 32) * 5 / 9
    fmt.Println(C)
}
