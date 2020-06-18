package main

import (
	"fmt"
	"testing"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println("")
}

func test() {
	if 1 == 1 {
		return
	}
}

func wode() {
}

// Area to caculate area
func Area(ok int) float64 {
	return float64(ok)
}

// Point to generate point
type Point struct {
}

// Test to test user snippets.
type Test struct {
	name  string
	color float64
}

// TestFunc to test fun snippets.
func TestFunc(name string) string {
	fmt.Printf("%s\n", "")

	return ""
}

// TestCommentMethod to test comment for method.
func TestCommentMethod(name string) string {
	return "ok"
}

// Do to do some thing for test.
func (test *Test) Do() {
	fmt.Println()
}

// Test .
func TestFun2(t *testing.T) {

}

// getHome to ge home.
func getHome() string {
	return "三脚架家"
}
