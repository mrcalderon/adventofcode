package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func sum(array []int) int {
  result := 0
  for _, v := range array {
    result += v
  }
  return result
}

func main() {
  fmt.Println("=== Day 1 (part 2) ===")

  // open file
  f, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  var puzzle_input []int
  for scanner.Scan() {
    number, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal(err)
    }
    puzzle_input = append(puzzle_input, number)
  }

  var a, b int
  var three_measurement_sum int
  var sliding_window []int
  previous_sum := 9999999
  larger_sums := 0
  for i, _ := range puzzle_input {
    a = i
    b = i+3
    // Return when we can no longer create a new three-measurement sum
    if b > len(puzzle_input) {
      fmt.Println("Number of larger sums:", larger_sums)
      break
    }

    sliding_window = puzzle_input[a:b]

    three_measurement_sum = sum(sliding_window)
    if three_measurement_sum > previous_sum {
      larger_sums++
    }

    previous_sum = three_measurement_sum
  }
}
