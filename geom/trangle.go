package geom

import "math"

type Triangle [3]Point

func NewTrangle(a, b, c Point) Triangle {
	return [3]Point{a, b, c}
}

func (triangle Triangle) Area() float64 {
	A, B, C := triangle.A(), triangle.B(), triangle.C()
	return math.Abs(((C.X()-A.X())*(B.Y()-A.Y()) - (B.X()-A.X())*(C.Y()-A.Y())) / 2)
}

func (triangle Triangle) A() Point {
	return triangle[0]
}

func (triangle Triangle) B() Point {
	return triangle[1]

}

func (triangle Triangle) C() Point {
	return triangle[2]

}
