package main

func valueReturn() int  {
  x := 5

  return x
}

func pointerReturn() *int {
  y := 5

  return &y
}

func main()  {

  valueReturn()
  pointerReturn()

}
