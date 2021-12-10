package day10

func FindClosingDelimiter(l string, delims []byte) LineStatus {
	if len(l) == 0 && len(delims) != 0 {
		return LineStatus{"incomplete", 0, 0, string(delims)}
	}
	if len(l) == 0 && len(delims) == 0 {
		return LineStatus{"valid", 0, 0, ""}
	}
	switch l[0] {
	case '{':
		return FindClosingDelimiter(
			l[1:], append([]byte{'}'}, delims...),
		)
	case '(':
		return FindClosingDelimiter(
			l[1:], append([]byte{')'}, delims...),
		)
	case '<':
		return FindClosingDelimiter(
			l[1:], append([]byte{'>'}, delims...),
		)
	case '[':
		return FindClosingDelimiter(
			l[1:], append([]byte{']'}, delims...),
		)
	case '}':
		if delims[0] != '}' {
			return LineStatus{"corrupt", delims[0], '}', ""}
		}
		return FindClosingDelimiter(l[1:], delims[1:])
	case ')':
		if delims[0] != ')' {
			return LineStatus{"corrupt", delims[0], ')', ""}
		}
		return FindClosingDelimiter(l[1:], delims[1:])
	case '>':
		if delims[0] != '>' {
			return LineStatus{"corrupt", delims[0], '>', ""}
		}
		return FindClosingDelimiter(l[1:], delims[1:])
	case ']':
		if delims[0] != ']' {
			return LineStatus{"corrupt", delims[0], ']', ""}
		}
		return FindClosingDelimiter(l[1:], delims[1:])
	}
	return LineStatus{"error", 0, 0, ""}
}

func Score(l LineStatus) int {
	switch l.got {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 1
}

func ScoreMissing(m string) int {
	total := 0
	for _, b := range m {
		switch b {
		case ')':
			total = total*5 + 1
		case ']':
			total = total*5 + 2
		case '}':
			total = total*5 + 3
		case '>':
			total = total*5 + 4
		}
	}
	return total
}

func Filter(pred func(LineStatus) bool, ls []LineStatus) []LineStatus {
	out := make([]LineStatus, 0)
	for _, l := range ls {
		if pred(l) {
			out = append(out, l)
		}
	}
	return out
}

func MapToInt(f func(LineStatus) int, ls []LineStatus) []int {
	is := make([]int, len(ls))
	for i, l := range ls {
		is[i] = f(l)
	}
	return is
}

func MapToStatus(f func(string) LineStatus, ss []string) []LineStatus {
	ls := make([]LineStatus, len(ss))
	for i, s := range ss {
		ls[i] = f(s)
	}
	return ls
}

func MapToString(f func(LineStatus) string, ls []LineStatus) []string {
	ss := make([]string, len(ls))
	for i, l := range ls {
		ss[i] = f(l)
	}
	return ss
}
