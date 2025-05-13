/*
Challenge: Write a program that:
1.	Launches 3 goroutines, each printing:
•	"Routine X: Hello!" (where X is 1–3)
2.	The main function waits long enough to let them all finish

Output should look like (order may vary):
Routine 1: Hello!
Routine 2: Hello!
Routine 3: Hello!
Main function done

Use time.Sleep() just for this example — we’ll learn better ways (like sync.WaitGroup) in the next lessons.
*/

package main

import (
  "fmt"
  "time"
)


func hello(r int) {
  fmt.Println("Routine", r,": Hello!")
}

func main()  {

  for _, x := range []int{1,2,3}  {
    go hello(x)
  }

  time.Sleep(1 * time.Second)
  fmt.Println("Main function done")
}
