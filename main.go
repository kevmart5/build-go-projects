package main

import "fmt"

type bot interface {
	getGreeting() string
}

type bot1 struct{}
type bot2 struct{}

// Assignment interfaces

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	botOne := bot1{}
	botTwo := bot2{}

	printGreeting(botOne)
	printGreeting(botTwo)

	triangleFigure := triangle{base: 2, height: 3}
	squareFigure := square{sideLength: 4}

	fmt.Println("")
	fmt.Println("#################")

	printArea(triangleFigure)
	printArea(squareFigure)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (bt1 bot1) getGreeting() string {
	return "Hello bot 1"
}

func (bt2 bot2) getGreeting() string {
	return "Hello bot 2"
}

// #####################################

func printArea(sp shape) {
	fmt.Println(sp.getArea())
}

func (t triangle) getArea() float64 {
	areaTriangle := 0.5 * t.base * t.height
	return areaTriangle
}

func (s square) getArea() float64 {
	squareArea := s.sideLength * s.sideLength
	return squareArea
}
