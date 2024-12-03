package main

import "regexp"

type RegexMatch struct {
	Index int
	Value string
}

func ExtractRegexFromMemory(memory, regex string) []RegexMatch {
	var regexMatches []RegexMatch

	reg := regexp.MustCompile(regex)

	matches := reg.FindAllStringSubmatch(memory, -1)
	matchesIndexes := reg.FindAllStringSubmatchIndex(memory, -1)

	for i := range matches {
		regMatch := RegexMatch{Index: matchesIndexes[i][0], Value: matches[i][0]}
		regexMatches = append(regexMatches, regMatch)
	}

	return regexMatches
}
