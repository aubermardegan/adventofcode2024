package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const inputFile = "input"

func XSearchMAS(puzzle [][]byte) int {
	var totalMatches int

	for i := 1; i < len(puzzle)-1; i++ {
		for j := 1; j < len(puzzle)-1; j++ {
			if puzzle[i][j] == 'A' {
				if checkMS(puzzle[i-1][j-1], puzzle[i+1][j-1], puzzle[i-1][j+1], puzzle[i+1][j+1]) {
					totalMatches++
				}
			}
		}
	}

	return totalMatches
}

func checkMS(TL, TR, BL, BR byte) bool {
	s := []byte{TL, TR, BL, BR}

	s = bytes.ReplaceAll(bytes.ReplaceAll(s, []byte{'M'}, []byte{}), []byte{'S'}, []byte{})
	if len(s) > 0 {
		return false
	}

	if TL != BR && TR != BL {
		return true
	}
	return false
}

func SearchWord(word string, puzzle [][]byte) int {
	var totalMatches int

	totalMatches = horizontalSearch(word, puzzle)

	totalMatches += verticalSearch(word, puzzle)

	totalMatches += diagonalSearch(word, puzzle)

	totalMatches += invertedDiagonalSearch(word, puzzle)

	return totalMatches
}

func horizontalSearch(word string, puzzle [][]byte) int {
	matches := 0

	for _, line := range puzzle {
		matches += findMatches(word, line)
	}

	return matches
}

func verticalSearch(word string, puzzle [][]byte) int {
	matches := 0

	transposedPuzzle := transposePuzzle(puzzle)

	for _, line := range transposedPuzzle {
		matches += findMatches(word, line)
	}
	return matches
}

func diagonalSearch(word string, puzzle [][]byte) int {
	matches := 0

	puzzleLength := len(puzzle)
	bottomLeftDiagonalPuzzle := make([][]byte, puzzleLength)
	topRightDiagonalPuzzle := make([][]byte, puzzleLength)

	for i := 0; i < len(puzzle); i++ {
		bottomLeftDiagonalPuzzle[i] = make([]byte, puzzleLength)
		for j := 0; j < puzzleLength; j++ {
			if i+j < puzzleLength {
				bottomLeftDiagonalPuzzle[i][j] = puzzle[i+j][j]
			}
		}
	}

	for _, line := range bottomLeftDiagonalPuzzle {
		matches += findMatches(word, line)
	}

	transposedPuzzle := transposePuzzle(puzzle)
	for i := 0; i < len(puzzle); i++ {
		topRightDiagonalPuzzle[i] = make([]byte, puzzleLength)
		for j := 0; j < puzzleLength; j++ {
			if i+j < puzzleLength {
				topRightDiagonalPuzzle[i][j] = transposedPuzzle[i+j][j]
			}
		}
	}

	for i, line := range topRightDiagonalPuzzle {
		if i == 0 {
			continue //para evitar conferir 2x a linha do meio, que já foi conferida no bottom left
		}
		matches += findMatches(word, line)
	}
	return matches
}

func invertedDiagonalSearch(word string, puzzle [][]byte) int {

	invertedPuzzle := invertPuzzle(puzzle)
	return diagonalSearch(word, invertedPuzzle)
}

func findMatches(word string, line []byte) int {
	var matches int
	wordLength := len(word)
	lineLength := len(line)

	for i := range line {
		match := true
		if line[i] == word[0] {
			if lineLength >= i+wordLength {
				for j := 1; j < wordLength; j++ {
					if line[i+j] != word[j] {
						match = false
						break
					}
				}
				if match {
					matches++
				}
			}
			match := true
			if (i+1)-wordLength >= 0 {
				for j := -1; j > -wordLength; j-- {
					if line[i+j] != word[-j] {
						match = false
						break
					}
				}
				if match {
					matches++
					match = false
				}
			}
		}
	}
	return matches
}

func transposePuzzle(puzzle [][]byte) [][]byte {
	puzzleLength := len(puzzle)

	transposedPuzzle := make([][]byte, puzzleLength)
	for i := 0; i < len(puzzle); i++ {
		transposedPuzzle[i] = make([]byte, puzzleLength)
		for j := 0; j < puzzleLength; j++ {
			transposedPuzzle[i][j] = puzzle[j][i]
		}
	}
	return transposedPuzzle
}

func invertPuzzle(puzzle [][]byte) [][]byte {
	puzzleLength := len(puzzle)

	invertedPuzzle := make([][]byte, puzzleLength)
	for i := 0; i < len(puzzle); i++ {
		invertedPuzzle[i] = make([]byte, puzzleLength)
		for j := 0; j < puzzleLength; j++ {
			invertedPuzzle[i][j] = puzzle[i][(puzzleLength-1)-j]
		}
	}
	return invertedPuzzle
}

func main() {
	var puzzle [][]byte

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("error reading file")
	}

	puzzle = bytes.Split(inputBytes, []byte("\n"))

	wordCount := SearchWord("XMAS", puzzle)
	fmt.Printf("\nPart 1 - WordCount: %d", wordCount)

	xmasCount := XSearchMAS(puzzle) //1905
	fmt.Printf("\nPart 2 - X-MAS count: %d", xmasCount)
}
