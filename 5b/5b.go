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
  password := []byte("xxxxxxxx")
  count := 0

  for i := 0; count < 8; i++ {
    salt = []byte(strconv.Itoa(i))
    hash := md5.Sum(append(input, salt...))
    hex := hex.EncodeToString(hash[:])

    if hex[0:5] == "00000" && hex[5] < '8' {
      index, _ := strconv.Atoi(string(hex[5]))
      if password[index] == 'x' {
        password[index] = hex[6]
        fmt.Println(string(password))
        count++
      }
    }
  }
}
