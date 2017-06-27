package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.FieldsFunc(string(b), func (r rune) bool {
    return r == ' ' || r == '\n' || r == '\r'
  })
  var t [9]int
  var triangles [][3]int
  var blocks = len(input)/9 //545

  for i := 0; i < blocks; i++ { //0-544
    for j := 0; j < 9; j++ { //0-8
      t[j], _ = strconv.Atoi(input[i*9+j])
    }
    for k := 0; k < 3; k++ {
      if t[k] < t[k+3] + t[k+6] && t[k+3] < t[k] + t[k+6] && t[k+6] < t[k] + t[k+3] {
        triangles = append(triangles, [3]int{t[k], t[k+3], t[k+6]})
      }
    }
  }
  fmt.Println(len(triangles))
}
