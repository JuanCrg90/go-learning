/*
Challenge: Write a program that:

Defines an interface Shape with a method Area() float64.
Defines two structs: Circle and Rectangle.
Implements the Area() method for both.
Writes a function printArea(s Shape) that prints the area.
*/

package main

import (
  "fmt"
  "math"
)

type Shape interface {
  Area() float64
}

type Circle struct {
  Radius float64
}

type Rectangle struct {
  Width float64
  Height float64
}

func (c Circle) Area() float64 {
  return math.Pow((math.Pi * c.Radius), 2)
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func printArea(s Shape) {
  fmt.Println(s.Area())
}

func main() {
  var circle Shape = Circle{Radius: 3}
  var rectangle Shape = Rectangle{Width: 5, Height: 8}

  printArea(circle)
  printArea(rectangle)
}
