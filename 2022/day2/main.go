package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Returns the round score using the strategy guide from part 1
func strategyGuideOne(opponentMove string, myMove string) int {
	var shapeScore int
	var outcomeScore int
	switch opponentMove {
	case "A":
		// opponent chooses rock
		switch myMove {
		case "X":
			// tie
			shapeScore = 1
			outcomeScore = 3
		case "Y":
			// win
			shapeScore = 2
			outcomeScore = 6
		case "Z":
			// lose
			shapeScore = 3
			outcomeScore = 0
		}
	case "B":
		// opponent chooses paper
		switch myMove {
		case "X":
			// lose
			shapeScore = 1
			outcomeScore = 0
		case "Y":
			// tie
			shapeScore = 2
			outcomeScore = 3
		case "Z":
			// win
			shapeScore = 3
			outcomeScore = 6
		}
	case "C":
		// opponent chooses scissors
		switch myMove {
		case "X":
			// win
			shapeScore = 1
			outcomeScore = 6
		case "Y":
			// lose
			shapeScore = 2
			outcomeScore = 0
		case "Z":
			// tie
			shapeScore = 3
			outcomeScore = 3
		}
	}
	return shapeScore + outcomeScore
}

// Returns the round score using the strategy guide from part 2
func strategyGuideTwo(opponentMove string, outcome string) int {
	var shapeScore int
	var outcomeScore int
	switch opponentMove {
	case "A":
		// opponent chooses rock
		switch outcome {
		case "X":
			// lose
			shapeScore = 3
			outcomeScore = 0
		case "Y":
			// tie
			shapeScore = 1
			outcomeScore = 3
		case "Z":
			// win
			shapeScore = 2
			outcomeScore = 6
		}
	case "B":
		// opponent chooses paper
		switch outcome {
		case "X":
			// lose
			shapeScore = 1
			outcomeScore = 0
		case "Y":
			// tie
			shapeScore = 2
			outcomeScore = 3
		case "Z":
			// win
			shapeScore = 3
			outcomeScore = 6
		}
	case "C":
		// opponent chooses scissors
		switch outcome {
		case "X":
			// lose
			shapeScore = 2
			outcomeScore = 0
		case "Y":
			// tie
			shapeScore = 3
			outcomeScore = 3
		case "Z":
			// win
			shapeScore = 1
			outcomeScore = 6
		}
	}
	return shapeScore + outcomeScore
}

func main() {
	// Read input file
	f, err := os.Open("./input")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var totalScore1 int
	var totalScore2 int
	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Fields(line)
		totalScore1 += strategyGuideOne(round[0], round[1])
		totalScore2 += strategyGuideTwo(round[0], round[1])
	}

	// Part 1
	fmt.Println("Part 1:", totalScore1)

	// Part 2
	fmt.Println("Part 2:", totalScore2)
}
