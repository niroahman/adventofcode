package main

import "strings"

func checkIfContains(arr1 []int, arr2 []int) bool {
	for _, num1 := range arr1 {
		for _, num2 := range arr2 {
			if num1 == num2 {
				return true
			}
		}
	}
	return false
}

func findOverLappedSections(line string) int {
	elfSections := strings.Split(line, ",")

	elf1 := convertSectionToNumbers(elfSections[0])
	elf2 := convertSectionToNumbers(elfSections[1])

	if checkIfContains(elf1, elf2) {
		return 1
	} else {
		return 0
	}
}

func solvePart2(input string) int {
	overlappedSections := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		overlappedSections = overlappedSections + findOverLappedSections(line)

	}

	return overlappedSections
}
