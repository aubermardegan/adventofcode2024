package main

import "testing"

func TestSearchWord(t *testing.T) {
	matrix := [][]byte{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	want := 18

	got := SearchWord("XMAS", matrix)

	if want != got {
		t.Errorf("TestSearchWord - want %d but got %d", want, got)
	}
}

func TestXSearchMAS(t *testing.T) {
	matrix := [][]byte{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	want := 9

	got := XSearchMAS(matrix)

	if want != got {
		t.Errorf("TestXSearchMAS - want %d but got %d", want, got)
	}
}
