package day7

import "testing"

func Test_CrabFuel(t *testing.T){
    tests := []struct {
        from int
        to int
        expected int
    }{
        {16, 5, 66},
        {1, 5, 10},
        {2, 5, 6},
        {0, 5, 15},
        {4, 5, 1},
        {7, 5, 3},
        {14, 5, 45},
    }
    for _, tt := range tests {
        actual := CrabFuel(tt.from, tt.to)
        if actual != tt.expected {
            t.Errorf("expected %v, got %v", tt.expected, actual)
        }
    }
}