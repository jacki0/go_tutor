package main

import "fmt"
/*
func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func main() {
    xs := []float64{98, 93, 77, 82, 83}
    fmt.Println(average(xs))
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1, 2, 3))
}

func main() {
    xs := []int{1, 2, 3}
    fmt.Println(add(xs...))
}

func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1, 1))
}

func main() {
    x := 0
    increment := func() int {
        x++
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
}

func maleEvenGenerator() func() uint {
    i :=  uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}


func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x - 1)
}

func main() {
    factorial(2)
}


func first() {
    fmt.Println("1st")
}
func second() {
    fmt.Println("2nd")
}
func main() {
    defer second()
    first()
}*/


func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
