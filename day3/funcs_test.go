package day3

import (
	"testing"

	"github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func Test_Count(t *testing.T) {
	tests := []struct {
		state    BitCount
		bit      string
		expected BitCount
	}{
		{
			BitCount{"1", 5},
			"0",
			BitCount{"1", 4},
		},
		{
			BitCount{"0", 1},
			"1",
			BitCount{"0", 0},
		},
		{
			BitCount{"0", 0},
			"1",
			BitCount{"1", 1},
		},
		{
			BitCount{"1", 0},
			"0",
			BitCount{"0", 1},
		},
		{
			BitCount{},
			"1",
			BitCount{"1", 1},
		},
		{
			BitCount{},
			"0",
			BitCount{"0", 1},
		},
	}
	for _, tt := range tests {
		actual := Count(tt.state, tt.bit)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_AsDecimal(t *testing.T) {
	tests := []struct {
		b        BitString
		expected float64
	}{
		{"0", 0},
		{"1", 1},
		{"01001", 9},
		{"10110", 22},
	}
	for _, tt := range tests {
		actual := AsDecimal(tt.b)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_Invert(t *testing.T) {
	tests := []struct {
		b        BitString
		expected BitString
	}{
		{"00", "11"},
		{"1010", "0101"},
		{"", ""},
	}
	for _, tt := range tests {
		actual := Invert(tt.b)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_Filter(t *testing.T) {
	tests := []struct {
		bs       []BitString
		p        func(b BitString) bool
		expected []BitString
	}{
		{
			[]BitString{"000", "", "10101", "111"},
			func(b BitString) bool { return len(b) == 3 },
			[]BitString{"000", "111"},
		},
		{
			[]BitString{},
			func(b BitString) bool { return true },
			[]BitString{},
		},
		{
			[]BitString{"000", "", "10101", "111"},
			func(b BitString) bool { return true },
			[]BitString{"000", "", "10101", "111"},
		},
	}
	for _, tt := range tests {
		actual := Filter(tt.p, tt.bs)
		if !strings.SliceEqual(actual, tt.expected) {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
