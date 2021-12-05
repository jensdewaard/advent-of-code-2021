package day5

import "testing"

func Test_Expand(t *testing.T) {
	tests := []struct {
		l        Line
		expected []Position
	}{
		{
			Line{Position{0, 0}, Position{0, 3}},
			[]Position{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			Line{Position{0, 0}, Position{0, 0}},
			[]Position{{0, 0}},
		},
		{
			Line{Position{3, 5}, Position{0, 5}},
			[]Position{{3, 5}, {2, 5}, {1, 5}, {0, 5}},
		},
		{
			Line{Position{1, 3}, Position{1, 0}},
			[]Position{{1, 3}, {1, 2}, {1, 1}, {1, 0}},
		},
		{
			Line{Position{2, 7}, Position{5, 7}},
			[]Position{{2, 7}, {3, 7}, {4, 7}, {5, 7}},
		},
		{
			Line{Position{0, 0}, Position{8, 8}},
			[]Position{},
		},
	}
	for _, tt := range tests {
		actual := Expand(tt.l)
		equals := true
		if len(actual) != len(tt.expected) {
			equals = false
		} else {
			for i, p := range tt.expected {
				if actual[i] != p {
					equals = false
				}
			}
		}
		if !equals {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_ExpandDiagonal(t *testing.T) {
	tests := []struct {
		l        Line
		expected []Position
	}{
		{
			Line{Position{0, 0}, Position{0, 3}},
			[]Position{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			Line{Position{0, 0}, Position{0, 0}},
			[]Position{{0, 0}},
		},
		{
			Line{Position{3, 5}, Position{0, 5}},
			[]Position{{3, 5}, {2, 5}, {1, 5}, {0, 5}},
		},
		{
			Line{Position{1, 3}, Position{1, 0}},
			[]Position{{1, 3}, {1, 2}, {1, 1}, {1, 0}},
		},
		{
			Line{Position{2, 7}, Position{5, 7}},
			[]Position{{2, 7}, {3, 7}, {4, 7}, {5, 7}},
		},
		{
			Line{Position{0, 0}, Position{8, 8}},
			[]Position{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}},
		},
		{
			Line{Position{2, 2}, Position{0, 0}},
			[]Position{{2, 2}, {1, 1}, {0, 0}},
		},
		{
			Line{Position{3, 3}, Position{0, 6}},
			[]Position{{3, 3}, {2, 4}, {1, 5}, {0, 6}},
		},
		{
			Line{Position{3, 3}, Position{6, 0}},
			[]Position{{3, 3}, {4, 2}, {5, 1}, {6, 0}},
		},
	}
	for _, tt := range tests {
		actual := ExpandDiagonal(tt.l)
		equals := true
		if len(actual) != len(tt.expected) {
			equals = false
		} else {
			for i, p := range tt.expected {
				if actual[i] != p {
					equals = false
				}
			}
		}
		if !equals {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}

}
