package distance

import (
	"github.com/yuhangch/gots/geom"
	"math"
)

func PointToSegment(C geom.Point, seg geom.LineSegment) float64 {
	/*
		assume distance = |CP|
		line AB, point C, L = |AB|,point P
		|AC|cos(∠CAB)	   AC·AB     |AC|
		-------       =   -------- · ----  =  r
		|AB|			  |AC||AB|	 |AB|

		r represent the ratio of the AC projection on AB to length of AB
		it has the following meaning:
		r = 0: P is A
		r = 1: P is B
		r < 0: P on the backward extension of AB
		r > 0: P on the forward extension of AB
		0 < r <1: P is interior to AB

	*/
	A, B := seg.Endpoint1(), seg.Endpoint2()
	L2 := (B.X-A.X)*(B.X-A.X) + (B.Y-A.Y)*(B.Y-A.Y)
	r := ((C.X-A.X)*(B.X-A.X) + (C.Y-A.Y)*(B.Y-A.Y)) / L2
	if r <= 0.0 {
		return C.Distance(A)
	}
	if r >= 1.0 {
		return C.Distance(B)
	}
	/*
								   CA×AB      CA×AB
		d = |AC|sin(∠CAB) = |AC|· -------- = --------
								  |CA||AB|      L

		ps: (x1,y1)×(x2,y2) =  x1 * y2 - x2 * y1
	*/
	return math.Abs((A.Y-C.Y)*(B.X-A.X)-(A.X-C.X)*(B.Y-A.Y)) / math.Sqrt(L2)

}
