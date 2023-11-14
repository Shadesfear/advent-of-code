package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {

  content, _ := ioutil.ReadFile("input.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")

  biggest := 0
  curr := 0

  lines := strings.Split(str, "\n")

  for _, line := range lines {
    snack, err := strconv.Atoi(line)
    curr += snack

    if err != nil {
      fmt.Println(err)
      if curr > biggest {
        biggest = curr
      }
      curr = 0
    }

  }   
  fmt.Println(biggest)
}
