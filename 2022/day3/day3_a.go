package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func convertCharToPoints(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}
func splitHalf(line string) (string, string) {
	return line[0 : len(line)/2], line[len(line)/2 : len(line)]
}

func calculateTotal(line string) int {
	firstPart, secondPart := splitHalf(line)
	for _, char := range firstPart {
		foundOrNot := strings.Index(secondPart, string(char))
		if foundOrNot != -1 {
			return convertCharToPoints(char)
		}
	}
	return 0
}

func solvePart1() {
	file, err := os.Open("day3.txt")
	checkErr(err)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {

		total = total + calculateTotal(fileScanner.Text())
	}

	fmt.Printf("Part1, Total points: %d \n", total)
	file.Close()

}
