package main

import "testing"

var sequences = [][]string{
	{"75", "47", "61", "53", "29"},
	{"97", "61", "53", "29", "13"},
	{"75", "29", "13"},
	{"75", "97", "47", "61", "53"},
	{"61", "13", "29"},
	{"97", "13", "75", "29", "47"},
}

var rules = []Rule{
	NewRule("47", "53"),
	NewRule("97", "13"),
	NewRule("97", "61"),
	NewRule("97", "47"),
	NewRule("75", "29"),
	NewRule("61", "13"),
	NewRule("75", "53"),
	NewRule("29", "13"),
	NewRule("97", "29"),
	NewRule("53", "29"),
	NewRule("61", "53"),
	NewRule("97", "53"),
	NewRule("61", "29"),
	NewRule("47", "13"),
	NewRule("75", "47"),
	NewRule("97", "75"),
	NewRule("47", "61"),
	NewRule("75", "61"),
	NewRule("47", "29"),
	NewRule("75", "13"),
	NewRule("53", "13"),
}

func TestSumMiddleNumbersFromValidSequences(t *testing.T) {

	want := 143

	got := SumMiddleNumbersFromValidSequences(sequences, rules)

	if want != got {
		t.Errorf("TestSumMiddleNumbersFromValidSequences - want %d but got %d", want, got)
	}
}

func TestSumMiddleNumbersFromCorrectedInvalidSequences(t *testing.T) {

	want := 123

	got := SumMiddleNumbersFromCorrectedInvalidSequences(sequences, rules)

	if want != got {
		t.Errorf("TestSumMiddleNumbersFromCorrectedInvalidSequences - want %d but got %d", want, got)
	}
}
