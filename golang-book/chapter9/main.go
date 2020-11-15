package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}
type Rectangle struct {
    x1, y1, x2, y2 float64
}
type Shape interface {
    area() float64
    perimeter() float64
}
type MultiShape struct {
    shapes []Shape
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
}
func (m *MultiShape) area() float64 {
    var area float64
    for _, s:= range m.shapes {
        area += s.area()
    }
    return area
}
/*
func main() {
    c := Circle{0, 0, 5}
    r := Rectangle{0, 0, 10, 10}

    fmt.Println(r.area())
    fmt.Println(c.area())
}
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("hi, my name is", p.Name)
}
type Android struct {
    Person
    Model string
}

func main() {
    a := new(Android)
    a.Talk()
}*/
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

}
func main() {
    c := Circle{0, 0, 5}
    r := Rectangle{0, 0, 10, 10}

    fmt.Println(r.area())
    fmt.Println(c.area())
    fmt.Println(totalArea(&c, &r))
    fmt.Println(perimeter(&r))
}

