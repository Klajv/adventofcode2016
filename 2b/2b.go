package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "bytes"
)

func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.Split(string(b), "\n")
  pos := [2]int{1,1} //[row, col]
  var code bytes.Buffer

  u := map[[2]int]struct{}{
    {2,0}: struct{}{},
    {1,1}: struct{}{},
    {0,2}: struct{}{},
    {1,3}: struct{}{},
    {2,4}: struct{}{},
  }

  d := map[[2]int]struct{}{
    {2,0}: struct{}{},
    {3,1}: struct{}{},
    {4,2}: struct{}{},
    {3,3}: struct{}{},
    {2,4}: struct{}{},
  }

  l := map[[2]int]struct{}{
    {0,2}: struct{}{},
    {1,1}: struct{}{},
    {2,0}: struct{}{},
    {3,1}: struct{}{},
    {4,2}: struct{}{},
  }

  r := map[[2]int]struct{}{
    {0,2}: struct{}{},
    {1,3}: struct{}{},
    {2,4}: struct{}{},
    {3,3}: struct{}{},
    {4,2}: struct{}{},
  }

  val := map[[2]int]string{
    {0,2}: "1",
    {1,1}: "2",
    {1,2}: "3",
    {1,3}: "4",
    {2,0}: "5",
    {2,1}: "6",
    {2,2}: "7",
    {2,3}: "8",
    {2,4}: "9",
    {3,1}: "A",
    {3,2}: "B",
    {3,3}: "C",
    {4,2}: "D",
  }

  for _, row := range input {
    for _, char := range row {
      switch char {
      case 'U':
        if _, ok := u[pos]; !ok {
          pos[0]--
        }
      case 'D':
        if _, ok := d[pos]; !ok {
          pos[0]++
        }
      case 'L':
        if _, ok := l[pos]; !ok {
          pos[1]--
        }
      case 'R':
        if _, ok := r[pos]; !ok {
          pos[1]++
        }
      }
    }
    code.WriteString(val[pos])
  }
  
  fmt.Println(code.String())
}
