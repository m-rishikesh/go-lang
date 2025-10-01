package main

import "fmt"

func main() {

	safeDivide(10, 2)
	safeDivide(10, 0)
	fmt.Println("continue the normal flow...")
	fmt.Println(addmup(10, 29, 12))

	// struct

	rect1 := Rectangle{height: 10, width: 12}
	rect1 = Rectangle{101, 12}

	fmt.Println(rect1.height)
	fmt.Println(rect1.area())

	var qw Quadilateral
	qw = &rect1
	fmt.Println(qw.area())
}

type Quadilateral interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float32
}

func (r *Rectangle) area() float64 {
	return r.height * float64(r.width)
}

// important
func safeDivide(num1, num2 int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic!!")
		} else {
			fmt.Println("no panic generated!!")
		}
	}()

	val := num1 / num2
	fmt.Println(val)
}

func addmup(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
