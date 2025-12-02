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
	if _, err := os.Stat("main.go"); err == nil {
		fmt.Println("Error: main.go already exists, refusing to overwrite")
		os.Exit(1)
	}

	if _, err := os.Stat("main_test.go"); err == nil {
		fmt.Println("Error: main_test.go already exists, refusing to overwrite")
		os.Exit(1)
	}

	err := os.WriteFile("main.go", []byte(MAIN), 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile("main_test.go", []byte(TEST), 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Created main.go and main_test.go")
}
