package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcLargest(input []int) int {
	for j := 1; j < len(input); j++ {
		if input[0] < input[j] {
			input[0] = input[j]
		}
	}
	return input[0]
}

func main() {
	file, err := os.Open("day1.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	index := 0
	elfsAndCalories := make([]int, 0)
	elfsAndCalories = append(elfsAndCalories, 0)

	for fileScanner.Scan() {

		if fileScanner.Text() == "" {

			index++
			elfsAndCalories = append(elfsAndCalories, 0)

		} else {
			lineAsInt, err := strconv.Atoi(fileScanner.Text())
			check(err)

			elfsAndCalories[index] = elfsAndCalories[index] + lineAsInt
		}
	}

	sort.Ints(elfsAndCalories)

	elf1 := elfsAndCalories[len(elfsAndCalories)-1]
	elf2 := elfsAndCalories[len(elfsAndCalories)-2]
	elf3 := elfsAndCalories[len(elfsAndCalories)-3]

	fmt.Printf("Largest value: %d \n", elf1)
	fmt.Printf("Sum of top 3: %d : %d %d %d \n", elf1+elf2+elf3, elf1, elf2, elf3)
	file.Close()
	return
}
