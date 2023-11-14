package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func modifyBit(n,p,b int) int {
  mask := 1 << p
  return (n & ^mask) | ((b << p) & mask)
}

func findPosish(n int) int {
  i := 1
  posish := 1
  for {
    if i & n > 0 {
      break
    }
    i = i << 1
    posish++
  }
  return posish
}


func main()  {

  priority := make(map[rune]int)
  for i, ch := range "abcdefghijklmnopqrstuvwxyz" {
    priority[ch] = i + 1
  }
  for i, ch := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
    priority[ch] = i + 27
  }

  content, _ := ioutil.ReadFile("input.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")
  backpacks := strings.Split(str, "\n")

  total := 0
  for i := 0; i <= len(backpacks)-3; i += 3 {
    bitMasks := []int{0,0,0}

    for j := 0; j < 3; j++ {
      backpack := backpacks[i+j]
      for _, b := range backpack {
        bitMasks[j] = modifyBit(bitMasks[j], priority[b] - 1, 1)
      }
    }
    hmm := bitMasks[0] & bitMasks[1] & bitMasks[2]
    total += findPosish(hmm)
  }

  fmt.Println(total)
}
