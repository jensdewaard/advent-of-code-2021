package day3

import (
	"log"
	"math"
)

func Count(state BitCount, bit string) BitCount {
	if state.Bit == bit {
		return BitCount{
			Bit:   state.Bit,
			Count: state.Count + 1,
		}
	} else {
		if state.Count == 0 {
			return BitCount{bit, 1}
		}
		if state.Count == 1 {
			return BitCount{
				Bit:   bit,
				Count: 1,
			}
		} else {
			return BitCount{
				Bit:   state.Bit,
				Count: state.Count - 1,
			}
		}
	}
}

func ZipWith(f func(BitCount, string) BitCount, bs []BitCount, ss string) []BitCount {
	if len(bs) != len(ss) {
		log.Fatal("zipWith: lengths mismatch")
	}
	out := make([]BitCount, len(bs))
	for i := range bs {
		out[i] = f(bs[i], string(ss[i]))
	}
	return out
}

func FoldToBitCount(
	acc []BitCount,
	f func([]BitCount, string) []BitCount,
	ss []string) []BitCount {
	for _, s := range ss {
		acc = f(acc, s)
	}
	return acc
}

func FoldToString(acc string, f func(string, string) string, bs []BitCount) string {
	for _, b := range bs {
		acc = f(acc, b.Bit)
	}
	return acc
}

func AsDecimal(b BitString) float64 {
	if b == "" {
		return 0
	}
	if string(b[0]) == "0" {
		return AsDecimal(b[1:])
	}
	power := len(b)
	return math.Pow(2, float64(power-1)) + AsDecimal(b[1:])
}

func Invert(b BitString) BitString {
	out := ""
	for _, c := range b {
		if string(c) == "0" {
			out += "1"
		} else {
			out += "0"
		}
	}
	return out
}
