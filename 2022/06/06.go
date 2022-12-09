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
	rawCrateData := ReadInFile(filename)
	unProcessedCommsData := ProcessFileToArray(rawCrateData)
	marker := FindMarker(unProcessedCommsData)

	fmt.Printf("Marker is at pos %d", marker)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(source string) int {
	returnInt, err := strconv.Atoi(strings.TrimSpace(source))
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

func FindMarker(rawData []string) int {

	for _, data := range rawData {
		for pos := range data {
			// Part 1:
			// if pos >= 4 {
			// 	lastFourChars := data[pos-4 : pos]
			// 	if uniqueChars(lastFourChars, 4) {
			// 		return pos
			// 	}
			// }

			// Part 2:
			if pos >= 14 {
				lastFourChars := data[pos-14 : pos]
				if uniqueChars(lastFourChars, 14) {
					return pos
				}
			}
		}
	}
	return -1
}

func uniqueChars(search string, numberToSearch int) bool {

	for x := 1; x < numberToSearch; x++ {
		if strings.Count(search, string(search[x])) > 1 {
			return false
		}
	}
	return true
}
