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
	calibrationData := ReadInFile(filename)

	elfCal := ProcessCalibration(calibrationData)
	total := calculateCoordinates(elfCal)
	fmt.Println(total)

	part2 := convertLetterToNumbers(calibrationData)
	elfCal2 := ProcessCalibration(part2)
	part2Total := calculateCoordinates(elfCal2)
	fmt.Println(part2Total)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInFile(filename string) string {
	file, err := os.ReadFile(filename)
	check(err)
	return string(file)
}

func ProcessCalibration(rawData string) []string {
	var coordinates []string

	scanner := bufio.NewScanner(strings.NewReader(rawData))

	for scanner.Scan() {
		row := scanner.Text()
		onlyNumbers := ""
		for _, a := range row {
			if int(a) >= 48 && int(a) <= 57 {
				onlyNumbers += string(a)
			}
		}

		firstDigit := onlyNumbers[0:1]
		lastDigit := onlyNumbers[0:1]
		if len(onlyNumbers) > 1 {
			lastDigit = onlyNumbers[len(onlyNumbers)-1:]
		}
		newRow := firstDigit + lastDigit
		coordinates = append(coordinates, newRow)
	}
	return coordinates
}

func calculateCoordinates(data []string) int {
	total := 0
	for _, row := range data {
		if len(row) != 2 {
			fmt.Println("Panic, row should have two digits")
		}
		value, err := strconv.Atoi(string(row))
		check(err)
		total = total + value
	}
	return total
}

func convertLetterToNumbers(data string) string {
	var newData []string

	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		row := scanner.Text()
		// fmt.Println(row)
		row = strings.ReplaceAll(row, "one", "one1one")
		row = strings.ReplaceAll(row, "two", "two2two")
		row = strings.ReplaceAll(row, "three", "three3three")
		row = strings.ReplaceAll(row, "four", "four4four")
		row = strings.ReplaceAll(row, "five", "five5five")
		row = strings.ReplaceAll(row, "six", "six6six")
		row = strings.ReplaceAll(row, "seven", "seven7seven")
		row = strings.ReplaceAll(row, "eight", "eight8eight")
		row = strings.ReplaceAll(row, "nine", "nine9nine")
		newData = append(newData, row)
	}

	return strings.Join(newData, "\n")
}
