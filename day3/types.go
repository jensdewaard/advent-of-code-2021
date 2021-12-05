package day3

type BitString = string

type BitCount struct {
	Bit   string
	Count int
}

type BitStringFilter func(BitString) bool