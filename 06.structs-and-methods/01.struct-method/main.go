package main

import "fmt"

type Rectangle struct {
  Width, Height float64
}

func (rectangle Rectangle) Area() float64 {
  return rectangle.Height * rectangle.Width
}

func main()  {

  rectangle := Rectangle{ Width: 2, Height: 4 }

  area := rectangle.Area()

  fmt.Println("The area is:" , area)
}
