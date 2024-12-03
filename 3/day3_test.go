package main

import "testing"

func TestSumMultiples(t *testing.T) {

	memory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	want := 161

	got := SumMultiples(memory)

	if want != got {
		t.Errorf("TestSumMultiples - want %d but got %d", want, got)
	}
}

func TestSumEnabledMultiples(t *testing.T) {

	memory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	want := 48

	got := SumEnabledMultiples(memory)

	if want != got {
		t.Errorf("TestSumEnabledMultiples - want %d but got %d", want, got)
	}
}
