package components

import "math"

func (circle CircleType) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (rectangle RectangleType) Area() float64 {
	return (rectangle.length * rectangle.width)
}

func (circle CircleType) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (rectangle RectangleType) Perimeter() float64 {
	return 2 * (rectangle.length + rectangle.width)
}
