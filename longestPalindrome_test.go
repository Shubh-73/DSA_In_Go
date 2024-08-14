package main

import (
	"dataStructuresInGo/strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},
		{"racecar", "racecar"},
		{"abdadf", "dad"},
		{"daz", "d"},
	}

	for _, test := range tests {
		result := strings.LongestPalindromicSubstring(test.input)
		if result != test.expected {
			t.Errorf("for input %s, expected is %s, output was %s", test.input, test.expected, result)
		}
	}
}
