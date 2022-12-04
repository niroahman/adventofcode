package main

import (
	"bufio"
	"os"
	"strings"
)

func pointsForGame2(row string) int {
	split := strings.Split(row, " ")
	elf := split[0]
	player := split[1]
	if player == "X" {
		if elf == "A" {
			return 3
		} else if elf == "B" {
			return 1
		} else if elf == "C" {
			return 2
		}

	} else if player == "Y" {

		if elf == "A" {
			return 4
		} else if elf == "B" {
			return 5
		} else if elf == "C" {
			return 6
		}
	} else if player == "Z" {

		if elf == "A" {
			return 8
		} else if elf == "B" {
			return 9
		} else if elf == "C" {
			return 7
		}

	}

	return 0
}

// a x rock
// b y paper
// c z scissors

func solvePart2() {
	file, err := os.Open("day2.txt")
	checkErr(err)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0

	for fileScanner.Scan() {
		totalPoints = totalPoints + pointsForGame2(fileScanner.Text())
	}
	file.Close()

	print(totalPoints)
}
