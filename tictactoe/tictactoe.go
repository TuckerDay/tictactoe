package main

import (
  "fmt"
  //"log"

  "example.com/render"
)

func main() {
  // Create a new game struct
  fmt.Println("Hello")
  // Loop that renders board while player and cpu take turns.


  board := make([][]int, 3)
  for i := range board {
    board[i] = make([]int, 3)
  }


  render.Render(board)
}
