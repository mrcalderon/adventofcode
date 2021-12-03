package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type BitCriteria struct {
	Bit      rune
	Position int
}

func intArrayToString(arr []int) string {
	var strArr string
	for _, integer := range arr {
		strArr = strArr + strconv.Itoa(integer)
	}
	return strArr
}

// From a list of binary numbers, return the most common bit in a given position (0 or 1)
func mostCommonBit(binaryNumbers []string, position int) rune {
	bitOneCount := 0
	bitZeroCount := 0

	// count ocurrences of each bit in the given position
	for _, binaryNumber := range binaryNumbers {
		if binaryNumber[position] == '1' {
			bitOneCount++
		} else {
			bitZeroCount++
		}
	}

	// return the most common bit
	if bitOneCount >= bitZeroCount {
		return '1'
	} else {
		return '0'
	}
}

// From a list of binary numbers, return the least common bit in a given position (0 or 1)
func leastCommonBit(binaryNumbers []string, position int) rune {
	bitOneCount := 0
	bitZeroCount := 0

	// count ocurrences of each bit in the given position
	for _, binaryNumber := range binaryNumbers {
		if binaryNumber[position] == '1' {
			bitOneCount++
		} else {
			bitZeroCount++
		}
	}

	// return the least common bit
	if bitOneCount < bitZeroCount {
		return '1'
	} else {
		return '0'
	}
}

// From a list of binary numbers, return a list of numbers that match a given bit criteria
func bitCriteriaMatches(binaryNumbers []string, criteria BitCriteria) []string {

	// find matches
	var matches []string
	for _, binaryNumber := range binaryNumbers {
		bit := []rune(binaryNumber)[criteria.Position] // convert string to array of runes
		if bit == criteria.Bit {
			matches = append(matches, binaryNumber)
		}
	}

	// return list of matches
	return matches
}

func main() {
	fmt.Println("=== Day 3 (part 2) ===")

	// Open input file
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read input file
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	binaryLength := len(input[0])

	// get oxygen generator rating value
	matches := input // start with all binary numbers from input
	for i := 0; i < binaryLength; i++ {
		commonBit := mostCommonBit(matches, i)
		criteria := BitCriteria{Bit: commonBit, Position: i}
		matches = bitCriteriaMatches(matches, criteria)
		if len(matches) == 1 {
			break
		}
	}
	oxygenRatingBinary := matches[0]
	oxygenRatingDecimal, err := strconv.ParseInt(oxygenRatingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// get CO2 scrubber rating value
	matches = input // start with all binary numbers from input
	for i := 0; i < binaryLength; i++ {
		commonBit := leastCommonBit(matches, i)
		criteria := BitCriteria{Bit: commonBit, Position: i}
		matches = bitCriteriaMatches(matches, criteria)
		if len(matches) == 1 {
			break
		}
	}
	co2RatingBinary := matches[0]
	co2RatingDecimal, err := strconv.ParseInt(co2RatingBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Print results
	fmt.Println(oxygenRatingDecimal * co2RatingDecimal)
}
