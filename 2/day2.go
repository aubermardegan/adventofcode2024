package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SumSafeReportsWithProblemDampener(reportMatrix [][]int) int {
	var safeReports int

	for _, report := range reportMatrix {
		if isSafeReportWithProblemDampener(report) {
			safeReports++
		}
	}

	return safeReports
}

func isSafeReportWithProblemDampener(report []int) bool {
	var errorCount int
	var desc bool
	for i, value := range report {
		if i == 0 {
			if isSafeReport(report[1:]) {
				return true
			}

			if report[i+1] < value {
				desc = true
			}
			continue
		}

		previousValue := report[i-1]
		levelDiff := 0

		if value == previousValue {
			if isSafeReportRemovingCurrentOrLastElement(report, i) {
				return true
			}
			errorCount++
		}

		if desc {
			if value > previousValue {
				if isSafeReportRemovingCurrentOrLastElement(report, i) {
					return true
				}
				errorCount++
			}
			levelDiff = previousValue - value
		} else {
			if value < previousValue {
				if isSafeReportRemovingCurrentOrLastElement(report, i) {
					return true
				}
				errorCount++
			}
			levelDiff = value - previousValue
		}

		if levelDiff > 3 {
			if isSafeReportRemovingCurrentOrLastElement(report, i) {
				return true
			}
			errorCount++
		}
	}
	if errorCount > 0 {
		return isSafeReport(report[:len(report)-1])
	}
	return true
}

func isSafeReportRemovingCurrentOrLastElement(report []int, index int) bool {
	return isSafeReport(copySliceRemovingElementByIndex(report, index)) || isSafeReport(copySliceRemovingElementByIndex(report, index-1))
}

func copySliceRemovingElementByIndex(slice []int, index int) []int {
	buf := make([]int, len(slice))
	copy(buf, slice)
	if index == len(slice) {
		return buf[:len(slice)-1]
	}
	return append(buf[:index], buf[index+1:]...)
}

func SumSafeReports(reportMatrix [][]int) int {
	var safeReports int

	for _, report := range reportMatrix {
		if isSafeReport(report) {
			safeReports++
		}
	}

	return safeReports
}

func isSafeReport(report []int) bool {
	var desc bool
	for i, value := range report {
		if i == 0 {
			if report[i+1] < value {
				desc = true
			}
			continue
		}

		previousValue := report[i-1]
		levelDiff := 0

		if value == previousValue {
			return false
		}

		if desc {
			if value > previousValue {
				return false
			}
			levelDiff = previousValue - value
		} else {
			if value < previousValue {
				return false
			}
			levelDiff = value - previousValue
		}

		if levelDiff > 3 {
			return false
		}
	}
	return true
}

const inputFile = "input"
const separator = " "

func main() {
	var matrix [][]int

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("error reading file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(inputBytes))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		numbers, err := extractNumbers(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		matrix = append(matrix, numbers)
	}

	countSafeReports := SumSafeReports(matrix)
	fmt.Printf("\nPart 1 - Number of safe reports: %d", countSafeReports)

	countSafeReportsWithProblemDampener := SumSafeReportsWithProblemDampener(matrix)
	fmt.Printf("\nPart 2 - Number of safe reports (with problem dampener): %d", countSafeReportsWithProblemDampener)
}

func extractNumbers(line string) ([]int, error) {
	var numbers []int

	values := strings.Split(line, separator)
	for _, value := range values {
		number, err := strconv.Atoi(value)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

//293
