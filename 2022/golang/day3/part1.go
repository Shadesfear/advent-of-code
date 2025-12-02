package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)


func main()  {

  priority := make(map[rune]int)
  // Populate lowercase
  for i, ch := range "abcdefghijklmnopqrstuvwxyz" {
    priority[ch] = i + 1
  }

  // Populate uppercase
  for i, ch := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
    priority[ch] = i + 27
  }
  content, _ := ioutil.ReadFile("../../inputs/day03.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")

  backpacks := strings.Split(str, "\n")

  total := 0
  for _, backpack := range backpacks {
    seen := map[rune]bool{}
    first := backpack[:len(backpack) / 2]
    second := backpack[len(backpack) / 2:] 
    
    for _, b := range first {
      seen[b] = true
    }

    for _, b := range second {
      if seen[b] {
        total += priority[b]
        break
      }
    }
  }
  fmt.Println(total)
}
