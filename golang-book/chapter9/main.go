package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}
type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
func (c *Circle) area() float64 {
    return math.Pi * math.Pow(c.r, 2)
}
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - y1
    b := y2 - y1
    return math.Sqrt(a * a + b * b)
}/*
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}
func circleArea(c *Circle) float64 {
    return math.Pi * math.Pow(c.r, 2)
}*/
func main() {
    c := Circle{0, 0, 5}
    r := Rectangle{0, 0, 10, 10}

    fmt.Println(r.area())
    fmt.Println(c.area())
}
