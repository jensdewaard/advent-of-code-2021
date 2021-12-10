package day10

import "testing"

func Test_FindClosingDelimiter(t *testing.T) {
	tests := []struct {
		input    string
		expected byte
		found    byte
		err      string
		missing  string
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", ']', '}', "corrupt", ""},
		{"[[<[([]))<([[{}[[()]]]", ']', ')', "corrupt", ""},
		{"[{[{({}]{}}([{[{{{}}([]", ')', ']', "corrupt", ""},
		{"[<(<(<(<{}))><([]([]()", '>', ')', "corrupt", ""},
		{"<{([([[(<>()){}]>(<<{{", ']', '>', "corrupt", ""},
		{"[({(<(())[]>[[{[]{<()<>>", 0, 0, "incomplete", "}}]])})]"},
		{"[(()[<>])]({[<{<<[]>>(", 0, 0, "incomplete", ")}>]})"},
		{"(((({<>}<{<{<>}{[]{[]{}", 0, 0, "incomplete", "}}>}>))))"},
		{"{<[[]]>}<{[{[{[]{()[[[]", 0, 0, "incomplete", "]]}}]}]}>"},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 0, 0, "incomplete", "])}>"},
	}
	for _, tt := range tests {
		actual := FindClosingDelimiter(tt.input, []byte{})
		if actual.expected != tt.expected {
			t.Errorf("expected expected %s, got expected %s", string(tt.expected), string(actual.expected))
		}
		if actual.got != tt.found {
			t.Errorf("expected found %s, got found %s", string(tt.found), string(actual.got))
		}
		if actual.err != tt.err {
			t.Errorf("expected error %s, got error %s", tt.err, actual.err)
		}
	}
}

func Test_ScoreMissing(t *testing.T) {
	tests := []struct {
		m        string
		expected int
	}{
		{"}}]])})]", 288957},
		{")}>]})", 5566},
		{"}}>}>))))", 1480781},
		{"]]}}]}]}>", 995444},
		{"])}>", 294},
	}
	for _, tt := range tests {
		actual := ScoreMissing(tt.m)
		if actual != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, actual)
		}
	}
}
