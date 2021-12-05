package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const rows int = 5
const cols int = 5

type Board struct {
	numbers [rows][cols]Number
}

type Number struct {
	value  int
	marked bool
}

// Create board from a 2D integer array
func newBoard(numbers [][]int) *Board {
	board := new(Board)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board.numbers[row][col].value = numbers[row][col]
		}
	}

	return board
}

// Check if a board exists in a list of boards
func containsBoard(board Board, boards []Board) bool {
	if len(boards) == 0 {
		return false
	}
contains:
	for _, b := range boards {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if b.numbers[i][j].value != board.numbers[i][j].value {
					continue contains
				}
			}
		}
		return true
	}
	return false
}

// Mark board if it has the provided number
func (b *Board) mark(number int) {
loop:
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if b.numbers[row][col].value == number {
				b.numbers[row][col].marked = true
				break loop
			}
		}
	}
}

// Print formatted board
func (b Board) pprint() {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			number := b.numbers[row][col]
			if number.marked == true {
				if col == cols-1 {
					fmt.Print(b.numbers[row][col].value, "*")
				} else {
					fmt.Print(b.numbers[row][col].value, "*", "\t")
				}
			} else {
				if col == cols-1 {
					fmt.Print(b.numbers[row][col].value)
				} else {
					fmt.Print(b.numbers[row][col].value, "\t")
				}
			}
		}
		fmt.Println()
	}
}

// Check if given board won bingo or not
func (b Board) won() bool {
	var marked int

	// check for row-based win
	for i := 0; i < rows; i++ {
		marked = 0
		for j := 0; j < cols; j++ {
			number := b.numbers[i][j]
			if number.marked == true {
				marked++
			}
		}
		if marked == cols {
			// fmt.Println("row-based win!")
			return true
		}
	}

	// check for column-based win
	for j := 0; j < cols; j++ {
		marked = 0
		for i := 0; i < rows; i++ {
			number := b.numbers[i][j]
			if number.marked == true {
				marked++
			}
		}
		if marked == rows {
			// fmt.Println("column-based win!")
			return true
		}
	}

	// no winning rows/columns found
	return false
}

func main() {
	fmt.Println("=== Day 4 (part 2) ===")

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

	// Get bingo numbers from the first line of input
	selectedNumbersStr := strings.Split(input[0], ",")
	selectedNumbers := make([]int, 0)
	for _, num := range selectedNumbersStr {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		selectedNumbers = append(selectedNumbers, numInt)
	}

	// Create the bingo boards
	boards := make([]Board, 0)

	boardNumbers := make([][]int, 0)
	for _, line := range input[2:] {
		if len(line) == 0 {
			board := newBoard(boardNumbers)
			boards = append(boards, *board)
			boardNumbers = make([][]int, 0)
		} else {
			inputRow := strings.Fields(line)
			inputRowInt := make([]int, 0)
			for _, num := range inputRow {
				num, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				inputRowInt = append(inputRowInt, num)
			}

			boardNumbers = append(boardNumbers, inputRowInt)
		}
	}

	// Find the last possible board to win bingo
	winningBoard := new(Board)
	winningBoards := make([]Board, 0)
	winningNumbers := make([]int, 0)
	var winningNumber int
	for _, num := range selectedNumbers {
		for i := range boards {
			board := &boards[i]
			board.mark(num)
			if board.won() == true {
				winningBoard = board
				winningNumber = num
				if containsBoard(*winningBoard, winningBoards) == false {
					winningBoards = append(winningBoards, *winningBoard)
					winningNumbers = append(winningNumbers, winningNumber)
				}
			}
		}
	}
	lastWinningNumber := winningNumbers[len(winningNumbers)-1]
	fmt.Println("last winning number:", lastWinningNumber)
	lastWinningBoard := winningBoards[len(winningBoards)-1]
	lastWinningBoard.pprint()

	// Get last winning board score
	unmarkedSum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			number := lastWinningBoard.numbers[i][j]
			if number.marked == false {
				unmarkedSum = unmarkedSum + number.value
			}
		}
	}
	score := lastWinningNumber * unmarkedSum
	fmt.Println("score:", score)
}
