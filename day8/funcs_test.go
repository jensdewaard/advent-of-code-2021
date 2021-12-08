package day8

import "testing"

func Test_Recognize(t *testing.T) {
	tests := []struct {
		s        LitSegments
		expected int
	}{
		{"fdgacbe", 8},
		{"gcbe", 4},
		{"ab", 1},
		{"bgf", 7},
		{"a", -1},
		{"abcde", -1},
	}
	for _, tt := range tests {
		actual := Recognize(tt.s)
		if actual.Number != tt.expected {
			t.Errorf("for %s expected %v, got %v", tt.s, tt.expected, actual)
		}
	}
}

func Test_Difference(t *testing.T) {
	tests := []struct {
		a, b     LitSegments
		expected LitSegments
	}{
		{
			"dab", // 7
			"ab",  // 1
			"d",   // top segment!
		},
		{
			"efgcdab", // 8
			"ebadf",
			"gc",
		},
		{
			"ebadf",
			"efgcdab", // 8
			"gc",
		},
		{
			"abc",
			"abc",
			"",
		},
		{"abc", "cab", ""},
	}
	for _, tt := range tests {
		actual := Difference(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, actual)
		}
	}
}

func Test_EquallyLit(t *testing.T) {
	tests := []struct {
		a, b     LitSegments
		expected bool
	}{
		{"abc", "abc", true},
		{"abd", "abc", false},
		{"abc", "cab", true},
		{"bca", "abc", true},
		{"abcdef", "ab", false},
	}
	for _, tt := range tests {
		actual := EquallyLit(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("%s == %s, expected %v, got %v", tt.a, tt.b, tt.expected, actual)
		}
	}
}

func Test_Solve(t *testing.T) {
	tests := []struct {
		rs, us   []RecognizedNumber
		expected []RecognizedNumber
	}{
		{
			[]RecognizedNumber{{"ab", 1}, {"acedgfb", 8}, {"eafb", 4}, {"dab", 7}},
			[]RecognizedNumber{{"cdfbe", -1}, {"gcdfa", -1}, {"fbcad", -1}, {"cefabd", -1}, {"cdfgeb", -1}, {"cagedb", -1}},
			[]RecognizedNumber{{"ab", 1}, {"acedgfb", 8}, {"eafb", 4}, {"dab", 7}, {"cdfbe", 5}, {"gcdfa", 2}, {"fbcad", 3}, {"cefabd", 9}, {"cdfgeb", 6}, {"cagedb", 0}},
		},
	}
	for _, tt := range tests {
		actual := Solve(tt.rs, tt.us)
		if !RecognizedEqual(actual, tt.expected) {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
