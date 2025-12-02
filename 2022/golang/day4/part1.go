package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func cleanMask(mask1, mask2 int) (int, int) {
	commonMask := mask1 | mask2
	return mask1 & commonMask, mask2 & commonMask
}

func toBitMask(rng string) int {
  
  split := strings.Split(rng, "-")

  left, _ := strconv.Atoi(split[0])
  right ,_ := strconv.Atoi(split[1])
  
  mask := 0
  for i := left; i <= right; i++ {
    mask |= 1 << uint(i)
  } 

  return mask
  
}

func setBit(n, b int) int {
  return n | (1 << b)
}

// func overLapping(first, second int) bool {
//   return (first | second) == first || (first | second) == second
// }

func overLapping(a, b int) bool {
	return (a & b == a && a != b && a | b == b) || (a & b == b && a != b && a | b == a)
}

func main()  {
  content, _ := ioutil.ReadFile("../../inputs/day04.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")
  lines := strings.Split(str, "\n")

  total := 0
  for _, line := range lines {
    if len(line) < 1 {
      break
    }

    s := strings.Split(line, ",")
    first := toBitMask(s[0])
    second := toBitMask(s[1])

		cleanMask1, cleanMask2 := cleanMask(first, second)

    if overLapping(cleanMask1, cleanMask2) {
      total++
    }

  }

  fmt.Println(total)
}
