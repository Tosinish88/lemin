package main

import (
	"fmt"
	"os"
	"time"

	lh "lemin/handler"
)

func main() {
	Start := time.Now()
	// ensuring that file argument is given
	if len(os.Args) < 2 {
		fmt.Println("Error: Invalid Arguments")
		fmt.Println("Usage: go run . <filename> or <filename -v> if you want to see the number of steps to complete")
		return
	}
	// reads the data from file provided
	str, err := os.ReadFile(os.Args[1])
	fmt.Println(string(str))
	if err != nil {
		fmt.Println("Error: Invalid file")
		fmt.Println(err)
		return
	}
	// checks if data is valid
	if !lh.ValidData(string(str)) {
		return
	} else {
		// finds the path and solve the problem
		lh.LeminSolver(string(str))
	}
	fmt.Println("Time taken to move ants:", time.Since(Start))
}
