package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "sort"
)

type room struct {
  name string
  sector int
  checksum string
}

type pairlist [][2]int

func (p pairlist) Less(i, j int) bool {
  if p[i][1] == p[j][1] {
    return p[i][0] > p[j][0]
  } else {
    return p[i][1] < p[j][1]
  }
}

func (p pairlist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairlist) Len() int {
	return len(p)
}

func main() {
  b, _ := ioutil.ReadFile("input.txt")
  input := strings.FieldsFunc(string(b), func (r rune) bool {
    return r == '[' || r == ']' || r == '\n' || r == '\r'
  })
  var rooms []room

  for i := 0; i < len(input); i=i+2 {
    split := strings.Split(input[i], "-")
    name := strings.Join(split[:len(split)-1], "")
    sector, _ := strconv.Atoi(split[len(split)-1])
    checksum := input[i+1]
    countmap := map[rune]int{}
    var countarr pairlist

    for _, r := range name {
      _, ok := countmap[r]
      if !ok {
        countmap[r] = 0
      }
      countmap[r] = countmap[r] + 1
    }

    for key, val := range countmap {
      countarr = append(countarr, [2]int{int(key), val})
    }

    sort.Sort(sort.Reverse(countarr))

    if byte(countarr[0][0]) == checksum[0] && byte(countarr[1][0]) == checksum[1] && byte(countarr[2][0]) == checksum[2] && byte(countarr[3][0]) == checksum[3] && byte(countarr[4][0]) == checksum[4] {
      rooms = append(rooms, room{strings.Join(split[:len(split)-1], " "), sector, checksum})
    }
  }

  for i, r := range rooms {
    var roomname []byte
    for _, b := range r.name {
      if b == ' ' {
        roomname = append(roomname, ' ')
      } else {
        roomname = append(roomname, byte(((b - 'a' + rune(r.sector)) % 26) + 'a'))
      }
    }
    rooms[i].name = string(roomname)
  }

  fmt.Println(rooms)
}
