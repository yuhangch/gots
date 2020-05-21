package simplify

import (
	"github.com/yuhangch/gots/algorithm/distance"
	"github.com/yuhangch/gots/geom"
)

type DpCfg struct {
	save      *[]bool
	line      *geom.LineString
	tolerance float64
}

func DP(line geom.LineString, tolerance float64) geom.LineString {
	save := make([]bool, len(line))
	for k := range save {
		save[k] = true
	}
	cfg := &DpCfg{
		save:      &save,
		line:      &line,
		tolerance: tolerance,
	}

	dpSection(0, len(line)-1, cfg)

	var points []geom.Point
	for k, v := range save {
		if v == true {
			points = append(points, line[k])
		}
	}
	return points

}

func dpSection(head, tail int, cfg *DpCfg) {
	var (
		dMax = -1.0
		iMax = 1
		pnts = cfg.line
	)
	if head+1 == tail {
		return
	}
	seg := geom.LineSegment{
		(*pnts)[head],
		(*pnts)[tail],
	}

	for k := head + 1; k < tail; k++ {
		d := distance.PointToSegment((*pnts)[k], seg)
		if d > dMax {
			dMax = d
			iMax = k
		}
	}

	if dMax <= cfg.tolerance {
		for k := head + 1; k < tail; k++ {
			(*cfg.save)[k] = false
		}
	} else {
		dpSection(head, iMax, cfg)
		dpSection(iMax, tail, cfg)
	}

}
