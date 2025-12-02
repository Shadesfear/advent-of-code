package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main(){
  content, _ := ioutil.ReadFile("../../inputs/day02.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")

  lines := strings.Split(str, "\n")

  scores := map[string]int{
    "B X": 1,
    "A X": 4,
    "C X": 7,
    "A Y": 8,
    "B Y": 5,
    "C Y": 2,
    "A Z": 3,
    "B Z": 9,
    "C Z": 6}

  totalScore := 0
  for _, line := range lines {
    totalScore += scores[line]
  }
  fmt.Println(totalScore)
}
