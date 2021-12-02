package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  fmt.Println("=== Day 2 (part 2) ===")

  // open file
  f, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)

  depth := 0
  horizontal_pos := 0
  aim := 0
  for scanner.Scan() {
    command := scanner.Text()
    command_split := strings.Fields(command)
    direction := command_split[0]
    direction_amount, err := strconv.Atoi(command_split[1])
    if err != nil {
      log.Fatal(err)
    }

    // Adjust depth, horizontal position, and aim based on input directions
    switch direction {
      case "forward":
        horizontal_pos = horizontal_pos + direction_amount
        depth = depth + (aim * direction_amount)
      case "up":
        aim = aim - direction_amount
      case "down":
        aim = aim + direction_amount
    }
  }

  // Print results
  fmt.Println("depth:", depth)
  fmt.Println("horizontal position:", horizontal_pos)
  fmt.Println("depth * horizontal position =", depth*horizontal_pos)
}
