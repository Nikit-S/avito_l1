package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

//конструктор под разное количество аргументов
func NewPoint(i ...float64) *Point {

	if len(i) == 2 {
		return &Point{
			x: i[0],
			y: i[1]}
	}
	return &Point{}

}

//подсчет расстояния по формуле
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint()
	fmt.Println(p1.DistanceTo(p2))
}
