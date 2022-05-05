package render

import (
  "fmt"
)

// Takes a matrix and prints ASCII to console.
func Render (board [][]int) {

  fmt.Print(board[0][0], " | ", board[0][1], " | ", board[0][2],
            "\n_   _   _\n\n",
            board[1][0], " | ", board[1][1], " | ", board[1][2],
            "\n_   _   _\n\n",
            board[2][0], " | ", board[2][1], " | ", board[2][2], "\n")

}
