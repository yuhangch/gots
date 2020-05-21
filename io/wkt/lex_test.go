package wkt

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer("abcdefg")
	for l.nextChar() != 0 {
		fmt.Println(l.index, l.wkt[l.index])
	}

}
