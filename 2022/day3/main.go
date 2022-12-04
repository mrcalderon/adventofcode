package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create priorities map
	itemTypes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorities := make(map[rune]int)
	for i, c := range itemTypes {
		priorities[c] = i + 1
	}

	// Read input file
	f, err := os.Open("./input")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Calculate priorities
	var sumOfCommonItemPriorities int
	var sumOfBadgePriorites int
	var lineNumber int
	badgeItems := make(map[rune]int)
	commonItems := make(map[rune]int)
	for scanner.Scan() {
		rucksack := scanner.Text()

		// Part 1
		rucksackLength := len(rucksack)
		compartment1 := rucksack[0 : rucksackLength/2]
		compartment2 := rucksack[rucksackLength/2 : rucksackLength]
		var commonItemType rune
		var badge rune
		for _, item := range compartment1 {
			commonItems[item] = 1
		}

		for _, item := range compartment2 {
			if commonItems[item] == 1 {
				commonItemType = item
				break
			}
		}
		commonItemTypePriority := priorities[commonItemType]
		sumOfCommonItemPriorities += commonItemTypePriority
		commonItems = make(map[rune]int)

		// Part 2
		if lineNumber%3 == 0 {
			for _, item := range rucksack {
				badgeItems[item] = 1
			}
		}
		if lineNumber%3 == 1 {
			for _, item := range rucksack {
				if _, ok := badgeItems[item]; ok {
					badgeItems[item] = 2
				}
			}
		}
		if lineNumber%3 == 2 {
			for _, item := range rucksack {
				if _, ok := badgeItems[item]; ok {
					if badgeItems[item] == 2 {
						badge = item
						break
					}
				}
			}
			badgePriority := priorities[badge]
			sumOfBadgePriorites += badgePriority
			badgeItems = make(map[rune]int)
		}
		lineNumber += 1
	}

	// Part 1
	fmt.Println("Part 1:", sumOfCommonItemPriorities)

	// Part 2
	fmt.Println("Part 2:", sumOfBadgePriorites)
}
