package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)


func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.Split(string(b), "\r\n")
  password := make([]rune, 8)
  count := [8]map[rune]int{}
  for i := 0; i < 8; i++ {
    count[i] = map[rune]int{}
  }

  for _, s := range input {
    for i, r := range s {
      if _, ok := count[i][r]; !ok {
        count[i][r] = 0
      }
      count[i][r] = count[i][r] + 1
    }
  }

  for i := 0; i < 8; i++ {
    maxr := ' '
    maxc := 0
    for r, c := range count[i] {
      if c > maxc {
        maxr = r
        maxc = c
      }
    }
    password[i] = maxr
  }
  fmt.Println(string(password))
}
