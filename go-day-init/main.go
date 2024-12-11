package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed main_go
var MAIN string

//go:embed main_test
var TEST string

func main() {
	err := os.WriteFile("main.go", []byte(MAIN), 0o644)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("main_test.go", []byte(TEST), 0o644)
	if err != nil {
		fmt.Println(err)
	}
}
