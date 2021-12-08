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
