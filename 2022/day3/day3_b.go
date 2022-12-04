package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateGroupBadgeValue(elfGroup []string) int {
	for _, char := range elfGroup[0] {
		foundInSecondElf := strings.Index(elfGroup[1], string(char))
		if foundInSecondElf != -1 {
			foundInThirdElf := strings.Index(elfGroup[2], string(char))
			if foundInThirdElf != -1 {
				return convertCharToPoints(char)
			}
		}
	}

	return 0
}

func solvePart2() {
	file, err := os.Open("day3.txt")
	checkErr(err)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	elfGroup := []string{}
	for fileScanner.Scan() {

		elfGroup = append(elfGroup, fileScanner.Text())

		if len(elfGroup) == 3 {
			total = total + calculateGroupBadgeValue(elfGroup)
			elfGroup = []string{}
		}
	}

	fmt.Printf("Part2, total points: %d \n", total)
	file.Close()

}
