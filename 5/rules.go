package main

type Rule struct {
	Before string
	After  string
}

func NewRule(before, after string) Rule {
	return Rule{Before: before, After: after}
}
