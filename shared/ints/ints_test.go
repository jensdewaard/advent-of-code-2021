package ints

import "testing"

func TestWindow(t *testing.T) {
	var tests = []struct {
		k        int
		in       []int
		expected [][]int
	}{
		{
			3,
			[]int{1, 2, 3, 4, 5},
			[][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
		{
			1,
			[]int{1, 2, 3},
			[][]int{{1}, {2}, {3}},
		},
		{
			3,
			[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			[][]int{
				{199, 200, 208},
				{200, 208, 210},
				{208, 210, 200},
				{210, 200, 207},
				{200, 207, 240},
				{207, 240, 269},
				{240, 269, 260},
				{269, 260, 263},
			},
		},
	}
	for _, tt := range tests {
		var actual = Window(tt.k, tt.in)
		if len(actual) != len(tt.expected) {
			t.Errorf(
				"expected length %d, got length %d",
				len(tt.expected),
				len(actual),
			)
		}
		for i := range tt.expected {
			if !SliceEqual(actual[i], tt.expected[i]) {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		}
	}
}

func TestFold(t *testing.T) {
	tests := []struct {
		acc      int
		f        func(int, int) int
		is       []int
		expected int
	}{
		{
			0,
			func(a, b int) int { return a + b },
			[]int{1, 2, 3},
			6,
		},
	}
	for _, tt := range tests {
		actual := Fold(tt.acc, tt.f, tt.is)
		if actual != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, actual)
		}
	}
}
