package wkt

import (
	"fmt"
	"github.com/yuhangch/gots/geom"
	"strconv"
)

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(wkt string) *Parser {
	l := NewLexer(wkt)
	return &Parser{
		lexer: l,
		token: l.nextToken(),
	}
}

func (p *Parser) consume() {
	p.token = p.lexer.nextToken()
}

func (p *Parser) isTokenType(tokenType TokenType) bool {
	return p.token.token == tokenType
}

func (p *Parser) match(tokenType TokenType) bool {
	isMatch := p.isTokenType(tokenType)
	if isMatch {
		p.consume()
	}
	return isMatch
}

func (p *Parser) parseGeometryType() (geom geom.Geometry) {
	token := p.token
	if p.match(TEXT) {
		geomType := token.value

		switch geomType {
		case "POINT":
			geom = p.handleParen(p.parsePoint)
		case "LINESTRING":
			geom = p.handleParen(p.parseLineString)
		case "POLYGON":
			geom = p.handleParen(p.parsePolygon)
		case "MULTIPOINT":
			geom = p.handleParen(p.parseMultiPoint)

		}

	}
	fmt.Println(geom)

	return
}

func (p *Parser) handleParen(parse func() geom.Geometry) (geom geom.Geometry) {
	if p.match(LEFT_PAREN) {
		geom = parse()
		if p.match(RIGHT_PAREN) {
			return
		} else {
			panic("error parse geom")
		}
	}
	return
}

func (t Token) tryParseDigit() float64 {
	value := t.value
	c, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("error parse")
		return 0
	}
	return c

}

func (p *Parser) parsePoint() geom.Geometry {
	var digits []float64
	for i := 0; i < 2; i++ {
		token := p.token
		if p.match(NUMBER) {
			digits = append(digits, token.tryParseDigit())
		}
	}

	return geom.NewCoordinate(digits)
}

func (p *Parser) parsePointList(t geom.GeometryType) geom.Geometry {
	var points []geom.Point
	points = append(points, p.parsePoint().(geom.Point))
	for p.match(COMMA) {
		points = append(points, p.parsePoint().(geom.Point))
	}
	switch t {
	case geom.LINE_STRING:
		return geom.NewLineString(points)
	case geom.LINEAR_RING:
		ring, err := geom.NewLinearRing(points)
		if err != nil {
			panic("ring 解析错误")
		}
		return ring
	case geom.MULTI_POINT:
		return geom.NewMultiPoint(points)

	}
	return nil

}

func (p *Parser) parseLineString() geom.Geometry {
	var line geom.LineString
	line = append(line, p.parsePoint().(geom.Point))
	for p.match(COMMA) {
		line = append(line, p.parsePoint().(geom.Point))
	}

	return line
}
func (p *Parser) parseLinearRing() geom.Geometry {
	var ring geom.LinearRing
	ring = append(ring, p.parsePoint().(geom.Point))
	for p.match(COMMA) {
		ring = append(ring, p.parsePoint().(geom.Point))
	}
	if !ring.Valid() {
		panic("数据出错")
	}
	return ring
}

func (p *Parser) parsePolygon() geom.Geometry {
	var pg geom.Polygon
	pg = append(pg, p.handleParen(p.parseLinearRing).(geom.LinearRing))
	for p.match(COMMA) {
		pg = append(pg, p.handleParen(p.parseLinearRing).(geom.LinearRing))
	}
	return pg
}

func (p *Parser) parseMultiPoint() geom.Geometry {
	var mp geom.MultiPoint
	if p.isTokenType(LEFT_PAREN) {
		mp = append(mp, p.handleParen(p.parsePoint).(geom.Point))
		for p.match(COMMA) {
			mp = append(mp, p.handleParen(p.parsePoint).(geom.Point))
		}
	} else {
		return p.parsePointList(geom.MULTI_POINT)
	}
	return mp
}
