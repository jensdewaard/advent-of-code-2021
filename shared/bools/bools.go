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
