package graphs

import "testing"

func Test_PathUnique(t *testing.T) {
	tests := []struct {
		p        Path
		expected bool
	}{
		{Path{Node{"a"}, Node{"b"}, Node{"c"}}, true},
		{Path{Node{"a"}, Node{"b"}, Node{"a"}}, false},
	}

	for _, tt := range tests {
		actual := PathUnique(tt.p)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
