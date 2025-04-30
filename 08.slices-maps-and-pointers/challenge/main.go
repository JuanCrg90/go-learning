/*
- Challenge: Write a program that:
- Creates a map[string]int representing a simple inventory ("apple": 5, "banana": 2).
- Write a function updateInventory that receives the inventory map and an item name, and increments the item count by 1.
- If the item doesn’t exist, add it with a count of 1.
- Test updating "apple", and adding a new "orange".


Example output:

map[apple:6 banana:2 orange:1]

*/


package main

import "fmt"

func updateInventory(inventory map[string]int, itemName string) {
  inventory[itemName] += 1
}

func main() {
  inventory := map[string]int {
    "apple": 5,
    "banana": 2,
  }

  fmt.Println(inventory)
  updateInventory(inventory, "apple")
  fmt.Println(inventory)
  updateInventory(inventory, "orange")
  fmt.Println(inventory)
}

