package main

import (
	"strconv"
	"strings"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func readContainerStacks(line string) [][]string {
	step1 := strings.Replace(line, "    ", " ", -1)

	containers := strings.Split(step1, " ")

	stackSliceOfSlices := make([][]string, len(containers))
	for index, container := range containers {
		rawContainers := strings.Split(container, " ")
		for index, rawContainer := range rawContainers {
			rawContainer = strings.Replace(rawContainer, "[", "", -1)
			rawContainer = strings.Replace(rawContainer, "]", "", -1)
			rawContainer = strings.Replace(rawContainer, " ", "", -1)
			rawContainers[index] = rawContainer
		}

		stackSliceOfSlices[index] = append(stackSliceOfSlices[index], rawContainers...)
	}
	return stackSliceOfSlices
}

func solvePart1(input string) []string {

	lines := strings.Split(input, "\n")
	stacks := [][]string{}
	firstTime := true
	firstMove := true
	for _, line := range lines {
		if strings.Contains(line, "[") {
			newContainers := readContainerStacks(line)
			if firstTime {
				stacks = newContainers
				firstTime = false

			} else {
				for index, container := range newContainers {

					stacks[index] = append(stacks[index], container...)
				}

			}

		} else if strings.Contains(line, "move") {

			// dont do this calc every move line :D
			if firstMove {

				for stack := 0; stack < len(stacks); stack++ {
					stacks[stack] = removeEmptyStrings(stacks[stack])
					for i, j := 0, len(stacks[stack])-1; i < j; i, j = i+1, j-1 {
						stacks[stack][i], stacks[stack][j] = stacks[stack][j], stacks[stack][i]
					}
				}

				firstMove = false
			}
			instructions := strings.Split(line, " ")
			from, err := strconv.Atoi(instructions[3])
			checkErr(err)
			amount, err := strconv.Atoi(instructions[1])
			checkErr(err)
			to, err := strconv.Atoi(instructions[5])
			checkErr(err)

			for i := 0; i < amount; i++ {
				stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
				stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
			}
		}
	}
	result := []string{}
	for i := 0; i < len(stacks); i++ {
		result = append(result, stacks[i][len(stacks[i])-1])
	}

	return result
}
