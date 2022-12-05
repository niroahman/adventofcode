package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input_bytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal("Failed to read input file.")
	}

	input := string(input_bytes)
	input = strings.Trim(input, "\n")

	fmt.Println("Part 1 solution: ", solvePart1(input))
	fmt.Println("Part 2 solution: ", solvePart2(input))
}
