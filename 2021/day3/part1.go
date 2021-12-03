package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
)

func intArrayToString(arr []int) string {
  var strArr string
  for _, integer := range arr {
    strArr = strArr + strconv.Itoa(integer)
  }
  return strArr
}

func main() {
  fmt.Println("=== Day 3 (part 1) ===")

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

  // get counts of each bit
  inputLength := len(input)
  binaryLength := len(input[0])

  var bitCount = make([]int, binaryLength) // array of n 0's

  for _ , binaryNumber := range input {
    for i, bit := range binaryNumber {
      if bit == '1' {
        bitCount[i] = bitCount[i]+1
      }
    }
  }

  // generate gamma and epsilon rates
  var gammaRate = make([]int, binaryLength)
  var epsilonRate = make([]int, binaryLength)
  for i, count := range bitCount {
    if count > (inputLength/2) {
      gammaRate[i] = 1
      epsilonRate[i] = 0
    } else {
      gammaRate[i] = 0
      epsilonRate[i] = 1
    }
  }

  // convert gamma rate to decimal
  gammaRateBinary := intArrayToString(gammaRate)
  gammaRateDecimal, err := strconv.ParseInt(gammaRateBinary, 2, 64)
  if err != nil {
    log.Fatal(err)
  }

  // convert epsilon rate to decimal
  epsilonRateBinary := intArrayToString(epsilonRate)
  epsilonRateDecimal, err := strconv.ParseInt(epsilonRateBinary, 2, 64)
  if err != nil {
    log.Fatal(err)
  }

  // Print results
  fmt.Println(gammaRateDecimal * epsilonRateDecimal)
}
