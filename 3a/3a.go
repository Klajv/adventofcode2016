package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.Split(string(b), "\n")
  var t [3]int
  var triangles [][3]int

  for _, k := range input {
    for i, val := range strings.Fields(k) {
      t[i], _ = strconv.Atoi(val)
    }
    if t[0] < t[1] + t[2] && t[1] < t[0] + t[2] && t[2] < t[0] + t[1] {
      triangles = append(triangles, t)
    }
  }
  fmt.Println(len(triangles))
}
