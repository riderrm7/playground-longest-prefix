package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	tc := []struct {
		input    []string
		expected string
	}{
		{input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{input: []string{"flower", "flow", "flight", "fill", "aero", "ae", "al"}, expected: "f"},
		{input: []string{"book", "boon", "boost", "boot"}, expected: "boo"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
		{input: []string{""}, expected: ""},
		{input: []string{"single"}, expected: "single"},
		{input: []string{}, expected: ""},
		{input: []string{"same", "same", "same"}, expected: "same"},
	}

	for i, testCase := range tc {
		result := findLongestPrefix(testCase.input)
		fmt.Printf("Test case %d: input: %v, expected: %s, got: %s\n", i+1, testCase.input, testCase.expected, result)
	}

}

func findLongestPrefix(arr []string) string {
	//Empty array
	if len(arr) == 0 {
		return ""
	}
	//Empty string
	if arr[0] == "" {
		return ""
	}
	/*
		sort when use case ex. ["flower", "flow", "flight","alpha","aero","al"]
	*/
	sort.Strings(arr)

	strGroup := map[string][]string{}

	for _, v := range arr {
		firstChar := v[:1]
		strGroup[firstChar] = append(strGroup[firstChar], v)
	}

	matched := ""

	for _, v := range strGroup {
		pf := findMatchingPrefix(v)

		if len(pf) >= len(matched) {
			matched = pf
		}
	}

	return matched
}

func findMatchingPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	pf := arr[0]
	for _, s := range arr[1:] {
		for len(pf) > 0 && !strings.HasPrefix(s, pf) {
			pf = pf[:len(pf)-1]
		}
		if pf == "" {
			break
		}
	}
	return pf
}
