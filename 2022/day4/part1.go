package main

import (
	"strconv"
	"strings"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func convertSectionToNumbers(section string) []int {
	points := strings.Split(section, "-")
	start := points[0]
	end := points[1]
	min, err := strconv.Atoi(start)
	checkErr(err)
	max, err := strconv.Atoi(end)
	checkErr(err)
	return makeRange(min, max)
}
func checkIfSubset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func findFullyContainedSections(line string) int {
	elfSections := strings.Split(line, ",")

	elf1 := convertSectionToNumbers(elfSections[0])
	elf2 := convertSectionToNumbers(elfSections[1])

	if checkIfSubset(elf1, elf2) {
		return 1
	} else {
		if checkIfSubset(elf2, elf1) {
			return 1
		} else {
			return 0
		}
	}
}

func solvePart1(input string) int {
	fullyContainedAmount := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		fullyContainedAmount = fullyContainedAmount + findFullyContainedSections(line)

	}

	return fullyContainedAmount
}
