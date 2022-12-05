package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	rawPackingGuide := ReadInFile(filename)
	packingGuide := ProcessFileToArray(rawPackingGuide)
	// score := ProcessPackingGuide(packingGuide)
	score := FindBadges(packingGuide)
	fmt.Println(score)
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

func ProcessFileToArray(rawData string) []string {
	var data []string

	scanner := bufio.NewScanner(strings.NewReader(rawData))

	for scanner.Scan() {
		row := scanner.Text()

		data = append(data, row)
	}
	return data
}

func ProcessPackingGuide(packing []string) int {
	totalPriority := 0

	for _, rucksack := range packing {
		firstCompartment := string(rucksack[0 : len(rucksack)/2])
		secondCompartment := string(rucksack[len(rucksack)/2 : len(rucksack)])
		duplicates := FindDuplicates(firstCompartment, secondCompartment)
		priorty := FindPriority(duplicates)
		fmt.Printf("Duplicates:%s. Priority:%d\n", duplicates, priorty)
		totalPriority = totalPriority + priorty
	}

	return totalPriority
}

func FindDuplicates(first string, second string) []string {
	var duplicates []string
	for _, firstWordLetter := range first {
		for _, secondWordLetter := range second {
			//			fmt.Printf("First: %v.  Second %v. {%s}{%s}\n", firstWordLetter, secondWordLetter, string(firstWordLetter), string(secondWordLetter))
			if firstWordLetter == secondWordLetter {
				duplicates = append(duplicates, string(firstWordLetter))
			}
		}
	}
	return duplicates
}

func FindPriority(duplicates []string) int {

	//Lowercase item types a through z have priorities 1 through 26.
	//Uppercase item types A through Z have priorities 27 through 52.
	priorty := 0
	for _, item := range duplicates {
		itemRune := rune(item[0])
		//Rune a=97.  Rune z=122
		if (itemRune >= 97) && (itemRune <= 122) {
			priorty = int(itemRune) - 96
		} else if (itemRune >= 65) && (itemRune <= 90) {
			//Rune A=65.  Rune Z=90
			priorty = int(itemRune) - (65 - 27)
		}
		//fmt.Printf("Rune %v. {%s}.  Priority: %d\n", rune(item[0]), item, priorty)
	}

	return priorty
}

// Part 2
func FindBadges(packing []string) int {
	totalPriority := 0

	for count, _ := range packing {
		if count%3 == 0 {
			badge := FindTheBadge(packing[count], packing[count+1], packing[count+2])
			priorty := FindPriority(badge)
			totalPriority = totalPriority + priorty
		}
	}

	return totalPriority
}

func FindTheBadge(firstBag string, secondBag string, thirdBag string) []string {
	var badge []string

	for _, firstBagItem := range firstBag {
		if strings.Contains(secondBag, string(firstBagItem)) {
			if strings.Contains(thirdBag, string(firstBagItem)) {
				badge = append(badge, string(firstBagItem))
			}
		}
	}
	return badge
}
