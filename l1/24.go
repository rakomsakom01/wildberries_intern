package main


import (
    "fmt"
    "math"
)

/*
Задание номер 24

Разработать программу нахождения расстояния между двумя точками, которые представлены в виде 
структуры Point с инкапсулированными параметрами x,y и конструктором.
*/


// структура
type Point struct {
    x, y float64
}


func (p *Point) X() float64 {
    return p.x
}


func (p *Point) Y() float64 {
    return p.y
}


func NewPoint(x, y float64) Point {
    point := Point{x: x, y: y}
    return point
}

// расстояние между точками
func distance(x, y Point) float64 {
    return math.Sqrt((x.X() - y.X()) * (x.X() - y.X()) + (x.Y() - y.Y()) * (x.Y() - y.Y()))
}



func main() {
    var x1, y1, x2, y2 float64

    fmt.Scanf("%f %f %f %f", &x1, &y1, &x2, &y2)

    var p1 = NewPoint(x1, y1)
    var p2 = NewPoint(x2, y2)

    fmt.Printf("%f\n", distance(p1, p2))
}
