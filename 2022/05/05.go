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
	unProcessedCrateData := ProcessFileToArray(rawCrateData)
	crates := CreateCrates(unProcessedCrateData)

	for cratesLoop, _ := range crates {
		fmt.Printf("%s,", crates[cratesLoop][0])
	}
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

func CreateCrates(rawData []string) [][]string {
	var crates [][]string

	numberCrates := FindNumberOfCrates(rawData)
	crates = LoadCrates(rawData, numberCrates)
	crates = ProcessRules(rawData, crates)

	return crates
}

func ProcessRules(rawData []string, crates [][]string) [][]string {
	OrderedCrates := GetAndProcessRules(rawData, crates)

	return OrderedCrates
}

func FindNumberOfCrates(rawData []string) int {
	var numberCrates int

	for _, line := range rawData {
		if !(strings.Contains(line, "[")) {
			numberItems := strings.Split(line, "   ")
			numberCrates = convertToInt(numberItems[len(numberItems)-1])
			break
		}
	}
	return numberCrates
}

func LoadCrates(rawData []string, numberCrates int) [][]string {
	crates := make([][]string, numberCrates)

	for _, line := range rawData {
		if strings.Contains(line, "[") {
			for pot := 0; pot < numberCrates; pot++ {
				crate := line[1+(pot*4)]

				if !(crate == 32) {
					crates[pot] = append(crates[pot], string(crate))
				}
			}
		}
		if len(line) == 0 {
			break
		}
	}

	return crates
}

func GetAndProcessRules(rawData []string, crates [][]string) [][]string {
	for _, line := range rawData {
		if strings.Contains(line, "move") {

			numberToMove := GetNumberToMove(line)
			fromStack := GetFromStack(line)
			toStack := GetToStack(line)

			crates = MoveCrates9001(crates, fromStack, toStack, numberToMove)
		}
	}
	return crates
}

func GetNumberToMove(data string) int {
	return convertToInt(FindBetweenTwoStrings(data, "move ", "from "))
}

func GetFromStack(data string) int {
	return convertToInt(FindBetweenTwoStrings(data, "from ", "to "))
}

func GetToStack(data string) int {
	return convertToInt(FindBetweenStringAndEnd(data, "to "))
}
func FindBetweenTwoStrings(source string, first string, second string) string {
	pos := strings.Index(source, first) + len(first)
	pos1 := strings.Index(source, second)
	return source[pos:pos1]
}

func FindBetweenStringAndEnd(source string, find string) string {
	pos := strings.Index(source, find) + len(find)
	pos1 := len(source)
	return source[pos:pos1]
}

func MoveCrates(crates [][]string, fromStack int, toStack int, numberToMove int) [][]string {

	zeroBasedFromStack := fromStack - 1
	zeroBasedToStack := toStack - 1
	var tmpCrates []string

	for moveLoop := 0; moveLoop < numberToMove; moveLoop++ {
		itemToMove := crates[zeroBasedFromStack][0]

		//Remove from
		tmpCrates = nil
		for itemLoop, _ := range crates[zeroBasedFromStack] {
			if itemLoop > 0 {
				tmpCrates = append(tmpCrates, crates[zeroBasedFromStack][itemLoop])
			}
		}
		crates[zeroBasedFromStack] = tmpCrates
		tmpCrates = nil

		//Add to Array (at beginnig)
		tmpCrates = append(tmpCrates, itemToMove)
		for itemLoop, _ := range crates[zeroBasedToStack] {
			tmpCrates = append(tmpCrates, crates[zeroBasedToStack][itemLoop])
		}

		crates[zeroBasedToStack] = tmpCrates

		// fmt.Println("Crate status:")
		// for cratesLoop, thisCrate := range crates {
		// 	for itemLoop, _ := range thisCrate {
		// 		fmt.Printf("Crate %d.  Pos: %d. Item: %s.  Length crate:%d\n", cratesLoop, itemLoop, crates[cratesLoop][itemLoop], len(thisCrate))
		// 	}
		// }
	}

	return crates
}

// Part 2
func MoveCrates9001(crates [][]string, fromStack int, toStack int, numberToMove int) [][]string {

	zeroBasedFromStack := fromStack - 1
	zeroBasedToStack := toStack - 1
	var tmpCrates []string

	// fmt.Println("Current Status")
	// showArrayMulti(crates)
	// fmt.Printf("Need to move %d from %d to %d\n", numberToMove, fromStack, toStack)

	itemsToMove := GetItemsToMove(crates[zeroBasedFromStack], numberToMove)

	// fmt.Printf("Items to Move\n")
	// showArray(itemsToMove)

	//Remove from
	tmpCrates = nil
	for itemLoop, _ := range crates[zeroBasedFromStack] {
		if itemLoop >= numberToMove {
			tmpCrates = append(tmpCrates, crates[zeroBasedFromStack][itemLoop])
		}
	}
	crates[zeroBasedFromStack] = tmpCrates
	tmpCrates = nil

	//Add to Array (at beginning)
	// fmt.Println("Before Adding")
	// showArray(crates[zeroBasedToStack])
	tmpCrates = itemsToMove
	for itemLoop, _ := range crates[zeroBasedToStack] {
		tmpCrates = append(tmpCrates, crates[zeroBasedToStack][itemLoop])
	}
	crates[zeroBasedToStack] = tmpCrates

	// fmt.Println("After Adding")
	// showArray(crates[zeroBasedToStack])

	return crates
}

func GetItemsToMove(stack []string, numberToMove int) []string {
	var ItemsToMove []string

	//	fmt.Printf("Need to move %d.  Items in array are %d", numberToMove, len(stack))
	for moveLoop := 0; moveLoop < numberToMove; moveLoop++ {
		ItemsToMove = append(ItemsToMove, stack[moveLoop])
	}

	return ItemsToMove
}

func showArray(myarray []string) {
	for cratesLoop, thisCrate := range myarray {
		fmt.Printf("Pos: %d. Item: %s.  Length crate:%d\n", cratesLoop, thisCrate, len(thisCrate))
	}
}

func showArrayMulti(multi [][]string) {
	for cratesLoop, thisCrate := range multi {
		for itemLoop, _ := range thisCrate {
			fmt.Printf("Crate %d.  Pos: %d. Item: %s.  Length crate:%d\n", cratesLoop, itemLoop, multi[cratesLoop][itemLoop], len(thisCrate))
		}
	}
}
