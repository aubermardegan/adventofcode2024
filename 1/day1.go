package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SumDistance(arr1, arr2 []int) int {

	var distance int

	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	for i := 0; i < len(arr1); i++ {
		sum := arr2[i] - arr1[i]
		if sum < 0 {
			sum = sum * -1
		}
		distance += sum
	}

	return distance
}

func SumSimilarityScore(arr1, arr2 []int) int {

	var similarityScore int

	for _, num1 := range arr1 {
		countSameNumber := 0

		for _, num2 := range arr2 {
			if num2 == num1 {
				countSameNumber++
			}
		}

		similarityScore += (num1 * countSameNumber)
	}

	return similarityScore
}

const separator = "   "

func main() {
	var arr1, arr2 []int

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("error reading file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(inputBytes))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		num1, num2, err := extractNumbers(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	distance := SumDistance(arr1, arr2)
	fmt.Printf("\nPart 1 - Distance travelled: %d", distance)

	similarityScore := SumSimilarityScore(arr1, arr2)
	fmt.Printf("\nPart 2 - Similarity Score: %d", similarityScore)

}

func extractNumbers(line string) (int, int, error) {
	nums := strings.Split(line, separator)
	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, err
	}
	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}
