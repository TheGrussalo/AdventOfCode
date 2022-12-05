package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	rawCleaningAssignments := ReadInFile(filename)
	CleaningAssignments := ProcessFileToArray(rawCleaningAssignments)
	// duplicates := FindDuplicates(CleaningAssignments)
	overlaps := FindOverlaps(CleaningAssignments)
	fmt.Println(overlaps)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(source string) int {
	returnInt, err := strconv.Atoi(source)
	check(err)
	return returnInt
}

func ReadInFile(filename string) string {
	file, err := os.ReadFile(filename)
	check(err)
	return string(file)
}

func ProcessFileToArray(rawData string) []string {
	var data []string

	scanner := bufio.NewScanner(strings.NewReader(rawData))

	for scanner.Scan() {
		row := scanner.Text()

		data = append(data, row)
	}
	return data
}

func FindDuplicates(CleaningAssignments []string) int {
	var duplicates int
	for _, assignment := range CleaningAssignments {
		pairAssignments := strings.Split(assignment, ",")

		firstPair := strings.Split(pairAssignments[0], "-")
		secondPair := strings.Split(pairAssignments[1], "-")

		firstPairLower := convertToInt(firstPair[0])
		firstPairHigher := convertToInt(firstPair[1])
		secondPairLower := convertToInt(secondPair[0])
		secondPairHigher := convertToInt(secondPair[1])

		// fmt.Printf("First pair lower: %d.  Higher: %d.  Second pair lower: %d.  Higher: %d\n", firstPairLower, firstPairHigher, secondPairLower, secondPairHigher)

		if secondPairLower >= firstPairLower && secondPairHigher <= firstPairHigher {
			duplicates = duplicates + 1
		} else if firstPairLower >= secondPairLower && firstPairHigher <= secondPairHigher {
			duplicates = duplicates + 1
		}
	}
	return duplicates
}

// Part 2
func FindOverlaps(CleaningAssignments []string) int {
	var overlaps int
	for _, assignment := range CleaningAssignments {
		pairAssignments := strings.Split(assignment, ",")

		firstPair := strings.Split(pairAssignments[0], "-")
		secondPair := strings.Split(pairAssignments[1], "-")

		firstPairLower := convertToInt(firstPair[0])
		firstPairHigher := convertToInt(firstPair[1])
		secondPairLower := convertToInt(secondPair[0])
		secondPairHigher := convertToInt(secondPair[1])

		// fmt.Printf("First pair lower: %d.  Higher: %d.  Second pair lower: %d.  Higher: %d\n", firstPairLower, firstPairHigher, secondPairLower, secondPairHigher)

		if firstPairLower >= secondPairLower && firstPairLower <= secondPairHigher {
			overlaps = overlaps + 1
		} else if firstPairHigher >= secondPairLower && firstPairHigher <= secondPairHigher {
			overlaps = overlaps + 1
		} else if secondPairLower >= firstPairLower && secondPairLower <= firstPairHigher {
			overlaps = overlaps + 1
		} else if secondPairHigher >= firstPairLower && secondPairHigher <= firstPairHigher {
			overlaps = overlaps + 1
		}
	}
	return overlaps
}
