package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	strategyGuide := ReadInFile(filename)
	rockPaperScissors := ProcessStrategyGuide(strategyGuide)
	score := PlayRockPaperScissors(rockPaperScissors)
	fmt.Println(score)
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

func ProcessStrategyGuide(rawStratgyGuide string) [][]string {
	var RockPaperScissors [][]string

	scanner := bufio.NewScanner(strings.NewReader(rawStratgyGuide))

	for scanner.Scan() {
		row := scanner.Text()
		data := strings.Split(row, " ")
		RockPaperScissors = append(RockPaperScissors, data)
	}
	return RockPaperScissors
}

func PlayRockPaperScissors(strategyGuide [][]string) int {
	score := 0

	for round, _ := range strategyGuide {
		//		fmt.Printf("Player 1: '%s' - Suggested '%s'\n", strategyGuide[round][0], strategyGuide[round][1])
		score = score + PlayARoundOfRockPaperScissorsPart2(strategyGuide[round][0], strategyGuide[round][1])
	}

	return score
}

func PlayARoundOfRockPaperScissors(player1 string, suggested string) int {
	score := 0
	bonus := 0

	// A = Rock
	// B = Paper
	// C = Scissors

	// X = Rock
	// Y = Paper
	// Z = Scissors

	if player1 == "A" { //Rock
		if suggested == "Y" { //Paper
			score = 6 // Win
		} else if suggested == "X" { //Rock
			score = 3 // Draw
		} else if suggested == "Z" { //Scissors
			score = 0 // Loss
		}
	}

	if player1 == "B" { //Paper
		if suggested == "Y" { //Paper
			score = 3 // Draw
		} else if suggested == "X" { //Rock
			score = 0 // Loss
		} else if suggested == "Z" { //Scissors
			score = 6 // Win
		}
	}

	if player1 == "C" { //Scissors
		if suggested == "Y" { //Paper
			score = 0 // Loss
		} else if suggested == "X" { //Rock
			score = 6 // Win
		} else if suggested == "Z" { //Scissors
			score = 3 // Draw
		}
	}

	//Bonus Score - Item played:
	switch suggested {
	case "Y":
		bonus = 2
	case "X":
		bonus = 1
	case "Z":
		bonus = 3
	}

	//	fmt.Printf("Score:%d - Bonus:%d.  Played:%s, Suggested:%s\n", score, bonus, player1, suggested)
	return score + bonus
}

func PlayARoundOfRockPaperScissorsPart2(player1 string, suggested string) int {
	score := 0
	bonus := 0
	cardPlayed := ""

	// A = Rock
	// B = Paper
	// C = Scissors

	//X means you need to lose,
	//Y means you need to end the round in a draw, and
	//Z means you need to win

	if player1 == "A" { //Rock
		if suggested == "Y" {
			score = 3 //Draw
			cardPlayed = "Rock"
		} else if suggested == "X" {
			score = 0 // Loss
			cardPlayed = "Scissors"
		} else if suggested == "Z" {
			score = 6 // Win
			cardPlayed = "Paper"
		}
	}

	if player1 == "B" { //Paper
		if suggested == "Y" {
			score = 3 //Draw
			cardPlayed = "Paper"
		} else if suggested == "X" {
			score = 0 // Loss
			cardPlayed = "Rock"
		} else if suggested == "Z" {
			score = 6 // Win
			cardPlayed = "Scissors"
		}
	}

	if player1 == "C" { //Scissors
		if suggested == "Y" {
			score = 3 //Draw
			cardPlayed = "Scissors"
		} else if suggested == "X" {
			score = 0 // Loss
			cardPlayed = "Paper"
		} else if suggested == "Z" {
			score = 6 // Win
			cardPlayed = "Rock"
		}
	}

	// 1 for Rock, 2 for Paper, and 3 for Scissors

	//Bonus Score - Item played:
	switch cardPlayed {
	case "Rock":
		bonus = 1
	case "Paper":
		bonus = 2
	case "Scissors":
		bonus = 3
	}

	fmt.Printf("Score:%d - Bonus:%d.  Played:%s, Suggested:%s\n", score, bonus, player1, suggested)
	return score + bonus
}
