package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pointsForGame(row string) int {
	split := strings.Split(row, " ")
	elf := split[0]
	player := split[1]
	if elf == "A" && player == "X" {
		return 4
	} else if elf == "B" && player == "Y" {
		return 5
	} else if elf == "C" && player == "Z" {
		return 6
	} else if elf == "A" && player == "Y" {
		return 8
	} else if elf == "B" && player == "Z" {
		return 9
	} else if elf == "C" && player == "X" {
		return 7
	} else if elf == "A" && player == "Z" {
		return 3
	} else if elf == "B" && player == "X" {
		return 1
	} else if elf == "C" && player == "Y" {
		return 2
	}
	return 0
}

// a x rock
// b y paper
// c z scissors

func solvePart1() {
	file, err := os.Open("day2.txt")
	checkErr(err)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0

	for fileScanner.Scan() {
		totalPoints = totalPoints + pointsForGame(fileScanner.Text())
	}
	file.Close()

	fmt.Println(totalPoints)
}
