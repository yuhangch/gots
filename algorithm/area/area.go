package area

import (
	"github.com/yuhangch/gots/geom"
	"math"
)

func OfRing(ring geom.LinearRing) float64 {
	if !ring.Valid() {
		return 0.0
	}
	sum := 0.0
	for i := 0; i < len(ring); i++ {
		var ip int
		if i+1 > len(ring)-1 {
			ip = 0
		} else {
			ip = i + 1
		}
		sum += (ring[i].X() * ring[ip].Y()) - (ring[i].Y() * ring[ip].X())
		//fmt.Println(ring[i].X() , ring[ip].Y(),ring[i].Y() , ring[ip].X())

	}
	return math.Abs(sum / 2)
}
