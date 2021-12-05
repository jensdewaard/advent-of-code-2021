package bools

func CountTrue(bs []bool) int {
	count := 0
	for _, b := range bs {
		if b {
			count++
		}
	}
	return count
}

func Invert(bs []bool) []bool {
	out := make([]bool, len(bs))
	for i, b := range bs {
		out[i] = !b
	}
	return out
}
