package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Given vertex is: ", v)
	fmt.Println("Absolute value of given vertex is: ", Abs(v))
	f := 10.0
	fmt.Printf("After scaling by %0.1f new vertex is: ", f)
	Scale(&v, f)
	fmt.Println(v)
	fmt.Println("Absolute value of given vertex is: ", Abs(v))
}
