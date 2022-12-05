package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	elfCalRaw := ReadInFile(filename)
	elfCal := ProcessCalories(elfCalRaw)
	highestCalorie := HighestCalorieCount(elfCal)
	fmt.Println(highestCalorie)
	elfCal = OrderCalories(elfCal)
	fmt.Printf("\nTop:%d\n", elfCal[len(elfCal)-1])
	fmt.Printf("Second:%d\n", elfCal[len(elfCal)-2])
	fmt.Printf("Third:%d\n", elfCal[len(elfCal)-3])
	fmt.Printf("In total that's %d calories", elfCal[len(elfCal)-1]+elfCal[len(elfCal)-2]+elfCal[len(elfCal)-3])

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInFile(filename string) string {
	elfCal, err := os.ReadFile(filename)
	check(err)
	return string(elfCal)
}

func ProcessCalories(rawData string) []int {
	var elfs []int
	individualElfCalories := 0

	scanner := bufio.NewScanner(strings.NewReader(rawData))

	for scanner.Scan() {
		row := scanner.Text()

		if len(row) == 0 {
			elfs = append(elfs, individualElfCalories)
			individualElfCalories = 0
		} else {
			intOfRow, err := strconv.ParseInt(row, 0, 64)
			check(err)
			individualElfCalories = individualElfCalories + int(intOfRow)
		}
	}
	return elfs
}

func HighestCalorieCount(calories []int) int {
	highest := -1

	for _, elfCalorie := range calories {
		if elfCalorie > highest {
			highest = elfCalorie
		}
	}

	return highest
}

func OrderCalories(calories []int) []int {
	sort.Ints(calories)
	return calories
}
