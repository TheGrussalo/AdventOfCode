package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Type string
	Name string
	Size int
}

type Tree map[string]TreeNode

func main() {
	filename := "input.txt"
	rawData := ReadInFile(filename)
	unProcessFileStructure := ProcessFileToArray(rawData)
	tree := ProcessFileStructure(unProcessFileStructure)

	totalSize := tree["/"].Size
	fmt.Printf("Total Size is %d\n", totalSize)
	maxSize := 70000000
	needed := 30000000
	unused := maxSize - totalSize
	toDelete := needed - unused
	fmt.Printf("Size to delete is %d\n", toDelete)

	//Part1:
	DisplayDirectories(tree, 100000)
	//Part2:
	FindDirectory(tree, toDelete)
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

func ProcessFileStructure(rawData []string) Tree {
	currentFolder := "/"
	tree := Tree{}
	for rawFilePos, data := range rawData {

		if data == "$ ls" { //Add following items to the tree.
			tree = ProcessList(rawData, rawFilePos, currentFolder, tree)
		}
		if strings.Contains(data, "$ cd") {
			newFolder := strings.Split(data, " ")
			if newFolder[2] == "/" {
				currentFolder = "/"
			} else if newFolder[2] == ".." {
				currentFolder = GetLevelUp(currentFolder)
			} else {
				currentFolder = currentFolder + newFolder[2] + "/"
			}
		}
	}

	return tree
}

func ProcessList(rawData []string, startingRow int, currentFolder string, tree Tree) Tree {
	for rawFilePos, data := range rawData {
		if rawFilePos > startingRow {
			if strings.Contains(data, "$ ") {
				break
			} else {
				rowSplit := strings.Split(data, " ")
				// Each row needs to be added to structure
				if strings.HasPrefix(data, "dir") {
					//We have a directory to add
					fileName := rowSplit[1]

					item := TreeNode{Type: "Directory", Name: fileName, Size: -1}
					tree[currentFolder+fileName+"/"] = item

				} else { //We have a file
					fileSize := convertToInt(rowSplit[0])
					fileName := rowSplit[1]

					item := TreeNode{Type: "File", Name: fileName, Size: fileSize}

					tree[currentFolder+fileName] = item
				}
			}
		}
	}

	tree = CalculateDirectorySizes(tree)
	return tree
}

func GetLevelUp(currentFolder string) string {
	newFolder := ""

	pos := strings.LastIndex(currentFolder[0:len(currentFolder)-1], "/")
	newFolder = currentFolder[0:pos] + "/"

	return newFolder
}

func CalculateDirectorySizes(tree Tree) Tree {
	size := -2
	for treeID, treeItem := range tree {
		if treeItem.Type == "Directory" {
			size = CalculateDirectorySize(treeID, tree)
			treeItem.Size = size
			tree[treeID] = treeItem
		}
	}

	rootItem := tree["/"]
	rootItem.Size = CalculateDirectorySize("/", tree)
	tree["/"] = rootItem

	return tree
}

func CalculateDirectorySize(path string, tree Tree) int {
	size := 0
	for treeID, treeItem := range tree {
		if treeItem.Type == "File" && strings.HasPrefix(treeID, path) {
			size = size + treeItem.Size
		}
	}
	return size
}

func DisplayDirectories(tree Tree, maxSize int) {
	totalSize := 0
	for treeID, treeItem := range tree {
		if treeItem.Type == "Directory" && treeItem.Size <= maxSize {
			fmt.Printf("%s - %d\n", treeID, treeItem.Size)
			totalSize = totalSize + treeItem.Size
		}
	}
	fmt.Printf("Total Size of items is %d\n", totalSize)
}

func FindDirectory(tree Tree, maxSize int) {
	found := TreeNode{"", "", 0}
	for _, treeItem := range tree {
		if treeItem.Type == "Directory" && treeItem.Size >= maxSize && (treeItem.Size < found.Size || found.Size == 0) {
			found = treeItem
		}
	}
	fmt.Printf("Directory to delete is %s which is %d bytes\n", found.Name, found.Size)
}

func DisplayTree(tree Tree) {

	for treeID, item := range tree {
		if item.Type == "Directory" {
			fmt.Printf("(%s) - %d -- %v\n", item.Name, item.Size, treeID)

		} else {
			fmt.Printf("%s - %d -- %v\n", item.Name, item.Size, treeID)
		}
	}
}
