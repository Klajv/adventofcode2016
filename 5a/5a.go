package main

import (
  "fmt"
  "crypto/md5"
  "encoding/hex"
  "strconv"
)


func main() {
  input := []byte("reyedfim")
  salt := []byte{}
  password := []byte{}

  for i := 0; len(password) < 8; i++ {
    salt = []byte(strconv.Itoa(i))
    hash := md5.Sum(append(input, salt...))
    hex := hex.EncodeToString(hash[:])

    if hex[0:5] == "00000" {
      password = append(password, hex[5])
      fmt.Println(string(password))
    }
  }
}
