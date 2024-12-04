package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SumEnabledMultiples(memory string) int {
	result := 0

	mulFuncRegex := `mul\(\d\d?\d?,\d\d?\d?\)`
	mulFuncMatches := ExtractRegexFromMemory(memory, mulFuncRegex)

	doRegex := `do\(\)`
	doMatches := ExtractRegexFromMemory(memory, doRegex)

	dontRegex := `don't\(\)`
	dontMatches := ExtractRegexFromMemory(memory, dontRegex)

	conditionalMatches := append(doMatches, dontMatches...)
	sort.Slice(conditionalMatches, func(i, j int) bool { return conditionalMatches[i].Index < conditionalMatches[j].Index })

	for _, match := range mulFuncMatches {
		if isEnabled(conditionalMatches, match.Index) {
			result += multiply(match.Value)
		}
	}

	return result
}

func SumMultiples(memory string) int {
	result := 0

	mulFuncRegex := `mul\(\d\d?\d?,\d\d?\d?\)`
	mulFuncMatches := ExtractRegexFromMemory(memory, mulFuncRegex)

	for _, match := range mulFuncMatches {
		result += multiply(match.Value)
	}

	return result
}

func multiply(functionDescription string) int {
	values := strings.Split(functionDescription[len("mul("):len(functionDescription)-1], ",")

	num1, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal(err)
	}

	return num1 * num2
}

func isEnabled(conditionalMatches []RegexMatch, index int) bool {
	for i, condition := range conditionalMatches {
		if index > condition.Index {
			continue
		}
		if i > 0 {
			return strings.EqualFold(conditionalMatches[i-1].Value, "do()")
		}
		return true
	}
	return strings.EqualFold(conditionalMatches[len(conditionalMatches)-1].Value, "do()")
}

func main() {

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("error reading file")
	}

	multiplicationResult := SumMultiples(strings.ReplaceAll(string(inputBytes), "\n", ""))
	fmt.Printf("\nPart 1 - Uncorrupted multiplication result: %d", multiplicationResult)

	enabledMultiplicationResult := SumEnabledMultiples(strings.ReplaceAll(string(inputBytes), "\n", ""))
	fmt.Printf("\nPart 2 - Uncorrupted enabled multiplication result: %d", enabledMultiplicationResult)
}
