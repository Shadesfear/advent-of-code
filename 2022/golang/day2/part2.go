package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

// X loss
// Y draw
// Z win

// A rock
// B paper 
// C Scissors

func main(){
  content, _ := ioutil.ReadFile("../../inputs/day02.txt")
  str := string(content)
  str = strings.ReplaceAll(str, "\r", "")

  lines := strings.Split(str, "\n")

  scores := map[string]int{
    "B X": 1, // They play Paper and i have to loose
    "A X": 3, 
    "C X": 2,

    "A Y": 4,
    "B Y": 5,
    "C Y": 6,

    "A Z": 8,
    "B Z": 9,
    "C Z": 7}

  totalScore := 0
  for _, line := range lines {
    totalScore += scores[line]
  }
  fmt.Println(totalScore)
}
