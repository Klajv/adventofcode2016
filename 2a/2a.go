package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.Split(string(b), "\n")
  pos := [2]int{1,1} //[row, col]
  code := 0

  for _, row := range input {
    for _, char := range row {
      switch char {
      case 'U':
        if pos[0] > 0 {
          pos[0]--
        }
      case 'D':
        if pos[0] < 2 {
          pos[0]++
        }
      case 'L':
        if pos[1] > 0 {
          pos[1]--
        }
      case 'R':
        if pos[1] < 2 {
          pos[1]++
        }
      }
    }
    code = code*10 + pos[0]*3 + pos[1] + 1
  }

  fmt.Println(code)
}
