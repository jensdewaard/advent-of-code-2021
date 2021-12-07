package day6

import (
	"testing"

	"github.com/jensdewaard/advent-of-code-2021/shared/ints"
)

func Test_Age(t *testing.T) {
	tests := []struct {
		fish     []Fish
		expected []Fish
	}{
		{[]Fish{3}, []Fish{2}},
		{[]Fish{2}, []Fish{1}},
		{[]Fish{1}, []Fish{0}},
		{[]Fish{0}, []Fish{6, 8}},
		{[]Fish{6, 8}, []Fish{5, 7}},
	}
	for _, tt := range tests {
		actual := ints.FlatMapToInt(
			Age,
			tt.fish,
		)
		if !ints.SliceEqual(actual, tt.expected) {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}

func Test_FishAfterDays(t *testing.T) {
	tests := []struct {
		days     int
		start    []Fish
		expected int
	}{
		{18, []Fish{3, 4, 3, 1, 2}, 26},
		{80, []Fish{3, 4, 3, 1, 2}, 5934},
		{256, []Fish{3, 4, 3, 1, 2}, 26984457539},
	}
	for _, tt := range tests {
		actual := FishAfterDaysSlice(tt.days, tt.start)
		if actual != tt.expected {
			t.Errorf("expected %v, got %v", tt.expected, actual)
		}
	}
}
