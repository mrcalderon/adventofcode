package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== Day 1 (part 1) ===")

	// open file
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	previous_measurement := 9999999
	larger_measurements := 0
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if number > previous_measurement {
			larger_measurements++
		}
		previous_measurement = number
	}

	fmt.Println("Number of measurements larger than the previous one: " + strconv.Itoa(larger_measurements))

}
