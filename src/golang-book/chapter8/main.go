package main

import "fmt"

/*
func zero(xPtr *int) {
    *xPtr = 0
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}


func one(xPtr *int) {
    *xPtr = 1
}
func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
}
*/

func swap(x int, y int) {
    x, y = y, x
    fmt.Println("x =", x, "y =", y)
}
func main() {
    x := 1
    y := 2
    fmt.Println("x =", x, "y =", y)
    swap(x, y)
}
