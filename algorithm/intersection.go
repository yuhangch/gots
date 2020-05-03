package algorithm

import (
	"github.com/yuhangch/gots/geom"
	"math"
)

func intersection(pnt1, pnt2, pnt3, pnt4 geom.Coordinate) geom.Coordinate {
	left1 := math.Min(pnt1.X, pnt2.X)
	right1 := math.Max(pnt1.X, pnt2.X)
	top1 := math.Max(pnt1.Y, pnt2.Y)
	bottom1 := math.Min(pnt1.Y, pnt2.Y)

	left2 := math.Min(pnt3.X, pnt4.X)
	right2 := math.Max(pnt3.X, pnt4.X)
	top2 := math.Max(pnt3.Y, pnt4.Y)
	bottom2 := math.Min(pnt3.Y, pnt4.Y)

	left := math.Max(left1, left2)
	right := math.Min(right1, right2)
	top := math.Min(top1, top2)
	bottom := math.Max(bottom1, bottom2)

	centX, centY := (right+left)/2.0, (top+bottom)/2.0

	p1x, p1y := pnt1.X-centX, pnt1.Y-centY
	p2x, p2y := pnt2.X-centX, pnt2.Y-centY
	q1x, q1y := pnt3.X-centX, pnt3.Y-centY
	q2x, q2y := pnt4.X-centX, pnt4.Y-centY

	//p1 (p1x,p1y) p1' (p1x,p1y,1)
	//p2 (p2x,p2y) p2' (p2x,p2y,1)
	//p1' X p2' -> ((p1y-p2y),-(p1x-p2x),(p1x*p2y-p1y*p2x))
	l1x, l1y, l1w := p1y-p2y, -(p1x - p2x), p1x*p2y-p1y*p2x
	l2x, l2y, l2w := q1y-q2y, -(q1x - q2x), q1x*q2y-q1y*q2x

	//ip = l1 X l2
	//
	x, y, w := l1y*l2w-l2y*l1w,-(l1x*l2w-l2x*l1w),l1x*l2y-l2x*l1y

	targetX,targetY := x/w,y/w

	if math.IsNaN(targetX) || math.IsInf(targetX,1) || math.IsInf(targetX,-1)||
		math.IsNaN(targetY) || math.IsInf(targetY,1) || math.IsInf(targetY,-1) {
		return geom.Coordinate{X:0,Y:0}
	}

	return geom.Coordinate{X:x+centX,Y:y+centY}

}
