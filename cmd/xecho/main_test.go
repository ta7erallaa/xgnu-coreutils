package main

import (
	"bytes"
	"testing"
)

var xechoTests = []struct {
	name       string
	args       []string
	interprted bool
	noNewline  bool
	want       string
}{
	{"Print word", []string{"random"}, false, false, "random\n"},
	{"Print word with no newline", []string{"random"}, false, true, "random"},
	{"Print word with interpretaion", []string{`\trandom`}, true, false, "\trandom\n"},
	{"Print word with interpretaion noNewline", []string{`\trandom`}, true, true, "\trandom"},
	{"Print word with no interpretaion", []string{`\trandom`}, false, false, `\trandom` + "\n"},
	{"Print word with no interpretaion noNewline", []string{`\trandom`}, false, true, `\trandom`},

	{"Print words", []string{"one", "two"}, false, false, "one two\n"},
	{"Print words with no newline", []string{"one", "two"}, false, true, "one two"},
	{"Print words with interpretaion", []string{`one\ttwo`}, true, false, "one\ttwo\n"},
	{"Print words with interpretaion noNewline", []string{`one\ttwo`}, true, true, "one\ttwo"},
	{"Print words with no interpretaion", []string{`one\ttwo`}, false, false, `one\ttwo` + "\n"},
	{"Print words with no interpretaion noNewline", []string{`one\ttwo`}, false, true, `one\ttwo`},
}

func TestXecho(t *testing.T) {
	for _, test := range xechoTests {
		var b bytes.Buffer
		t.Run(test.name, func(t *testing.T) {
			xecho(&b, test.args, test.interprted, test.noNewline)
			if b.String() != test.want {
				t.Errorf("output:%s not equal wanted:%s\n", b.String(), test.want)
			}
		})
	}
}
