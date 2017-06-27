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
    if hasssl(supernet, hypernet) {
      count++
    }
  }
  fmt.Println(count)
}

func hasssl(s []string, h []string) bool {
  for _, ss := range s {
    for i := 0; i < len(ss)-2; i++ {
      if ss[i] != ss[i+1] && ss[i] == ss[i+2] {
        for _, hh := range h {
          for j := 0; j < len(hh)-2; j++ {
            if hh[j] == ss[i+1] && ss[i] == hh[j+1] && hh[j] == hh[j+2] {
              return true
            }
          }
        }
      }
    }
  }
  return false
}
