package main

import "testing"

var matrix = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestSumSafeReports(t *testing.T) {

	want := 2

	got := SumSafeReports(matrix)

	if want != got {
		t.Errorf("TestSumSafeReports - want %d but got %d", want, got)
	}
}

func TestSumSafeReportsWithProblemDampener(t *testing.T) {

	want := 4

	got := SumSafeReportsWithProblemDampener(matrix)

	if want != got {
		t.Errorf("TestSumSafeReportsWithProblemDampener - want %d but got %d", want, got)
	}
}
