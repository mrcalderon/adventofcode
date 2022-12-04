package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read input file
	f, err := os.Open("./input")
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	pairs := make([]int, 4)
	var fullyContainedAssignments int
	var overlappingAssignments int
	for scanner.Scan() {
		line := scanner.Text()
		a := regexp.MustCompile(`-|,`)
		b := a.Split(line, -1)
		for i, v := range b {
			pairs[i], err = strconv.Atoi(v)
			if err != nil {
				return
			}
		}
		if ((pairs[0] - pairs[2]) * (pairs[1] - pairs[3])) <= 0 {
			fullyContainedAssignments += 1
		}
		if ((pairs[0] - pairs[3]) * (pairs[1] - pairs[2])) <= 0 {
			overlappingAssignments += 1
		}
	}

	// Part 1
	fmt.Println("Part 1:", fullyContainedAssignments)

	// Part 2
	fmt.Println("Part 2:", overlappingAssignments)
}
