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

func SumMiddleNumbersFromValidSequences(sequences [][]string, rules []Rule) int {
	var result int

	ruleSet := make(map[string][]Rule)
	for _, rule := range rules {
		rules, ok := ruleSet[rule.After]
		if !ok {
			rules = []Rule{}
		}
		rules = append(rules, rule)
		ruleSet[rule.After] = rules
	}

	for _, sequence := range sequences {
		validSequence := true
		for i := len(sequence) - 1; i >= 0; i-- {
			valid, _ := isValidSubSequence(ruleSet, sequence[i], sequence[i:])
			if !valid {
				validSequence = false
			}
		}
		if validSequence {
			number, err := strconv.Atoi(sequence[(len(sequence)-1)/2])
			if err != nil {
				log.Fatal(err)
			}
			result += number
		}
	}

	return result
}

func SumMiddleNumbersFromCorrectedInvalidSequences(sequences [][]string, rules []Rule) int {
	var result int

	ruleSet := make(map[string][]Rule)
	for _, rule := range rules {
		rules, ok := ruleSet[rule.After]
		if !ok {
			rules = []Rule{}
		}
		rules = append(rules, rule)
		ruleSet[rule.After] = rules
	}

	for _, sequence := range sequences {
		for i := len(sequence) - 1; i >= 0; i-- {
			valid, wrongIndex := isValidSubSequence(ruleSet, sequence[i], sequence[i:])
			if !valid {
				correctedSequence := orderSequence(ruleSet, sequence, i, i+wrongIndex)
				number, err := strconv.Atoi(correctedSequence[(len(correctedSequence)-1)/2])
				if err != nil {
					log.Fatal(err)
				}
				result += number
			}
		}
	}

	return result
}

func orderSequence(ruleSet map[string][]Rule, sequence []string, index, wrongIndex int) []string {
	value := sequence[index]
	unorderedValue := sequence[wrongIndex]
	sequence[index] = unorderedValue
	sequence[wrongIndex] = value

	for i := len(sequence) - 1; i >= 0; i-- {
		valid, wrongIndex := isValidSubSequence(ruleSet, sequence[i], sequence[i:])
		if !valid {
			return orderSequence(ruleSet, sequence, i, i+wrongIndex)
		}
	}
	return sequence
}

func isValidSubSequence(ruleSet map[string][]Rule, page string, sequence []string) (valid bool, wrongIndex int) {
	rules, ok := ruleSet[page]
	if !ok {
		return true, 0
	}

	for i, page := range sequence {
		for _, rule := range rules {
			if rule.Before == page {
				return false, i
			}
		}
	}
	return true, 0
}

func main() {
	var rules []Rule
	var sequences [][]string

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("error reading file")
	}

	scanner := bufio.NewScanner(bytes.NewReader(inputBytes))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			before, after, ok := strings.Cut(line, "|")
			if !ok {
				log.Fatal("problema no parsing da linha")
			}
			rule := NewRule(before, after)
			rules = append(rules, rule)
		}
		if strings.Contains(line, ",") {
			sequences = append(sequences, strings.Split(line, ","))
		}
	}

	resultPart1 := SumMiddleNumbersFromValidSequences(sequences, rules)
	fmt.Printf("\nPart 1 - Sum of Middle Numbers on Valid Sequences: %d", resultPart1)

	resultPart2 := SumMiddleNumbersFromCorrectedInvalidSequences(sequences, rules)
	fmt.Printf("\nPart 2 - Sum of Middle Numbers on Corrected Invalid Sequences: %d", resultPart2)
}
