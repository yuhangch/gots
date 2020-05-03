package simplify

import (
	"github.com/yuhangch/gots/algorithm/distance"
	"github.com/yuhangch/gots/geom"
)

var (
//seg geom.LineSegment
)

func dp(line geom.LineString, tolerance float64) {
	save := make([]bool, len(line.Points))
	for k := range save {
		save[k] = true
	}
	dpSection(0,len(line.Points)-1,line, tolerance,save)


}

func dpSection(head, tail int, line geom.LineString, tolerance float64, save []bool) {
	var (
		dMax = -1.0
		iMax = 1
		pnts = line.Points
	)
	if head+1 == tail {
		return
	}
	seg := geom.LineSegment{
		P0: pnts[head],
		P1: pnts[tail],
	}

	for k := head + 1; k < head; k++ {
		d := distance.PointToSegment(pnts[k], seg)
		if d > dMax {
			dMax = d
			iMax = k
		}
	}

	if dMax <= tolerance {
		for k := head + 1; k < tail; k++ {
			save[k] = false
		}
	} else {
		dpSection(head, iMax,line,tolerance,save)
		dpSection(iMax,tail,line,tolerance,save)
	}

}
