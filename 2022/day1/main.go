package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Return the sum of the top N numbers from an array
func sumOfTopN(n int, arr []int) int {
	sum := 0
	arrLen := len(arr)
	for i := 1; i <= n; i++ {
		sum = sum + arr[arrLen-i]
	}

	return sum
}

func main() {
	// Read input file
	f, err := os.Open("./input")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	elfCalories := make([]int, 0)
	calorieSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfCalories = append(elfCalories, calorieSum)
			calorieSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Failed to parse %s\n", line)
				return
			}
			calorieSum += calories
		}
	}
	sort.Ints(elfCalories)

	// Part 1 - max calories
	part1 := sumOfTopN(1, elfCalories)
	fmt.Println("part 1:", part1)

	// Part 2 - top three max calories
	part2 := sumOfTopN(3, elfCalories)
	fmt.Println("part 2:", part2)
}
