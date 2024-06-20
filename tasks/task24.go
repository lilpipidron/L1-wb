package tasks

import "math"

// По условию координаты инкапсулированы => мы до них напрямуб дотянуться не можем. В go подобное зависит от того
// с какой буквы называется поле, с заглавной, то к полю мы можем спокойно дотянутся, иначе нет
// поэтому чтобы дотянуться нужны геттеры.
// Само решение использует пакет math: извлечение корня и возведение в степень

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Solve24(point1, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point1.GetX()-point2.GetX(), 2) + math.Pow(point1.GetY()-point2.GetY(), 2))
}
