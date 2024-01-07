package components

type CircleType struct {
	radius float64
}

type RectangleType struct {
	length float64
	width  float64
}

type ShapeType interface {
	Area() float64
	Perimeter() float64
}
