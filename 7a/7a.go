package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)


func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.Split(string(b), "\r\n")
  count := 0

  for _, line := range input {
    split := strings.Split(line, "[")
    supernet := []string{}
    hypernet := []string{}

    for i, s := range split {
      if i == 0 {
        supernet = append(supernet, s)
      } else {
        ss := strings.Split(s, "]")
        supernet = append(supernet, ss[1])
        hypernet = append(hypernet, ss[0])
      }
    }
    // Look for ABBAs in hypernet and supernet
    if hasabba(supernet) {
      if !hasabba(hypernet) {
        count++
      }
    }
  }
  fmt.Println(count)
}

func hasabba(s []string) bool {
  for _, ss := range s {
    for i := 0; i < len(ss)-3; i++ {
      if ss[i] != ss[i+1] && ss[i] == ss[i+3] && ss[i+1] == ss[i+2] {
        return true
      }
    }
  }
  return false
}
