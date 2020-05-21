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

func NewParser(l *Lexer) *Parser {
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
			fmt.Println(geom)
		case "LINESTRING":
			geom = p.handleParen(p.parseLineString)
			fmt.Println(geom)
		}

	}
	return
}

func (p *Parser) handleParen(parse func() geom.Geometry) (coordinates geom.Geometry) {
	if p.match(LEFT_PAREN) {
		coordinates = parse()
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

func (p *Parser) parseLineString() geom.Geometry {
	var line geom.LineString
	line = append(line, p.parsePoint().(geom.Point))
	for p.match(COMMA) {
		line = append(line, p.parsePoint().(geom.Point))
	}

	return line
}
