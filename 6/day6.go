package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func CountDistinctPositions(matrix [][]byte) int {
	y, x := getStartingPosition(matrix)
	md := InitMovementDirector()

	matrixLength := len(matrix)
	guardInTheArea := true

	for guardInTheArea {
		matrix[y][x] = 'X'
		nextX, nextY := md.GetNextCoordinates(x, y)
		if nextX >= matrixLength || nextY >= matrixLength || nextX < 0 || nextY < 0 {
			matrix[y][x] = 'X'
			guardInTheArea = false
			break
		}

		if matrix[nextY][nextX] == '#' {
			md.ChangeDirection()
			continue
		}

		x = nextX
		y = nextY
	}

	return countDistinctPositions(matrix)
}

func getStartingPosition(matrix [][]byte) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == '^' {
				return i, j
			}
		}
	}
	return 0, 0
}

func countDistinctPositions(matrix [][]byte) int {
	var distinctPositionsCount int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == 'X' {
				distinctPositionsCount++
			}
		}
	}
	return distinctPositionsCount
}

func main() {
	var matrix [][]byte

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("error reading file")
	}

	matrix = bytes.Split(inputBytes, []byte("\n"))

	distinctPositionsCount := CountDistinctPositions(matrix)
	fmt.Printf("\nPart 1 - Distinct Positions Count: %d", distinctPositionsCount)
}
