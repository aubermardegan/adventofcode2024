package main

import "testing"

func TestSumDistance(t *testing.T) {
	arr1 := []int{3, 4, 2, 1, 3, 3}
	arr2 := []int{4, 3, 5, 3, 9, 3}

	want := 11

	got := SumDistance(arr1, arr2)

	if want != got {
		t.Errorf("TestSumDistance - want %d but got %d", want, got)
	}
}

func TestSumSimilarityScore(t *testing.T) {
	arr1 := []int{3, 4, 2, 1, 3, 3}
	arr2 := []int{4, 3, 5, 3, 9, 3}

	want := 31

	got := SumSimilarityScore(arr1, arr2)

	if want != got {
		t.Errorf("TestSumSimilarityScore - want %d but got %d", want, got)
	}
}
