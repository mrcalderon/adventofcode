package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Vent struct {
	pointA Point
	pointB Point
}

func newPoint(x int, y int) *Point {
	point := new(Point)
	point.x = x
	point.y = y
	return point
}

func newVent(pointA Point, pointB Point) *Vent {
	vent := new(Vent)
	vent.pointA = pointA
	vent.pointB = pointB
	return vent
}

// Calculate absolute value of an integer
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	// Determine whether to count diagonal lines or not
	argsCount := len(os.Args)
	if argsCount == 1 || argsCount > 2 {
		fmt.Println("Error: must indicate 1 or 2")
		os.Exit(1)
	}
	part := os.Args[1]
	var includeDiagonals bool
	if part == "1" {
		fmt.Println("=== Day 5 (part 1) ===")
		includeDiagonals = false
	} else if part == "2" {
		fmt.Println("=== Day 5 (part 2) ===")
		includeDiagonals = true
	} else {
		fmt.Println("Error: valid arguments are: 1, 2")
		os.Exit(1)
	}

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

	// Get vents
	vents := make([]Vent, 0)
	_ = vents
	for _, line := range input {
		re := regexp.MustCompile(" -> ")
		split := re.Split(line, -1)
		strA := strings.Split(split[0], ",")
		strB := strings.Split(split[1], ",")

		x1, _ := strconv.Atoi(strA[0])
		y1, _ := strconv.Atoi(strA[1])
		pointA := newPoint(x1, y1)

		x2, _ := strconv.Atoi(strB[0])
		y2, _ := strconv.Atoi(strB[1])
		pointB := newPoint(x2, y2)

		vent := newVent(*pointA, *pointB)
		vents = append(vents, *vent)
	}

	/**
	 *   Generate all coordinates where vents are located,
	 *   keeping track of how many times a particular
	 *   coordinate comes up.
	 */
	coordinatesMap := make(map[Point]int)
	var x, y int
	for _, vent := range vents {
		pointA := vent.pointA
		pointB := vent.pointB

		if pointA.x == pointB.x {
			// vertical line
			x = pointA.x
			if pointA.y < pointB.y { // example: A=(2,2), B=(2,5)
				for i := pointA.y; i <= pointB.y; i++ {
					if _, ok := coordinatesMap[*newPoint(x, i)]; ok {
						coordinatesMap[*newPoint(x, i)]++
					} else {
						coordinatesMap[*newPoint(x, i)] = 1
					}
				}
			} else if pointA.y > pointB.y { // example: A=(2,7), B=(2,5)
				for i := pointB.y; i <= pointA.y; i++ {
					if _, ok := coordinatesMap[*newPoint(x, i)]; ok {
						coordinatesMap[*newPoint(x, i)]++
					} else {
						coordinatesMap[*newPoint(x, i)] = 1
					}
				}
			}
		} else if pointA.y == pointB.y {
			// horizontal line
			y = pointA.y
			if pointA.x < pointB.x { // example: A=(1,3), B=(3,3)
				for i := pointA.x; i <= pointB.x; i++ {
					if _, ok := coordinatesMap[*newPoint(i, y)]; ok {
						coordinatesMap[*newPoint(i, y)]++
					} else {
						coordinatesMap[*newPoint(i, y)] = 1
					}
				}
			} else if pointA.x > pointB.x { // example: A=(5,3), B=(3,3)
				for i := pointB.x; i <= pointA.x; i++ {
					if _, ok := coordinatesMap[*newPoint(i, y)]; ok {
						coordinatesMap[*newPoint(i, y)]++
					} else {
						coordinatesMap[*newPoint(i, y)] = 1
					}
				}
			}
		} else if abs(pointA.x-pointB.x) == abs(pointA.y-pointB.y) && includeDiagonals {
			// diagonal line
			diff := abs(pointA.x - pointB.x) // absolute value of Xa - Xb
			diffX := pointA.x - pointB.x
			diffY := pointA.y - pointB.y
			if diffX == diffY { // example: A=(1,1), B=(3,3)
				for i := 0; i <= diff; i++ {
					if diffX > 0 {
						x = pointA.x - i
						y = pointA.y - i
					} else if diffX < 0 {
						x = pointA.x + i
						y = pointA.y + i
					}
					if _, ok := coordinatesMap[*newPoint(x, y)]; ok {
						coordinatesMap[*newPoint(x, y)]++
					} else {
						coordinatesMap[*newPoint(x, y)] = 1
					}
				}
			} else if pointA.x > pointB.x { // example: A=(9,7), B=(7,9)
				for i := 0; i <= diff; i++ {
					x = pointA.x - i
					y = pointA.y + i
					if _, ok := coordinatesMap[*newPoint(x, y)]; ok {
						coordinatesMap[*newPoint(x, y)]++
					} else {
						coordinatesMap[*newPoint(x, y)] = 1
					}
				}
			} else if pointB.x > pointA.x { // example: A=(7,9), B=(9,7)
				for i := 0; i <= diff; i++ {
					x = pointB.x - i
					y = pointB.y + i
					if _, ok := coordinatesMap[*newPoint(x, y)]; ok {
						coordinatesMap[*newPoint(x, y)]++
					} else {
						coordinatesMap[*newPoint(x, y)] = 1
					}
				}
			}
		} else {
			// not a valid line; skip
			continue
		}
	}

	// Find number of points with two or more lines overlapping
	pointsWithOverlap := 0
	for _, v := range coordinatesMap {
		if v > 1 {
			pointsWithOverlap++
		}
	}
	fmt.Println("Points with two or more lines overlapping:", pointsWithOverlap)
}
