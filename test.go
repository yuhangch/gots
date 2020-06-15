package main

import "fmt"

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

func Area(ok int) float64 {
	return float64(ok)
}

type Point struct {
}

type jj struct {
	x float64
}

func (p Point) Area() float64 {
	fmt.Println("wtf")
	return 0

}

type Hello struct {
	name  string
	first string
}

type Haha interface {
}

type Person struct {
	name   string
	age    int
	school string
}

// test commit
